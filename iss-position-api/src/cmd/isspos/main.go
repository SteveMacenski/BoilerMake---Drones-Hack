package main

import (
	iss "issposlib"
	"location"
	"log"
	"time"
)

func main() {
	// Connect to Arduino
	/*conf := &goserial.Config{
		Name: "/dev/ttyACM0",
		Baud: 9600,
	}
	_, err := goserial.OpenPort(conf)
	if err != nil {
		panic(err)
	}*/

	// Get my location
	myLocation, err := location.GetMyLocation()
	if err != nil {
		panic(err)
	}

	log.Print("Current location: (",
		myLocation.Latitude,
		",",
		myLocation.Longitude,
		")")

	for {
		// Fetch ISS location
		pos, err := iss.Fetch()
		if err != nil {
			panic(err)
		}

		log.Print("ISS located at (",
			pos.Coordinates.Latitude,
			",",
			pos.Coordinates.Longitude,
			")")

		// TODO write to Arduino

		time.Sleep(60 * time.Second)
	}
}
