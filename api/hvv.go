package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var hvvConfig = &apiConfig{
	host: "https://api-test.geofox.de",
	path: "/gti/public/getRoute",
}

const (
	REALTIMEON = "REALTIME"
	checkName  = "/gti/public/checkName" //@todo implement checkName
)

func (c *Client) HVVGetRoute(r *HVVGetRouteRequest) (*HVVGetRouteResponse, error) {
	if err := checkHVVParams(r); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("couldn't Marshal HVV-Request")
	}

	signature := ComputeHmac256(reqBody, string(r.Apikey))

	header := map[string]string{"Content-Type": "application/json", "geofox-auth-signature": signature, "geofox-auth-user": r.Username}
	var response struct {
		HVVGetRouteResponse
		HVVCommonResponse
	}

	if err := c.post(hvvConfig, r, &response, reqBody, header); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.HVVGetRouteResponse, nil
}

//@todo not implemented yet
func (c *Client) DepartureList(r *HVVDepartureListRequest) (*HVVDepartureListResponse, error) {
	/*if err := checkHVVParams(r); err != nil {
		return nil, err
	}*/

	hvvConfig.path = "/gti/public/departureList"
	reqBody, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("couldn't marshal HVV-Request")
	}
	signature := ComputeHmac256(reqBody, string(r.ApiKey))

	header := map[string]string{"Content-Type": "application/json", "geofox-auth-signature": signature, "geofox-auth-user": r.Username}
	var response struct {
		HVVDepartureListResponse
		HVVCommonResponse
	}

	if err := c.post(hvvConfig, r, &response, reqBody, header); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}
	return &response.HVVDepartureListResponse, nil
}

//Generates the Signature for the Header
func ComputeHmac256(message []byte, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func checkHVVParams(r *HVVGetRouteRequest) error {
	if len(r.Origin.Name) == 0 {
		return errors.New("origins empty")
	}
	if len(r.Destinations.Name) == 0 {
		return errors.New("destinations empty")
	}
	if len(r.Apikey) == 0 {
		return errors.New("no apiKey selected")
	}
	if len(r.DateTime.Time) == 0 || len(r.DateTime.Date) == 0 {
		return errors.New("no Date or Time selected")
	}
	if len(r.Username) == 0 {
		return errors.New("no username selected")
	}
	//@todo 0 might be okay, should return only 1 route ?
	if r.Amount <= 0 {
		return errors.New("no Amount set")
	}
	if len(r.Language) == 0 {
		r.Language = Language("en")
	}
	return nil
}

// Implements the Json Interface and
// adds additional information to the body when called with json.Marshal(r)
func (r *HVVGetRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		HVVGetRouteRequest
		TimeIsDeparture bool `json:"timeIsDeparture"`
		SchedulesBefore int  `json:"schedulesBefore"`
		// Provides you with Realtime data
		// Optional
		RealTime string `json:"realtime"`
	}{
		HVVGetRouteRequest: HVVGetRouteRequest(*r),
		RealTime:           "REALTIME",
		TimeIsDeparture:    true,
		SchedulesBefore:    0,
	})
}

/* HVVGetRouteRequest is the request struct for HVV GetRoute APi
Example of Request Body
{"start":{"name":"START"},"dest":{"name":"DESTINATION"},"time":{"date":"12.06.2019","time":"14:00"},"language":"de","schedulesAfter":3,"timeIsDeparture":true,"schedulesBefore":0,"realtime":"REALTIME"}
*/
type HVVGetRouteRequest struct {
	// Origins is a list of addresses and/or textual values. (example = Station{Name: "MyStation"})
	// Required.
	Origin Station `json:"start"`
	// Destinations is a list of addresses and/or textual values.(example = Station{Name: "MyDestionation"})
	// to which to calculate distance and time.
	// Required.
	Destinations Station `json:"dest"`
	//Date and Time of the Request. (example = DateTime{Date: "11.06.2019", Time: "14:00"})
	//Required.
	DateTime DateTime `json:"time"`
	// Language in which to return results.
	// Optional. (default is english)
	Language Language `json:"language"`
	// Amount of routes to return.
	// Required
	Amount int `json:"schedulesAfter"`
	// apiKey from HBT
	// Required
	Apikey apiKey `json:"-"`
	// Username from HBT.
	//Required
	Username string `json:"-"`
}

