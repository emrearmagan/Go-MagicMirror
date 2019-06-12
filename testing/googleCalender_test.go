package testing

import (
	"Go-MagicMirror/api"
	"fmt"
	"testing"
)

func TestCalender(t *testing.T) {
	c := api.NewClient()

	var request api.GoogleCalenderRequest

	d, err := c.GoogleCalender(&request)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(d)
}
