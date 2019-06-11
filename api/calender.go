package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"time"
)

var calenderAPI = &ApiConfig{
	host: "https://www.googleapis.csom/",
	path: "calendar/v3/calendars/primary/events",
}

//@todo use client.get() do make a request
func (c *Client) GoogleCalender(r *GoogleCalenderRequest) (*GoogleCalenderResponse, error) {
	conf, err := getConfig(CalendarReadonlyScope)
	if err != nil {
		return nil, err
	}

	token, err := getToken()
	if err != nil {
		requestTokenFromWeb(conf)
	}

	cl, err := calendar.NewService(context.Background(), option.WithTokenSource(conf.TokenSource(context.Background(), token)))

	t := time.Now().Format(time.RFC3339)
	events, err := cl.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(8).OrderBy("startTime").Do()
	if err != nil {
		fmt.Printf("Unable to retrieve next ten of the user's events: %v", err)
	}

	jsobBody, err := events.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var response GoogleCalenderResponse
	if err = json.Unmarshal(jsobBody, &response); err != nil {
		return nil, err
	}

	/*
		fmt.Println("Upcoming events:")
		if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
		} else {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				fmt.Printf("%v (%v)\n", item.Summary, date)
			}
		}
	*/
	return &response, nil
}

// Reads the file with the clients credentials. Credentials can be downloaded
// from https://console.developers.google.com under "Credentials".
func getConfig(scope ...string) (*oauth2.Config, error) {
	f, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		return nil, errors.New("no client file found. Visit the following page to download your credentials: \n https://console.developers.google.com")
	}
	var conf *Credentials
	if err = json.Unmarshal(f, &conf); err != nil {
		return nil, errors.New("unable to unmarshal the credentials")
	}

	return &oauth2.Config{
		ClientID:     conf.Installed.ClientID,
		ClientSecret: conf.Installed.ClientSecret,
		RedirectURL:  conf.Installed.RedirectURL[0],
		Scopes:       scope,
		Endpoint: oauth2.Endpoint{
			AuthURL:  conf.Installed.AuthURL,
			TokenURL: conf.Installed.TokenURL,
		},
	}, nil
}

func getToken() (*oauth2.Token, error) {
	f, err := ioutil.ReadFile("./token.json")
	if err != nil {
		return nil, err
	}

	tok := &oauth2.Token{}
	err = json.Unmarshal(f, tok)
	if err != nil {
		return nil, errors.New("unable to unmarshal the token file")
	}

	return tok, nil
}

func requestTokenFromWeb(config *oauth2.Config) error {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	//png, err := qrcode.Encode(authURL, qrcode.Medium, 128)

	err := qrcode.WriteFile(authURL, qrcode.Medium, 128, "qr.png")
	if err != nil {
		return errors.New("unable to request the token")
	}

	//todo how are we gonna get the input from the mirror ?
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return errors.New(fmt.Sprintf("Unable to read authorization code: %v", err))
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"Unable to retrieve token from web: %v", err))
	}

	//Save token to file
	f, err := os.Create("./token.json")
	if err != nil {
		return err
	}
	defer f.Close()

	if err = json.NewEncoder(f).Encode(tok); err != nil {
		return err
	}
	return nil
}

type Credentials struct {
	Installed struct {
		// ClientID is the application's ID.
		ClientID string `json:"client_id"`
		// ClientSecret is the application's secret.
		ClientSecret string `json:"client_secret"`
		// Endpoint contains the resource server's token endpoint
		// URLs. These are constants specific to each server and are
		// often available via site-specific packages, such as
		// google.Endpoint or github.Endpoint.
		AuthURL  string `json:"auth_uri"`
		TokenURL string `json:"token_uri"`
		// RedirectURL is the URL to redirect users going through
		// the OAuth flow, after the resource owner's URLs.
		RedirectURL []string `json:"redirect_uris"`
		// Scope specifies optional requested permissions.
		Scopes []string

		// AuthStyle optionally specifies how the endpoint wants the
		// client ID & client secret sent. The zero value means to
		// auto-detect.
		AuthStyle int
	} `json:"installed"`
}

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page
// that asks for permissions for the required scopes explicitly.
//
// State is a token to protect the user from CSRF attacks. You must
// always provide a non-empty string and validate that it matches the
// the state query parameter on your redirect callback.
// See http://tools.ietf.org/html/rfc6749#section-10.12 for more info.
//
// Opts may include AccessTypeOnline or AccessTypeOffline, as well
// as ApprovalForce.
// It can also be used to pass the PKCE challenge.
// See https://www.oauth.com/oauth2-servers/pkce/ for more info.
func (c *Credentials) AuthCodeURL(state, accessType string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Installed.AuthURL)
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.Installed.ClientID},
	}
	if c.Installed.RedirectURL[0] != "" {
		v.Set("redirect_uri", c.Installed.RedirectURL[0])
	}
	if len(c.Installed.Scopes) > 0 {
		v.Set("scope", strings.Join(c.Installed.Scopes, " "))
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		v.Set("state", state)
	}
	if strings.Contains(c.Installed.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

type GoogleCalenderResponse struct {
	Items []struct {
		Start struct {
			Date     string `json:"date"`
			DateTime string `json:"dateTime,omitempty"`
		} `json:"start,omitempty"`
		End struct {
			Date string `json:"date"`
		} `json:"end,omitempty"`
		Summary string `json:"summary"`
	} `json:"items"`
}

type GoogleCalenderRequest struct {
	ClientId     string
	ClientSecret string

	calenderID string
	MaxResults int
	OrderBy    string
}

// Sets and returns the Urls parameters for the request
func (c *Credentials) params() url.Values {
	urls := make(url.Values)

	urls.Set("alt", "json")
	urls.Set("maxResult", "5")
	urls.Set("orderBy", "startTime")
	urls.Set("prettyPrint", "false")
	urls.Set("showDeleted", "false")
	urls.Set("singleEvents", "true")
	urls.Set("timeMin", "2019-06-09T20%3A43%3A12%2B02%3A00")
	return urls
}
