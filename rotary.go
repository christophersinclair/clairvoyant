package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type Rotary struct {
	name string
	sw   gpio.PinIO
	clk  gpio.PinIO
	dt   gpio.PinIO
}

const DEBOUNCE_DURATION = 5 * time.Millisecond

func initializeRotary(swPin string, aPin string, bPin string) *Rotary {
	// Load all the drivers
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Lookup switch pin by its number
	swP := gpioreg.ByName(swPin)
	if swP == nil {
		log.Fatal("Could not find pin " + swPin)
	}

	// Lookup A pin (CLK) by its number
	aP := gpioreg.ByName(aPin)
	if aP == nil {
		log.Fatal("Could not find pin " + aPin)
	}

	// Lookup B pin (DT) by its number
	bP := gpioreg.ByName(bPin)
	if bP == nil {
		log.Fatal("Could not find pin " + bPin)
	}

	rot := Rotary{name: "Rotary", sw: swP, clk: aP, dt: bP}
	return &rot
}

func (r *Rotary) monitor(channel *chan string) {
	pinA := r.clk
	pinB := r.dt

	pinA.In(gpio.PullUp, gpio.BothEdges) // default state for both pins is high, so pull up
	pinB.In(gpio.PullUp, gpio.BothEdges) // default state for both pins is high, so pull up

	for {
		if pinA.WaitForEdge(time.Second) {
			fmt.Println("Detected rotation")

			// Debounce
			time.Sleep(DEBOUNCE_DURATION)

			// Read state
			stateA := pinA.Read()
			stateB := pinB.Read()

			if stateA != stateB {
				// Knob turned clockwise
				// Update channel with clockwise rotation event
				*channel <- "Clockwise"
			}

			if stateA == stateB {
				// Knob turned counterclockwise
				// Update channel with counterclockwise rotation event
				*channel <- "Counterclockwise"
			}
		}
		// You might want to include debouncing logic
	}
}
