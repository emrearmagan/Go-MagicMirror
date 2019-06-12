package api

import (
	"io/ioutil"
	"math/rand"
	"net/http"
)

var NumbersApi = &apiConfig{
	"http://numbersapi.com/",
	"",
}

//Path for different facts
var paths = []string{"random/trivia", "random/year", "random/date", "random/math"}

func (c *Client) Numbers() (string, error) {
	//Selects a random path
	NumbersApi.path = paths[rand.Intn(4-0)+0]

	resp, err := http.Get(NumbersApi.host + NumbersApi.path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}