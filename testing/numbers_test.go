package testing

import (
	"Go-MagicMirror/api"
	"fmt"
	"testing"
)

func TestNumbers(t *testing.T) {
	c := api.NewClient()
	resp, err := c.Numbers()
	if err != nil {
		t.Errorf("returned non nill error: %s", err.Error())
	}

	fmt.Println(resp)
}
