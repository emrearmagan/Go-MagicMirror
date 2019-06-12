package testing

import (
	"Go-MagicMirror/api"
	"fmt"
	"reflect"
	"testing"
)

func TestDistanceMatrixHamburgToBerlin(t *testing.T) {
	//expected response
	response := `{
		"destination_addresses" : [
			 "Berlin, Germany"
		],
		"origin_addresses" : [
			 "Hamburg, Germany"
		],
		"rows" : [
			 {
					"elements" : [
						 {
								"distance" : {
									 "text" : "291 km",
									 "value" : 290607
								},
								"duration" : {
									 "text" : "3 hours 10 mins",
									 "value" : 11412
								},
								"status" : "OK"
						 }
					]
			 }
		],
		"status" : "OK"
}`

	server := testServer(200, response)
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)

	r := &api.DistanceMatrixRequest{
		Origins:      "Hamburg, Germany",
		Destinations: []string{"Berlin, Germany"},
		Units:        api.UnitsMetric,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	resp, err := c.DistanceMatrix(r)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

	fmt.Println(resp)
	fmt.Println(reflect.TypeOf(resp))
	if resp.OriginAddresses[0] != r.Origins {
		t.Errorf("Incorrect origin address")
	}
	if resp.DestinationAddresses[0] != r.Destinations[0] {
		t.Error("Incorrect destination address")
	}
	if resp.Rows[0].Elements[0].Distance.Text != "291 km" {
		t.Error("Incorrect human readable distance")
	}
	if resp.Rows[0].Elements[0].Distance.Value != 290607 {
		t.Error("Incorrect distance value")
	}
	if resp.Rows[0].Elements[0].Duration.Value != 11412 {
		t.Errorf("Incorrect Duration: %v", resp.Rows[0].Elements[0].Duration)
	}
	if resp.Rows[0].Elements[0].Status != "OK" {
		t.Error("Incorrect element status")
	}
}

func TestDistanceMatrixWithCoords(t *testing.T) {
	//expected response
	response := `{
		"destination_addresses" : [
			 "Berlin, Germany"
		],
		"origin_addresses" : [
			 "Hamburg, Germany"
		],
		"rows" : [
			 {
					"elements" : [
						 {
								"distance" : {
									 "text" : "291 km",
									 "value" : 290607
								},
								"duration" : {
									 "text" : "3 hours 10 mins",
									 "value" : 11412
								},
								"status" : "OK"
						 }
					]
			 }
		],
		"status" : "OK"
}`

	server := testServer(200, response)
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)

	r := &api.DistanceMatrixRequest{
		Origins:      "9.993682,53.551085",
		Destinations: []string{"52.520008, 13.404954"},
		Units:        api.UnitsImperial,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	resp, err := c.DistanceMatrix(r)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

	correctResponse := &api.DistanceMatrixResponse{
		OriginAddresses:      []string{"Hamburg, Germany"},
		DestinationAddresses: []string{"Berlin, Germany"},
		Rows: []api.DistanceMatrixElementsRow{
			{
				Elements: []api.DistanceMatrixElement{
					{
						Distance: api.Time{Text: "291 km", Value: 290607},
						Duration: api.Time{Text: "3 hours 10 mins", Value: 11412},
						Status:   "OK",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(resp, correctResponse) {
		t.Errorf("expected %v, was %v", correctResponse, resp)
	}

}

func TestDistanceMatrixMissingOrigins(t *testing.T) {
	server := testServer(200, "")
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)
	r := &api.DistanceMatrixRequest{
		Origins:      "",
		Destinations: []string{"53.566569,9.98464"},
		Units:        api.UnitsImperial,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	if _, err := c.DistanceMatrix(r); err == nil {
		t.Errorf("Missing Origins should've return error")
	}
}

func TestDistanceMatrixMissingDestinations(t *testing.T) {
	server := testServer(200, "")
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)
	r := &api.DistanceMatrixRequest{
		Origins:      "53.561933, 9.953944",
		Destinations: []string{},
		Units:        api.UnitsImperial,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	if _, err := c.DistanceMatrix(r); err == nil {
		t.Errorf("Missing Destinations should've return error")
	}
}

func TestDistanceMatrixMissingApiKey(t *testing.T) {
	c := api.NewClient()
	r := &api.DistanceMatrixRequest{
		Origins:      "53.561933, 9.953944",
		Destinations: []string{"53.566569,9.98464"},
		Units:        api.UnitsImperial,
	}

	if _, err := c.DistanceMatrix(r); err == nil {
		t.Errorf("Missing apiKey should've return error")
	}
}

func TestDistanceMatrixTransitRequestURL(t *testing.T) {
	expectedQuery := "destinations=Berlin%2C+Germany&key=" + string(api.APIKEY_DISTANCEMATRIX) + "&language=de&origins=Hamburg%2C+Germany&units=metric"
	server := testServerForQuery(expectedQuery, 200, `{"status":"OK"}"`)
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)

	r := &api.DistanceMatrixRequest{
		Origins:      "Hamburg, Germany",
		Destinations: []string{"Berlin, Germany"},
		Language:     api.GERMAN,
		Units:        api.UnitsMetric,
		ApiKey:       api.APIKEY_DISTANCEMATRIX,
	}

	_, err := c.DistanceMatrix(r)
	if err != nil {
		t.Errorf("Unexpected error in constructing request URL: %+v", err)
	}
}
