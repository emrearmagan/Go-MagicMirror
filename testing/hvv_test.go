package testing

import (
	"Go-MagicMirror/api"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

//---------------------------HVVGetRoute Test-------------------------------
func TestHvvGetRoute(t *testing.T) {

	response := `{"realtimeSchedules":[{"start":{"name":"START"},"dest":{"name":"DESTINATION"},"time":41,"footpathTime":4,"scheduleElements":[{"from":{"name":"START","depTime":{"date":"12.06.2019","time":"13:54"}},"to":{"name":"S Neuwiedenthal","arrTime":{"date":"12.06.2019","time":"13:58"}},"line":{"name":"340","direction":"Ehestorfer Heuweg","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neuwiedenthal","depTime":{"date":"12.06.2019","time":"14:01"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"14:28"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"14:30"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"14:35"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]},{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":50,"footpathTime":13,"scheduleElements":[{"from":{"name":"Moorburger Ring","depTime":{"date":"12.06.2019","time":"14:05"}},"to":{"name":"Rehrstieg","arrTime":{"date":"12.06.2019","time":"14:14"}},"line":{"name":"Fußweg","direction":"","origin":"","type":{"simpleType":"FOOTPATH","shortInfo":""}}},{"from":{"name":"Rehrstieg","depTime":{"date":"12.06.2019","time":"14:14"}},"to":{"name":"S Neuwiedenthal","arrTime":{"date":"12.06.2019","time":"14:16"}},"line":{"name":"251","direction":"Heykenaukamp (Kehre)","origin":"","type":{"simpleType":"BUS","shortInfo":"Bus"}}},{"from":{"name":"Neuwiedenthal","depTime":{"date":"12.06.2019","time":"14:21"}},"to":{"name":"Jungfernstieg","arrTime":{"date":"12.06.2019","time":"14:48"}},"line":{"name":"S3","direction":"Pinneberg","origin":"","type":{"simpleType":"TRAIN","shortInfo":"S"}}},{"from":{"name":"Jungfernstieg","depTime":{"date":"12.06.2019","time":"14:50"}},"to":{"name":"Schlump","arrTime":{"date":"12.06.2019","time":"14:55"}},"line":{"name":"U2","direction":"Niendorf Nord","origin":"","type":{"simpleType":"TRAIN","shortInfo":"U"}}}]}],"returnCode":"OK"}`
	server := testServer(200, response)
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "START"},
		Destinations: api.Station{Name: "DESTINATION"},
		DateTime:     api.DateTime{Date: "12.06.2019", Time: "14:00"},
		Language:     api.GERMAN,
		Amount:       3,
		Apikey:       "randomApiKey",
		Username:     "username",
	}

	resp, err := c.HVVGetRoute(h)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

	if resp.RealtimeSchedules[0].Start.Name != h.Origin.Name {
		fmt.Errorf("returned non nill error, was %s", err)
	}
	if resp.RealtimeSchedules[0].Dest.Name != h.Destinations.Name {
		fmt.Errorf("returned non nill error, was %s", err)
	}
	if resp.RealtimeSchedules[0].Time != 53 {
		fmt.Errorf("returned non nill error, was %s", err)
	}
}

func TestHvvGetRouteMissingOrigin(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		Amount:       3,
		Apikey:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c := api.NewClientWithTestUrl(server.URL)

	if _, err := c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Origin should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingDestinaion(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:   api.Station{Name: "Moorburger Ring"},
		DateTime: api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language: api.GERMAN,
		Amount:   3,
		Apikey:   api.APIKEY_HVV,
		Username: api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c := api.NewClientWithTestUrl(server.URL)

	if _, err := c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Destionation should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingDateTime(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		Language:     api.GERMAN,
		Amount:       3,
		Apikey:       api.APIKEY_HVV,
		Username:     api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c := api.NewClientWithTestUrl(server.URL)

	if _, err := c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing DateTime should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingAPIKEY(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		Amount:       3,
		Username:     api.APIKEY_HVV_USER,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c := api.NewClientWithTestUrl(server.URL)

	if _, err := c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Apikey should've return err, %s", err)
	}
}

func TestHvvGetRouteWrongApi(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		Amount:       3,
		Apikey:       "TEST",
		Username:     api.APIKEY_HVV_USER,
	}

	c := api.NewClient()

	_, err := c.HVVGetRoute(h)
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
		Amount:       3,
		Apikey:       api.APIKEY_HVV,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c := api.NewClientWithTestUrl(server.URL)

	if _, err := c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Username should've return err, %s", err)
	}
}

func TestHvvGetRouteMissingLanguage(t *testing.T) {

	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Amount:       3,
		Apikey:       api.APIKEY_HVV,
	}

	server := testServer(200, `{"status":"OK"}"`)
	c := api.NewClientWithTestUrl(server.URL)

	if _, err := c.HVVGetRoute(h); err == nil {
		t.Errorf("Missing Lanugaune should not return err, %s", err)
	}
}

func TestHvvGetRouteRequestBody(t *testing.T) {
	expectedBody := `{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":{"date":"11.05.2019","time":"14:00"},"language":"de","schedulesAfter":3,"timeIsDeparture":true,"schedulesBefore":0,"realtime":"REALTIME"}`
	var h = &api.HVVGetRouteRequest{
		Origin:       api.Station{Name: "Moorburger Ring"},
		Destinations: api.Station{Name: "Schlump"},
		DateTime:     api.DateTime{Date: "11.05.2019", Time: "14:00"},
		Language:     api.GERMAN,
		Amount:       3,
	}

	reqBody, _ := json.Marshal(h)

	if !reflect.DeepEqual(string(reqBody), expectedBody) {
		t.Errorf("expected %v, was %v", expectedBody, string(reqBody))
	}
}

func TestHvvGetRouteSignature(t *testing.T) {
	reqBody := []byte(`{"start":{"name":"Moorburger Ring"},"dest":{"name":"Schlump"},"time":{"date":"11.05.2019","time":"14:00"},"language":"de","realtime":"REALTIME"}`)
	var h = &api.HVVGetRouteRequest{
		Apikey:   "testSignature",
		Username: "testUsername",
	}

	expectedSignature := "BZpAZcNY1An89aEFGaPkVZsNTMw="
	s := api.ComputeHmac256(reqBody, string(h.Apikey))

	if !reflect.DeepEqual(s, expectedSignature) {
		t.Errorf("expected %v, was %v", expectedSignature, s)
	}

}

//---------------------------HVVDepartureList Test-------------------------------
