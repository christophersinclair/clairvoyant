package main

import (
	"math/rand"
	"strconv"
)

func waitForPlayerSelection(hardwareEventChannel *chan string, display *Display) int {
	NUM_PLAYERS := 0
	display.show(strconv.Itoa(NUM_PLAYERS))

	for {
		select {
		case event := <-*hardwareEventChannel:
			if event == "CW" {
				// Display increment of player count
				NUM_PLAYERS += 1
				display.show(strconv.Itoa(NUM_PLAYERS))
			}

			if event == "CCW" {
				// Display decrement of player count
				if NUM_PLAYERS != 0 {
					NUM_PLAYERS -= 1
					display.show(strconv.Itoa(NUM_PLAYERS))
				} else {
					display.show(strconv.Itoa(0))
				}
			}

			if event == "KNOB" {
				// Set player count and continue
				return NUM_PLAYERS
			}
		}
	}
}

func main() {
	GAME_START := false

	hardwareEventChannel := make(chan string)
	setupInput(&hardwareEventChannel)

	// Initialize display driver
	display := initializeDisplay()

	// Start game with welcome message
	welcomeMessage := "Welcome to Clairvoyant. Please press the knob to get started."
	display.show(welcomeMessage)

	for {
		select {
		case event := <-hardwareEventChannel:
			if event == "KNOB" {
				GAME_START = true
				break
			}
		}
	}

	// Retrieve input on number of players and display
	if GAME_START {
		// Detect and display counter based on rotations of the rotary encoder
		NUM_PLAYERS := waitForPlayerSelection(&hardwareEventChannel, display)

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

		display.show(firstPlayer + " goes first! Hand the crystal ball to them!")

		guessLoop(&hardwareEventChannel, display)
	}
}
