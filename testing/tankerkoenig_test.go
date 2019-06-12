package testing

import (
	"Go-MagicMirror/api"
	"fmt"
	"strconv"
	"testing"
)

const (
	lat_testTanker    = "52.521"
	lon_testTanker    = "13.438"
	radius_testTanker = 1.5
	sortby_testTanker = "dist"
	gastyp_testTanker = "all"

	apiKey_testTanker = "00000000-0000-0000-0000-000000000002" // Dont worry, this is just a test Api from Tankerkoenig :)
)

func TestTankerkoenig(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Lat:    lon_testTanker,
		Radius: radius_testTanker,
		Sortby: sortby_testTanker,
		GasTyp: gastyp_testTanker,
		ApiKey: apiKey_testTanker,
	}

	c := api.NewClient()

	res, err := c.Tankerkoenig(tk)
	if err != nil {
		t.Errorf("returned non nill error, was %s", err)
	}

	fmt.Println(res)
}

func TestTankerkoenigMissingLon(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lat:    lon_testTanker,
		Radius: radius_testTanker,
		Sortby: sortby_testTanker,
		GasTyp: gastyp_testTanker,
		ApiKey: apiKey_testTanker,
	}

	c := api.NewClient()

	if _, err := c.Tankerkoenig(tk); err == nil {
		t.Errorf("Missing Lon should've return error, %s", err)
	}
}

func TestTankerkoenigMissingLat(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Radius: radius_testTanker,
		Sortby: sortby_testTanker,
		GasTyp: gastyp_testTanker,
		ApiKey: apiKey_testTanker,
	}

	c := api.NewClient()


	if _, err := c.Tankerkoenig(tk); err == nil {
		t.Errorf("Missing Lat should've return error, %s", err)
	}
}

func TestTankerkoenigMissingRadius(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Lat:    lon_testTanker,
		Sortby: sortby_testTanker,
		GasTyp: gastyp_testTanker,
		ApiKey: apiKey_testTanker,
	}

	c := api.NewClient()

	if _, err := c.Tankerkoenig(tk); err == nil {
		t.Errorf("Missing Radius should've return error, %s", err)
	}
}

func TestTankerkoenigMissingGasTyp(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Lat:    lon_testTanker,
		Radius: radius_testTanker,
		Sortby: sortby_testTanker,
		ApiKey: apiKey_testTanker,
	}

	c := api.NewClient()

	if _, err := c.Tankerkoenig(tk); err == nil {
		t.Errorf("Missing GasTyp should've return error, %s", err)
	}
}

//Can only miss "Sortby", when Gastyp is not "all"
func TestTankerkoenigMissingSortby(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Lat:    lon_testTanker,
		Radius: radius_testTanker,
		GasTyp: "diesel",
		ApiKey: apiKey_testTanker,
	}

	c := api.NewClient()


	if _, err := c.Tankerkoenig(tk); err == nil {
		t.Errorf("Missing Sortby should've return error, %s", err)
	}
}

func TestTankerkoenigMissingApiKey(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Lat:    lon_testTanker,
		Radius: radius_testTanker,
		Sortby: sortby_testTanker,
		GasTyp: gastyp_testTanker,
	}

	c := api.NewClient()


	if _, err := c.Tankerkoenig(tk); err == nil {
		t.Errorf("Missing apiKey should've return error, %s", err)
	}
}

func TestTankerkoenigRequestURL(t *testing.T) {
	var tk = &api.TankerkoenigRequest{
		Lon:    lat_testTanker,
		Lat:    lon_testTanker,
		Radius: radius_testTanker,
		Sortby: sortby_testTanker,
		GasTyp: gastyp_testTanker,
		ApiKey: apiKey_testTanker,
	}

	expectedQuery := fmt.Sprintf("apikey=%s&lat=%s&lng=%s&rad=%s&sort=%s&type=%s", tk.ApiKey, tk.Lat, tk.Lon, strconv.FormatFloat(tk.Radius, 'f', 2, 64), tk.Sortby, tk.GasTyp)
	server := testServerForQuery(expectedQuery, 200, `{"status":"OK"}"`)
	defer server.Close()

	c := api.NewClientWithTestUrl(server.URL)

	_, err := c.Tankerkoenig(tk)
	if err != nil {
		t.Errorf("Unexpected error in constructing request URL: %+v", err)
	}
}