type Station struct {
	Name string `json:"name"`
}
type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

//no URL parameter needed for HBT
func (r *HVVGetRouteRequest) params() url.Values {
	return nil
}
// Expected Response from HBT API
type HVVGetRouteResponse struct {
	RealtimeSchedules []struct {
		Start            Station `json:"start"`
		Dest             Station `json:"dest"`
		Time             float64 `json:"time"`
		FootpathTime     float64 `json:"footpathTime"`
		ScheduleElements []struct {
			From struct {
				Name    string   `json:"name"`
				DepTime DateTime `json:"depTime"`
			} `json:"from,omitempty"`
			To struct {
				Name    string   `json:"name"`
				ArrTime DateTime `json:"arrTime"`
			} `json:"to,omitempty"`
			Line struct {
				BusLine   string `json:"name"`
				Direction string `json:"direction"`
				Origin    string `json:"origin"`
				Type      struct {
					SimpleType string `json:"simpleType"`
					ShortInfo  string `json:"shortInfo"`
				} `json:"type"`
			} `json:"line,omitempty"`
		} `json:"scheduleElements"`
	} `json:"realtimeSchedules"`
}

type HVVCommonResponse struct {
	Status       Status       `json:"returnCode"`
	ErrorMessage ErrorMessage `json:"errorDevInfo,omitempty"`
	ErrorText    string       `json:"errorText,omitempty"`
}

//StatusError returns an error if this object has a Status different
func (c *HVVCommonResponse) StatusError() error {
	if c.Status != "OK" && c.Status != "ok" && c.Status != "200" {
		return fmt.Errorf("maps: %s - %s -%s", c.Status, c.ErrorText, c.ErrorMessage)
	}

	return nil

}

/* HVVDepartureListRequest is the request struct for HVV departureList APi
Example of Request Body
....
*/
type HVVDepartureListRequest struct {
	// Origins is a list of addresses and/or textual values. (example = Station{Name: "MyStation"})
	// Required.
	Origin Station `json:"start"`
	// Destinations is a list of addresses and/or textual values.(example = Station{Name: "MyDestionation"})
	// to which to calculate distance and time.
	// Required.
	Destinations Station `json:"dest"`
	//Date and Time of the Request. (example = DateTime{Date: "11.06.2019", Time: "14:00"})
	//Required.
	DateTime     DateTime `json:"time"`
	MaxList      int      `json:"maxList"`
	ServiceTypes []string `json:"serviceTypes"`
	// Language in which to return results.
	// Optional. (default is english)
	Language Language `json:"language"`
	// Provides you with Realtime data
	RealTime string `json:"realtime"`
	// apiKey from HBT
	// Required
	MaxTimeOffset int    `json:"maxTimeOffset"`
	ApiKey        apiKey `json:"-"`
	Username      string `json:"-"`
}

/*
 Expected Response for DepatureList from HBT API
*/
type HVVDepartureListResponse struct {
	Time struct {
		Date string `json:"date"`
		Time string `json:"time"`
	} `json:"time"`
	Departures []struct {
		Line struct {
			Name      string `json:"name"`
			Direction string `json:"direction"`
			Origin    string `json:"origin"`
			Type      struct {
				SimpleType string `json:"simpleType"`
				ShortInfo  string `json:"shortInfo"`
			} `json:"type"`
		} `json:"line"`
		TimeOffset int `json:"timeOffset"`
	} `json:"departures"`
}

func (r *HVVDepartureListRequest) params() url.Values {
	return nil
}
