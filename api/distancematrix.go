package api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var distanceMatrixAPI = &ApiConfig{
	"https://maps.googleapis.com",
	"/maps/api/distancematrix/json",
}

func (c *Client) DistanceMatrix(r *DistanceMatrixRequest) (*DistanceMatrixResponse, error) {

	if len(r.Origins) == 0 {
		return nil, errors.New("origins empty")
	}
	if len(r.Destinations) == 0 {
		return nil, errors.New("destinations empty")
	}
	if len(r.ApiKey) == 0 {
		return nil, errors.New("no ApiKey selected")
	}

	var response struct {
		DistanceMatrixCommonResponse
		DistanceMatrixResponse
	}

	if err := c.get(distanceMatrixAPI, r, &response); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.DistanceMatrixResponse, nil

	/*
		var jsonResponses = make(chan []byte)
		var wg sync.WaitGroup

		wg.Add(len(urls))
		for _, url := range urls {
			go func(url string) {
				defer wg.Done()
				resp, err := http.Get(url)
				if err != nil {
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
				}

				jsonResponses <- body
			}(url)
		}

		go func() {
			for response := range jsonResponses {
				var data DistanceMatrixResponse
				_ = json.Unmarshal(response, &data)
				tmp = append(tmp, data)
			}
		}()

		wg.Wait()

		fmt.Println(tmp)

		return nil, nil*/
}

// DistanceMatrixRequest is the request struct for Distance Matrix APi
type DistanceMatrixRequest struct {
	// Origins is a list of addresses and/or textual latitude/longitude values
	// from which to calculate distance and time.
	// Required.
	Origins string
	// Destinations is a list of addresses and/or textual latitude/longitude values
	// to which to calculate distance and time.
	// Required.
	Destinations []string
	// Mode specifies the mode of transport to use when calculating distance.
	// Valid values are `ModeDriving`, `ModeWalking`, `ModeBicycling` and `ModeTransit`.
	// Default is `Modedriving
	// Optional.
	// @Todo implement Mode
	Mode Mode
	// Language in which to return results.
	// Optional.
	Language Language
	// Units Specifies the unit system to use when expressing distance as text.
	// Valid values are `UnitsMetric` and `UnitsImperial`.
	// Optional.
	Units Units
	//Restrictions are indicated by use of the avoid parameter, and an argument to that parameter indicating the restriction to avoid.
	// Valid values are `AvoidTolls`, 'AvoidHighways`, 'AvoidFerries` and `AvoidIndoor`
	//Optional
	Avoid string
	// ApiKey from google.developers.
	// Required
	ApiKey ApiKey
}

/*Example link
https://maps.googleapis.com/maps/api/distancematrix/json?units=imperial&origins=Washington,DC&destinations=New+York+City,NY&key=YOUR_API_KEY
*/
// Sets and returns the Urls parameters for the request
func (r *DistanceMatrixRequest) params() url.Values {
	urls := make(url.Values)
	if r.Units != "" {
		urls.Set("units", string(r.Units))
	} else {
		urls.Set("units", "metric")
	}
	urls.Set("origins", r.Origins)
	urls.Set("destinations", strings.Join(r.Destinations, "|"))
	if r.Mode != "" {
		urls.Set("mode", string(r.Mode))
	}
	if r.Language != "" {
		urls.Set("language", string(r.Language))
	}

	urls.Set("key", string(r.ApiKey))

	return urls
}

//----------------------------------------------Response------------------------------------------

// Expected Response from google API
type DistanceMatrixResponse struct {
	DestinationAddresses []string                    `json:"destination_addresses"`
	OriginAddresses      []string                    `json:"origin_addresses"`
	Rows                 []DistanceMatrixElementsRow `json:"rows"`
}

// DistanceMatrixElementsRow is a row of distance elements.
type DistanceMatrixElementsRow struct {
	Elements []DistanceMatrixElement `json:"elements"`
}

// DistanceMatrixElement is the travel distance and time for a pair of origin
// and destination.
type DistanceMatrixElement struct {
	Distance Time   `json:"distance"`
	Duration Time   `json:"duration"`
	Status   string `json:"status"`
}

type Time struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type DistanceMatrixCommonResponse struct {
	//ErrorMessage1 for Google DistanceMatrix
	Status       Status       `json:"status,omitempty"`
	ErrorMessage ErrorMessage `json:"error_message,omitempty"`
}

//StatusError returns an error if this object has a Status different
func (c *DistanceMatrixCommonResponse) StatusError() error {
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
		if status != "OK" && c.Status != "ok" && status != "ZERO_RESULTS" && c.Status != "200" {
			return fmt.Errorf("maps: %s - %s", c.Status, c.ErrorMessage)
		}
	}
	return nil
}
