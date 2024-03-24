package main

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type Button struct {
	name string
	dev  gpio.PinIO
}

func initializeButton(name string, gpioname string) *Button {
	// Load all the drivers
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Lookup a pin by its number
	p := gpioreg.ByName(gpioname)
	if p == nil {
		log.Fatal("Could not find pin " + gpioname)
	}

	button := Button{name: name, dev: p}
	return &button
}

func (b *Button) monitor(channel *chan string) {
	pin := b.dev
	err := pin.In(gpio.PullDown, gpio.BothEdges)

	if err != nil {
		log.Fatal(err)
	}

	for {
		pin.WaitForEdge(time.Second)
		state := pin.Read()

		*channel <- state.String()
	}
}
