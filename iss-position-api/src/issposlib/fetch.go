package issposlib

import (
	"encoding/json"
	"io/ioutil"
	"location"
	"net/http"
)

const openNotifyUrl = "https://api.wheretheiss.at/v1/satellites/25544"

type ISSPosition struct {
	location.GPSCoords
	Altitude float64 `json:"altitude"`
}

func Fetch() (*ISSPosition, error) {
	resp, err := http.Get(openNotifyUrl)
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pos := &ISSPosition{}
	err = json.Unmarshal([]byte(contents), pos)
	return pos, err
}
