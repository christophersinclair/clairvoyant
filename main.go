package main

import (
	"math/rand"
	"strconv"
)

func main() {
start:
	GAME_START := false
	NUM_PLAYERS := 0

	hardwareEventChannel := make(chan string)
	setupInput(&hardwareEventChannel)

	// Initialize display driver
	display := initializeDisplay()

	// Start game with welcome message
	welcomeMessage := "Welcome to Clairvoyant. Please press the knob to get started."
	welcomeMessageBytes := []byte(welcomeMessage)
	display.dev.Write(welcomeMessageBytes)

knob:
	for {
		select {
		case event := <-hardwareEventChannel:
			if event == "KNOB" {
				GAME_START = true
				break knob
			}
		}
	}

	// Retrieve input on number of players and display
	if GAME_START {
		display.dev.Write([]byte(strconv.Itoa(NUM_PLAYERS)))

		// Detect and display counter based on rotations of the rotary encoder
	rot:
		for {
			select {
			case event := <-hardwareEventChannel:
				if event == "CW" {
					// Display increment of player count
					NUM_PLAYERS += 1
					display.dev.Write([]byte(strconv.Itoa(NUM_PLAYERS)))
				}

				if event == "CCW" {
					// Display decrement of player count
					if NUM_PLAYERS != 0 {
						NUM_PLAYERS -= 1
						display.dev.Write([]byte(strconv.Itoa(NUM_PLAYERS)))
					} else {
						display.dev.Write([]byte(strconv.Itoa(0)))
					}
				}

				if event == "KNOB" {
					// Set player count and continue
					break rot
				}
			}
		}

		// Randomly select a player to start
		firstPlayer := ""

		randInt := rand.Intn(NUM_PLAYERS)
		switch randInt {
		case 0:
			firstPlayer = "RED"
		case 1:
			firstPlayer = "BLUE"
		case 2:
			firstPlayer = "GREEN"
		case 3:
			firstPlayer = "YELLOW"
		case 4:
			firstPlayer = "BLACK"
		}

		display.dev.Write([]byte(firstPlayer + " goes first! Hand the crystal ball to them!"))

		guessLoop(&hardwareEventChannel, display)

		goto start

	}
}
