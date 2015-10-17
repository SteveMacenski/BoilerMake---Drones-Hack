package issposlib

import (
	"encoding/json"
	"io/ioutil"
	"location"
	"net/http"
)

const openNotifyUrl = "http://api.open-notify.org/iss-now.json"

type ISSPosition struct {
	Message string `json:"message"`
	//Timestamp   time.Time `json:"timestamp"`
	Coordinates location.GPSCoords `json:"iss_position"`
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
