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

var hvvAPI = &ApiConfig{
	host: "https://api-test.geofox.de",
	path: "/gti/public/getRoute",
}

const (
	REALTIMEON  = "REALTIME"
	checkName = "/gti/public/checkName"	//@todo implement checkName
)

func (c *Client) HVVGetRoute(r *HVVGetRouteRequest) (*HVVGetRouteResponse, error) {
	if err := checkHVVParams(r); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("couldn't Marshal HVV-Request")
	}

	signature := ComputeHmac256(reqBody, string(r.APIKEY))

	header := map[string]string{"Content-Type": "application/json", "geofox-auth-signature": signature, "geofox-auth-user": r.Username}
	var response struct {
		HVVGetRouteResponse
		HVVCommonResponse
	}

	if err := c.post(hvvAPI, r, &response, reqBody, header); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.HVVGetRouteResponse, nil
}

//@todo implement DepatureList for Alex echo
func (c *Client) DepartureList(r *HVVDepartureListRequest) (*HVVDepartureListResponse, error) {
	/*if err := checkHVVParams(r); err != nil {
		return nil, err
	}*/

	hvvAPI.path = "/gti/public/departureList"
	reqBody, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("couldn'testGoogleCal Marshal HVV-Request")
	}
	signature := ComputeHmac256(reqBody, string(r.APIKEY))

	header := map[string]string{"Content-Type": "application/json", "geofox-auth-signature": signature, "geofox-auth-user": r.Username}
	var response struct {
		HVVDepartureListResponse
		HVVCommonResponse
	}

	if err := c.post(hvvAPI, r, &response, reqBody, header); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}
	return &response.HVVDepartureListResponse, nil
}

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
	if len(r.APIKEY) == 0 {
		return errors.New("no ApiKey selected")
	}
	if len(r.DateTime.Time) == 0 || len(r.DateTime.Date) == 0 {
		return errors.New("no Date or Time selected")
	}
	if len(r.Username) == 0 {
		return errors.New("no username selected")
	}
	return nil
}

/* HVVDepartureListRequest is the request struct for HVV GetRoute APi
		Example of Request
		{"version":30,"station":{"id":"Master:41905","name":"Test Origin","type":"STATION"},"time":{"date":"13.05.2019","time":"14:48"},"maxList":30,"serviceTypes":["BUS","ZUG","FAEHRE"],"useRealtime":true,"maxTimeOffset":120}
 */
type HVVDepartureListRequest struct {
	/*	Version int `json:"version"`
		Station struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"station"`
		Time struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"time"`
		MaxList       int      `json:"maxList"`
		ServiceTypes  []string `json:"serviceTypes"`
		UseRealtime   bool     `json:"useRealtime"`
		MaxTimeOffset int      `json:"maxTimeOffset"`*/

	// Origins is a list of addresses and/or textual latitude/longitude values
	// from which to calculate distance and time.
	// Required.
	Origin Station `json:"start"`
	// Destinations is a list of addresses and/or textual latitude/longitude values
	// to which to calculate distance and time.
	// Required.
	DateTime     DateTime `json:"time"`
	MaxList      int      `json:"maxList"`
	ServiceTypes []string `json:"serviceTypes"`
	// Language in which to return results.
	// Optional.
	//Language Language `json:"language"`
	// Provides you with Realtime data
	// Optional (haven'testGoogleCal tried yet)
	RealTime string `json:"realtime"`
	// ApiKey from google.developers.
	// Required
	MaxTimeOffset int    `json:"maxTimeOffset"`
	APIKEY        ApiKey `json:"-"`
	Username      string `json:"-"`
}

/*
 Expected Response from HBT API
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

/* HVVGetRouteRequest is the request struct for HVV GetRoute APi
	Example of Request
		{"version":36,"language":"de","start":{"name":"Test Origin"},"dest":{"name":"Test Destination"},"time":{"date":"11.05.2019","time":"14:00"},"realtime":"REALTIME"}
*/
type HVVGetRouteRequest struct {
	// Origins is a list of addresses and/or textual latitude/longitude values
	// from which to calculate distance and time.
	// Required.
	Origin Station `json:"start"`
	// Destinations is a list of addresses and/or textual latitude/longitude values
	// to which to calculate distance and time.
	// Required.
	Destinations Station  `json:"dest"`
	DateTime     DateTime `json:"time"`
	// Language in which to return results.
	// Optional.
	Language Language `json:"language"`
	// Provides you with Realtime data
	// Optional (haven'testGoogleCal tried yet)
	RealTime string `json:"realtime"`
	// ApiKey from google.developers.
	// Required
	APIKEY   ApiKey `json:"-"`
	Username string `json:"-"`
}

type Station struct {
	Name string `json:"name"`
}
type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

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
	ErrorMessage ErrorMessage `json:"errorDevInfo"`
}

//StatusError returns an error if this object has a Status different
func (c *HVVCommonResponse) StatusError() error {
	fmt.Println(c.Status)
	fmt.Println(c.ErrorMessage)
	switch status := c.Status.(type) {
	case int:
		if status != 200 {
			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
		}
	case float64:
		if status != 200 {
			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
		}
	case string:
		if status != "OK" && c.Status != "ok" && c.Status != "200" {
			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
		}
	}
	return nil
}
