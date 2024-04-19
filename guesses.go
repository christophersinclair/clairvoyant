package main

import (
	"strconv"
	"time"
)

type Guess struct {
	guesser      string
	target       string
	turnsToGo    int
	turnItIs     int
	turnItWillBe int
}

func waitForButton(hardwareChannel *chan string) string {
	for {
		select {
		case event := <-*hardwareChannel:
			if event != "CW" && event != "CCW" && event != "KNOB" {
				return event
			}
		}
	}
}

func waitForTurns(hardwareChannel *chan string, display *Display) int {
	turns := 2
	for {
		select {
		case event := <-*hardwareChannel:
			if event == "CW" {
				// Display increment of turns
				turns += 1
				display.show(strconv.Itoa(turns))
			}

			if event == "CCW" {
				// Display decrement of turns
				if turns != 2 {
					turns -= 1
					display.show(strconv.Itoa(turns))
				} else {
					display.show(strconv.Itoa(2))
				}
			}

			if event == "KNOB" {
				// Set turns and continue
				return turns
			}
		}
	}
}

func guessLoop(hardwareChannel *chan string, display *Display) {
	guesses := []Guess{}
	turn := 0

	for {
		turn += 1

		display.show("Hello! Which player are YOU? (Please press your colored button!)")

		// Select holding player
		g := waitForButton(hardwareChannel)

		// Check if anyone has won yet
		if turn != 1 {
			for i := 0; i < len(guesses); i++ {
				gu := guesses[i]

				if gu.turnItWillBe == turn && gu.target == g {
					// Game end
					display.show("CONGRATULATIONS TO PLAYER " + gu.guesser + "! They guessed that player " + gu.target + " would have the crystal ball in " + strconv.Itoa(gu.turnsToGo) + ", " + strconv.Itoa(gu.turnsToGo) + " turns ago!")

					time.Sleep(20 * time.Second)
					display.show("GAME OVER. THANKS FOR PLAYING!")
				}
			}
		}

		display.show("Hello " + g + "! How many turns into the future do you see? (Use the dial to select how far into the future you can look)")

		// Select turns
		turns := waitForTurns(hardwareChannel, display)

		display.show("Now, who do you want to target? (Press the colored button matching the player you think will have it in " + strconv.Itoa(turns) + " turns!)")

		// Select target player
		target := waitForButton(hardwareChannel)

		guess := Guess{guesser: g, turnsToGo: turns, target: target, turnItIs: turn, turnItWillBe: (turn + turns)}
		guesses = append(guesses, guess)
	}
}
