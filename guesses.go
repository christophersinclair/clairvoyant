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

func guessLoop(hardwareChannel *chan string, display *Display) {
	guesses := []Guess{}
	turn := 0

main:
	for {
		turn += 1

		display.dev.Write([]byte("Hello! Which player are YOU? (Please press your colored button!)"))

		// Select holding player
		g := ""
	guesser:
		for {
			select {
			case event := <-*hardwareChannel:
				if event != "CW" && event != "CCW" && event != "KNOB" {
					g = event
					break guesser
				}
			}
		}

		// Check if anyone has won yet
		if turn != 1 {
			for i := 0; i < len(guesses); i++ {
				gu := guesses[i]

				if gu.turnItWillBe == turn && gu.target == g {
					// Game end
					display.dev.Write([]byte("CONGRATULATIONS TO PLAYER " + gu.guesser + "! They guessed that player " + gu.target + " would have the crystal ball in " + strconv.Itoa(gu.turnsToGo) + ", " + strconv.Itoa(gu.turnsToGo) + " turns ago!"))

					time.Sleep(20 * time.Second)
					display.dev.Write([]byte("GAME OVER. THANKS FOR PLAYING!"))
					break main
				}
			}
		}

		display.dev.Write([]byte("Hello " + g + "! How many turns into the future do you see? (Use the dial to select how far into the future you can look)"))

		// Select turns
		turns := 2
	turns:
		for {
			select {
			case event := <-*hardwareChannel:
				if event == "CW" {
					// Display increment of turns
					turns += 1
					display.dev.Write([]byte(strconv.Itoa(turns)))
				}

				if event == "CCW" {
					// Display decrement of turns
					if turns != 2 {
						turns -= 1
						display.dev.Write([]byte(strconv.Itoa(turns)))
					} else {
						display.dev.Write([]byte(strconv.Itoa(2)))
					}
				}

				if event == "KNOB" {
					// Set turns and continue
					break turns
				}
			}
		}

		display.dev.Write([]byte("Now, who do you want to target? (Press the colored button matching the player you think will have it in " + strconv.Itoa(turns) + " turns!)"))

		// Select target player
		target := ""
	target:
		for {
			select {
			case event := <-*hardwareChannel:
				if event != "CW" && event != "CCW" && event != "KNOB" {
					target = event
					break target
				}
			}
		}

		guess := Guess{guesser: g, turnsToGo: turns, target: target, turnItIs: turn, turnItWillBe: (turn + turns)}
		guesses = append(guesses, guess)
	}
}
