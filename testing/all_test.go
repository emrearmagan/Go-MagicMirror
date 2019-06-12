package testing

import (
	"Go-MagicMirror/api"
	"encoding/json"
	"fmt"
	"testing"
)

/*
	This is a proper request to the API Server without the Testserver. Only for testing purposes.
	Its commented out, since we don't want to make a request every time we test.
	We don't want want our Apikey to reach the limit :)
*/

func TestHVV(t *testing.T) {
	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Pinneberg"},
		DateTime:     api.DateTime{Date: "12.06.2019", Time: "14:00"},
		Language:     api.GERMAN,
		Amount:       1,
		Apikey:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	c := api.NewClient()

	res, err := c.HVVGetRoute(h)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", res)
}

func TestDeparture(t *testing.T) {
	var h = &api.HVVDepartureListRequest{
		Origin:        api.Station{Name: "Neuwiedenthal"},
		DateTime:      api.DateTime{Date: "11.06.2019", Time: "14:00"},
		MaxList:       30,
		RealTime:      api.REALTIMEON,
		MaxTimeOffset: 120,
		ServiceTypes:  []string{"BUS", "ZUG", "FAEHRE"},
		ApiKey:        api.APIKEY_HVV,
		Username:      api.APIKEY_HVV_USER,
	}

	c := api.NewClient()

	res, err := c.DepartureList(h)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

	fmt.Println(res)
}

func TestOpen(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    api.Longitude,
		Lat:    api.Latitude,
		Units:  "metric",
		ApiKey: api.APIKEY_OPENWEATHER,
	}

	c := api.NewClient()

	res, err := c.OpenWeather(r)
	if err != nil {
		t.Error(err)
	}


	fmt.Println(res)
}

func TestForecast(t *testing.T) {
	var r = &api.OpenWeatherRequest{
		Lon:    api.Longitude,
		Lat:    api.Latitude,
		Units:  "metric",
		ApiKey: api.APIKEY_OPENWEATHER,
	}

	c := api.NewClient()

	res, err := c.OpenForecast(r)
	if err != nil {
		t.Error(err)
	}

	s, _ := json.Marshal(res)
	fmt.Println(string(s))
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

	c := api.NewClient()

	res, err := c.Tankerkoenig(tk)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestGoogleDistanceMatrix(t *testing.T) {
	g := &api.DistanceMatrixRequest{
		Origins:      "Hamburg, Germany",
		Destinations: []string{"Berlin, Germany", "Bremen, Germany", "Frankfurt, Germany"},
		Units:        api.UnitsMetric,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	c := api.NewClient()

	res, err := c.DistanceMatrix(g)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestGoogleCalender(t *testing.T) {
	c := api.NewClient()

	var request api.GoogleCalenderRequest

	d, err := c.GoogleCalender(&request)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(d)
}
