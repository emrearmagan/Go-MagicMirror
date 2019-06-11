package testing

import (
	"Go-MagicMirror/api"
	"fmt"
	"log"
	"testing"
)

/*
	This is a proper request to the API Server without the Testserver. Only for testing purposes.
	Its commented out, since we don't want to make a request every time we test.
	We don't want want our APIKEY to reach the limit :)
*/

func TestHVV(t *testing.T) {
	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Rehrstieg"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		APIKEY:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	c, err := api.NewClient()
	if err != nil {
		t.Error(err)
	}

	res, err := c.HVVGetRoute(h)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", res)
}

func TestDeparture(t *testing.T) {

	var h = &api.HVVDepartureListRequest{
		Origin:        api.Station{Name: "Rehrstieg"},
		DateTime:      api.DateTime{Date: "11.05.2019", Time: "14:00"},
		MaxList:       30,
		RealTime:      api.REALTIMEON,
		MaxTimeOffset: 120,
		ServiceTypes:  []string{"BUS", "ZUG", "FAEHRE"},
		APIKEY:        api.APIKEY_HVV,
		Username:      api.APIKEY_HVV_USER,
	}

	c, err := api.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.DepartureList(h)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

}

func TestOpen(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Lat:    lat_testOpenWeather,
		ApiKey: api.APIKEY_OPENWEATHER,
	}

	c, err := api.NewClient()
	if err != nil {
		t.Error(err)
	}

	res, err := c.OpenWeather(r)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestForecast(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    lon_testOpenweather,
		Lat:    lat_testOpenWeather,
		ApiKey: api.APIKEY_OPENWEATHER,
	}

	c, err := api.NewClient()
	if err != nil {
		t.Error(err)
	}

	res, err := c.OpenForecast(r)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestTanker(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    api.Longitude,
		Lat:    api.Latitude,
		Radius: api.Radius,
		Sortby: api.SortbyDistance,
		GasTyp: api.GasTypAll,
		ApiKey: api.APIKEY_TANKERKOENIG,
	}

	c, err := api.NewClient()
	if err != nil {
		t.Error(err)
	}

	res, err := c.Tankerkoenig(tk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(*res)
}

func TestGoogleDistanceMatrix(t *testing.T) {
	g := &api.DistanceMatrixRequest{
		Origins:      "Hamburg, Germany",
		Destinations: []string{"Berlin, Germany"},
		Units:        api.UnitsMetric,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	c, err := api.NewClient()
	if err != nil {
		t.Error(err)
	}

	res, err := c.DistanceMatrix(g)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestGoogleCalender(t *testing.T) {
	c, err := api.NewClient()
	if err != nil {
		t.Error(err)
	}

	var request api.GoogleCalenderRequest

	d, err := c.GoogleCalender(&request)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(d)
}
