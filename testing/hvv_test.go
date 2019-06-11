package testing

import (
	"Go-MagicMirror/api"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

//---------------------------HVVGetRoute Test-------------------------------
func TestHvvGetRoute(t *testing.T) {

	response := `{"realtimeSchedules":[{"Start":{"Name":"Moorburger Ring"}"Dest":{"Name":"Schlump"}"Time":53 "FootpathTime":6 "ScheduleElements":[{"From":{"Name":"Moorburger Ring" "DepTime":{"Date":11.05.2019 "Time":13:07}} To:{"Name":"S Neugraben" "ArrTime":{"Date":11.05.2019 "Time":13:14}} "Line":{"BusLine":340 "Direction":"S Neugraben" "Origin": "Type":{"SimpleType":"BUS" "ShortInfo":"Bus"}}}}`
	server := testServer(200, response)
	defer server.Close()

	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		log.Fatal(err)
	}

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		APIKEY:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	r, err := c.HVVGetRoute(h)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

	fmt.Println(r)
	////@Todo returns a string and not type of HVVResponse like in disntancematrix test
	//fmt.Println(len(resp.RealtimeSchedules))
	//if resp.RealtimeSchedules[0].Start.Name != h.Origin.Name {
	//	testGoogleCal.Errorf("returned non nill error, was %s", err)
	//}
	//if resp.RealtimeSchedules[0].Dest.Name != h.Destinations.Name {
	//	testGoogleCal.Errorf("returned non nill error, was %s", err)
	//}
	//if resp.RealtimeSchedules[0].Time != 53 {
	//	testGoogleCal.Errorf("returned non nill error, was %s", err)
	//}
}

func TestHvvGetRouteMissingOrigin(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		APIKEY:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		t.Error(err)
	}

	if _, err = c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Origin should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingDestinaion(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:   api.Station{Name: "Moorburger Ring"},
		DateTime: api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language: api.GERMAN,
		RealTime: api.REALTIMEON,
		APIKEY:   api.APIKEY_HVV,
		Username: api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		t.Error(err)
	}

	if _, err = c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Destionation should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingDateTime(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		APIKEY:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		t.Error(err)
	}

	if _, err = c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing DateTime should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingAPIKEY(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		Username:     api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		t.Error(err)
	}

	if _, err = c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing APIKEY should've return err, %s", err)
	}
}

func TestHvvGetRouteWrongApi(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		APIKEY:       "TEST",
		Username:     api.APIKEY_HVV_USER,
	}

	c, err := api.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	_, err = c.HVVGetRoute(h)
	if err == nil {
		t.Errorf("returned non nill error, was %s", err)
	}
}

func TestHvvGetRouteMissingUsername(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
		APIKEY:       api.APIKEY_HVV,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		t.Error(err)
	}

	if _, err = c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Username should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingLanguage(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		RealTime:     api.REALTIMEON,
		APIKEY:       api.APIKEY_HVV,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c, err := api.NewClientWithTestUrl(server.URL)
	if err != nil {
		t.Error(err)
	}

	if _, err = c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Lanugaune should not return err, %s", err)
	}
}

func TestHvvGetRouteRequestBody(t *testing.T) {
	expectedBody := `{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":{"date":"11.05.2019","time":"14:00"},"language":"de","realtime":"REALTIME"}`
	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		RealTime:     api.REALTIMEON,
	}

	reqBody, _ := json.Marshal(h)

	if !reflect.DeepEqual(string(reqBody), expectedBody) {
		t.Errorf("expected %v, was %v", expectedBody, string(reqBody))
	}
}

func TestHvvGetRouteSignature(t *testing.T) {
	reqBody := []byte(`{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":{"date":"11.05.2019","time":"14:00"},"language":"de","realtime":"REALTIME"}`)
	var h = &api.HVVGetRouteRequest{
		APIKEY:   "testSignature",
		Username: api.APIKEY_HVV_USER,
	}

	expectedSignature := "BZpAZcNY1An89aEFGaPkVZsNTMw="
	s := api.ComputeHmac256(reqBody, string(h.APIKEY))

	if !reflect.DeepEqual(s, expectedSignature) {
		t.Errorf("expected %v, was %v", expectedSignature, s)
	}
}

//---------------------------HVVDepartureList Test-------------------------------
