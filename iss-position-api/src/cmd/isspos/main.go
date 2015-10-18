package main

import (
	"bufio"
	"coordinates"
	"fmt"
	iss "issposlib"
	"location"
	"log"
	"time"

	goserial "github.com/tarm/serial"
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
	bufreader := bufio.NewReader(port)
	time.Sleep(5 * time.Second)

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

		log.Printf("%.2f %.2f\n", inclination, azimuth)

		// write to Arduino
		fmt.Fprintf(port, "%.2f %.2f\n", inclination, azimuth)
		port.Flush()

		line, _ := bufreader.ReadString('\n')

		log.Print("Read \"", string(line), "\" from serial")

		time.Sleep(60 * time.Second)
	}
}
