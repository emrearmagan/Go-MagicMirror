package api

import (
	"errors"
	"fmt"
	"net/url"
)

var openWeatherApi = &ApiConfig{
	host: "http://api.openweathermap.org",
}

// Returns the current Weather of given Lon and Lat
func (c *Client) OpenWeather(r *OpenWeatherRequest) (*WeatherResponse, error) {
	openWeatherApi.path = "/data/2.5/weather"

	if err := checkWeatherParams(r); err != nil {
		return nil, err
	}

	var response struct {
		WeatherResponse
		OpenWeatherCommonResponse
	}

	if err := c.get(openWeatherApi, r, &response); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.WeatherResponse, nil
}

// Returns weather forecast for 5 days with data every hours of given Lon and Lat
func (c *Client) OpenForecast(r *OpenWeatherRequest) (*ForecastResponse, error) {
	openWeatherApi.path = "/data/2.5/forecast"

	if err := checkWeatherParams(r); err != nil {
		return nil, err
	}

	var response struct {
		ForecastResponse
		OpenWeatherCommonResponse
	}

	if err := c.get(openWeatherApi, r, &response); err != nil {
		return nil, err
	}

	if err := response.StatusError(); err != nil {
		return nil, err
	}

	return &response.ForecastResponse, nil

}

// Checks if all required parameter are given. If not return error without starting request
func checkWeatherParams(r *OpenWeatherRequest) error {
	if len(r.Lon) == 0 {
		return errors.New("longtitude empty")
	}
	if len(r.Lat) == 0 {
		return errors.New("latiude empty")
	}
	if len(r.ApiKey) == 0 {
		return errors.New("no ApiKey selected")
	}

	return nil
}

/* Example link
https://samples.openweathermap.org/data/2.5/weather?lat=35&lon=139&appid=YOUR_APIKEY
*/
// Sets and returns the Urls parameters for the request
func (r *OpenWeatherRequest) params() url.Values {
	urls := make(url.Values)
	if r.Units != "" {
		urls.Set("units", string(r.Units))
	}
	urls.Set("lat", string(r.Lat))
	urls.Set("lon", string(r.Lon))
	urls.Set("appid", string(r.ApiKey))

	return urls
}

// OpenWeatherRequest is the request struct for OpenWeather APi
type OpenWeatherRequest struct {
	// Longitude of City
	// Required.
	Lon Coords
	// Latitude of City
	// Required.
	Lat Coords
	// Units Specifies the unit system to use when expressing distance as text.
	// Valid values are `UnitsMetric` and `UnitsImperial`.
	// Optional.
	Units Units
	// ApiKey for OpenWeather API.
	// Required
	ApiKey ApiKey
}

//----------------------------------------------Response------------------------------------------
//@Todo combine Forecast and Weather in one struct and make code simpler
// Expected Response from OpenWeather API
type ForecastResponse struct {
	List []struct {
		Time int `json:"dt"` //time of date forecasted, unix, UTC
		Main struct {
			Temp     float64 `json:"temp"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
			Humidity int     `json:"humidity"` // in %
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"` //cloudiness in %
		} `json:"clouds"`
	} `json:"list"`
}

type WeatherResponse struct {
	Name    string `json:"name"`
	Time    int    `json:"dt"` //time of date, unix, UTC
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity int     `json:"humidity"` //in %
	} `json:"main"`
	Clouds struct {
		All int `json:"all"` //cloudiness in %
	} `json:"clouds"`
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`
}

type OpenWeatherCommonResponse struct {
	Status       Status       `json:"cod"` //
	ErrorMessage ErrorMessage `json:"message,omitempty"`
}

//StatusError returns an error if this object has a Status different
func (c *OpenWeatherCommonResponse) StatusError() error {
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
