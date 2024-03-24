package main

import (
	"image"
	"log"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/host/v3"
)

type Display struct {
	name string
	dev  *ssd1306.Dev
}

func initializeDisplay() *Display {
	// Load all the drivers
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Open a handle to the first available I2C bus
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	// Open a handle to a ssd1306 device on the bus
	dev, err := ssd1306.NewI2C(bus, &ssd1306.DefaultOpts)
	if err != nil {
		log.Fatal(err)
	}

	display := Display{name: "OLED", dev: dev}
	return &display
}

func (d *Display) show(img image.Image) {
	d.dev.Draw(img.Bounds(), img, image.Point{})
}
