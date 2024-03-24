package main

func main() {
	// Initialize display driver
	display := initializeDisplay()

	// Initialize four player buttons and one power button
	redButton := initializeButton("red", "GPIO17")
	blueButton := initializeButton("blue", "GPIO27")
	yellowButton := initializeButton("yellow", "GPIO22")
	greenButton := initializeButton("green", "GPIO23")
	blackButton := initializeButton("black", "GPIO24")

	// Initialize rotary encoder
	rotaryEncoder := initializeRotary("GPIO16", "GPIO20", "GPIO21")

	// Create channel for hardware events
	redButtonEventChannel := make(chan string)
	blueButtonEventChannel := make(chan string)
	yellowButtonEventChannel := make(chan string)
	greenButtonEventChannel := make(chan string)
	blackButtonEventChannel := make(chan string)

	rotaryEncoderEventChannel := make(chan string)

	// Kick off monitoring goroutines
	go redButton.monitor(&redButtonEventChannel)
	go blueButton.monitor(&blackButtonEventChannel)
	go yellowButton.monitor(&yellowButtonEventChannel)
	go greenButton.monitor(&greenButtonEventChannel)
	go blackButton.monitor(&blackButtonEventChannel)

	go rotaryEncoder.monitor(&rotaryEncoderEventChannel)

	for {
		select {
		case event := <-redButtonEventChannel:
			if event == "High" {
				// Red button was pressed
			}
		case event := <-blueButtonEventChannel:
			if event == "High" {
				// Blue button was pressed
			}
		case event := <-yellowButtonEventChannel:
			if event == "High" {
				// Yellow button was pressed
			}
		case event := <-greenButtonEventChannel:
			if event == "High" {
				// Green button was pressed
			}
		case event := <-blackButtonEventChannel:
			if event == "High" {
				// Black button was pressed
			}
		case event := <-rotaryEncoderEventChannel:
			if event == "CW" {
				// Rotary encoder was turned clockwise
				display.show(make([]byte, 0))
			} else if event == "CCW" {
				// Rotary encoder was turned counterclockwise
			}
		}
	}
}
