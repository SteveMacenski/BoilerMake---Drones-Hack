package main

import (
	"coordinates"
	"fmt"
	iss "issposlib"
	"location"
	"log"
	"time"

	"github.com/huin/goserial"
)

func main() {
	// Connect to Arduino
	conf := &goserial.Config{
		Name: "/dev/ttyACM0",
		Baud: 9600,
	}
	port, err := goserial.OpenPort(conf)
	if err != nil {
		panic(err)
	}

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
			pos.Latitude,
			",",
			pos.Longitude,
			")")

		azimuth, inclination := coordinates.ToRelative(pos, myLocation)

		log.Print("Inclination ", inclination, ", azimuth ", azimuth)

		// write to Arduino
		fmt.Fprint(port, inclination)
		fmt.Fprint(port, " ")
		fmt.Fprint(port, azimuth)

		time.Sleep(60 * time.Second)
	}
}
