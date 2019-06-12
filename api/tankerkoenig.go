package api

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

var tankerkoenigApi = &apiConfig{
	host: "https://creativecommons.tankerkoenig.de",
	path: "/json/list.php",
}

func (c *Client) Tankerkoenig(r *TankerkoenigRequest) (*FuelPriceResponse, error) {
	if err := checkTankerkoenigParams(r); err != nil {
		return nil, err
	}

	var response struct {
		FuelPriceResponse
		TankerkoenigCommonResponse
	}

	if err := c.get(tankerkoenigApi, r, &response); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return removeStations(&response.FuelPriceResponse, 5), nil

}

//Only keep a few Stations to list
func removeStations(f *FuelPriceResponse, n int) *FuelPriceResponse {
	removedStations := &FuelPriceResponse{}

	if n >= len(f.Stations) {
		return f
	}

	for i := 0; i < n; i++ {
		removedStations.Stations = append(removedStations.Stations, f.Stations[i])

		//Turn Brand and Street to Title in Format := A Title
		removedStations.Stations[i].Brand = strings.Title(strings.ToLower(removedStations.Stations[i].Brand))
		removedStations.Stations[i].Street = strings.Title(strings.ToLower(removedStations.Stations[i].Street))

	}
	return removedStations
}

// Checks if all required parameter are given. If not return error without starting request
func checkTankerkoenigParams(r *TankerkoenigRequest) error {
	if len(r.Lon) == 0 {
		return errors.New("no longtitude selected")
	}
	if len(r.Lat) == 0 {
		return errors.New("no latiude selected")
	}
	if r.Radius == 0 {
		return errors.New("no radius selected")
	}
	if len(r.GasTyp) == 0 {
		return errors.New("no Gastyp selected - 'e5, e10, diesel or all")
	}
	if r.GasTyp != "all" && len(r.Sortby) == 0 {
		return errors.New("no Sortby selected")
	}
	if len(r.ApiKey) == 0 {
		return errors.New("no apiKey selected")
	}
	return nil
}

/* Example link
https://creativecommons.tankerkoenig.de/json/list.php?lat=52.521&lng=13.438&rad=1.5&sort=dist&type=all&apikey=YOUR_APIKEY
*/
// Sets and returns the Urls parameters for the request
func (r *TankerkoenigRequest) params() url.Values {
	urls := make(url.Values)

	urls.Set("apikey", string(r.ApiKey))
	urls.Set("type", string(r.GasTyp))
	urls.Set("sort", r.Sortby)
	urls.Set("rad", strconv.FormatFloat(r.Radius, 'f', 2, 64))
	urls.Set("lng", string(r.Lon))
	urls.Set("lat", string(r.Lat))

	return urls
}

// TankerkoenigRequest is the request struct for Tankerkoenig API
type TankerkoenigRequest struct {
	// Longitude of City
	// Required.
	Lon Coords
	// Latitude of City
	// Required.
	Lat Coords
	// Radius Specifies the radius to search.
	// Required
	Radius float64
	// Sorts the response by given form
	// Valid values are `SortbyDistance` and `SortbyPrice`.
	// Optional.
	Sortby string
	// Specifiys the Gastyp to search,
	// Valid values are 'GasTypAll', 'GasTypE10', 'GasTypE5' and 'GasTypDiesel'
	// Required
	GasTyp GasTyp
	// apiKey for Tankerkoenig API.
	// Required
	ApiKey apiKey
}

// Expected Response from Tankerkoenig API
type FuelPriceResponse struct {
	Stations []struct {
		Brand       string  `json:"brand"`
		Street      string  `json:"street,omitempty"`
		HouseNumber string  `json:"houseNumber"`
		Price       float64 `json:"price,omitempty"`
		Diesel      float64 `json:"diesel"`
		E5          float64 `json:"e5"`
		E10         float64 `json:"e10"`
		IsOpen      bool    `json:"isOpen"`
	} `json:"stations"`
}

type TankerkoenigCommonResponse struct {
	Status       Status       `json:"status,omitempty"`
	ErrorMessage ErrorMessage `json:"message,omitempty"`
}

//StatusError returns an error if this object has a Status different
func (c *TankerkoenigCommonResponse) StatusError() error {
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
