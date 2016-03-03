package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	TEMPLATE = `
{
  "date8601": "2015-07-06T12:34:45Z",
  "dateyymmdd": "150706",
  "dateyyyymmdd": "20150706"
}
`
)

type Dates struct {
	//DateISO8601  DateISO8601  `json:"date8601"`
	DateYYMMDD   DateYYMMDD   `json:"dateyymmdd"`
	DateYYYYMMDD DateYYYYMMDD `json:"dateyyyymmdd"`
}

func TestDate(t *testing.T) {
	var d Dates

	if err := json.Unmarshal([]byte(TEMPLATE), &d); err != nil {
		t.Fatal(err)
	}

	str, err := json.MarshalIndent(&d, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if string(str) != TEMPLATE {
		t.Fatal(fmt.Errorf("Input and output does not match %s %v", str, TEMPLATE))
	}
}
