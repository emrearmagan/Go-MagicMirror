package testing

import (
	"Go-MagicMirror/api"
	"fmt"
	"testing"
)

const (
	lat_testOpenWeather    = "35"
	lon_testOpenweather    = "139"
	apiKey_testOpenweather = "b6907d289e10d714a6e88b30761fae22" // Dont worry, this is just a test Api from OpenWeather :)
	URL_testOpenWeather = "https://samples.openweathermap.org"
)

//Test server works only without units
func TestOpenWeather(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Lat:    lat_testOpenWeather,
		ApiKey: apiKey_testOpenweather,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)

	res, err := c.OpenWeather(r)
	if err != nil {
		t.Errorf("Unexpected error in constructing request URL: %+v", err)
	}

	fmt.Println(res)
}

func TestOpenForecast(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Lat:    lat_testOpenWeather,
		ApiKey: apiKey_testOpenweather,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)

	res, err := c.OpenForecast(r)
	if err != nil {
		t.Errorf("Unexpected error in constructing request URL: %+v", err)
	}

	fmt.Println(res)

}

func TestOpenweatherMissingLon(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lat:    lat_testOpenWeather,
		Units:  api.UnitsMetric,
		ApiKey: apiKey_testOpenweather,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)

	if _, err := c.OpenWeather(r); err == nil {
		t.Errorf("Missing Lon should've return error, %s", err)
	}
}

func TestOpenWeatherMissingLat(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Units:  api.UnitsMetric,
		ApiKey: apiKey_testOpenweather,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)


	if _, err := c.OpenWeather(r); err == nil {
		t.Errorf("Missing Lat should've return error, %s", err)
	}
}

func TestOpenForecastMissingLon(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lat:    lat_testOpenWeather,
		Units:  api.UnitsMetric,
		ApiKey: apiKey_testOpenweather,
	}

	c:= api.NewClientWithTestUrl(URL_testOpenWeather)


	if _, err := c.OpenForecast(r); err == nil {
		t.Errorf("Missing Lon should've return error, %s", err)
	}
}

func TestOpenForecastMissingLat(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Units:  api.UnitsMetric,
		ApiKey: apiKey_testOpenweather,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)

	if _, err := c.OpenForecast(r); err == nil {
		t.Errorf("Missing Lat should've return error, %s", err)
	}
}

func TestOpenWeatherMissingUnit(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Lat:    lat_testOpenWeather,
		ApiKey: apiKey_testOpenweather,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)

	if _, err := c.OpenWeather(r); err != nil {
		t.Errorf("Missing unit should not return error, %s", err)
	}
}

func TestOpenWeatherMissingApiKey(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:   lon_testOpenweather,
		Lat:   lat_testOpenWeather,
		Units: api.UnitsMetric,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)

	if _, err := c.OpenWeather(r); err == nil {
		t.Errorf("Missing apiKey should've return error, %s", err)
	}
}

func TestOpenForecastMissingApiKey(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:   lon_testOpenweather,
		Lat:   lat_testOpenWeather,
		Units: api.UnitsMetric,
	}

	c := api.NewClientWithTestUrl(URL_testOpenWeather)


	if _, err := c.OpenForecast(r); err == nil {
		t.Errorf("Missing apiKey should've return error, %s", err)
	}
}

func TestOpenWeatherRequestURL(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Lat:    lat_testOpenWeather,
		Units:  api.UnitsMetric,
		ApiKey: apiKey_testOpenweather,
	}

	expectedQuery := fmt.Sprintf("appid=%s&lat=%s&lon=%s&units=%s", string(r.ApiKey), r.Lat, r.Lon, r.Units)
	server := testServerForQuery(expectedQuery, 200, `{"status":"OK"}"`)
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)

	_, err := c.OpenWeather(r)
	if err != nil {
		t.Errorf("Unexpected error in constructing request URL: %+v", err)
	}
}
