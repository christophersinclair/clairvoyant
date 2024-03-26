package main

func setupInput(hwEventChan *chan string) {
	// Initialize five player buttons
	redButton := initializeButton("red", RED_BUTTON_GPIO)
	blueButton := initializeButton("blue", BLUE_BUTTON_GPIO)
	yellowButton := initializeButton("yellow", YELLOW_BUTTON_GPIO)
	greenButton := initializeButton("green", GREEN_BUTTON_GPIO)
	blackButton := initializeButton("black", BLACK_BUTTON_GPIO)

	// Initialize rotary encoder
	rotaryEncoder := initializeRotary(ROTARY_SW_GPIO, ROTARY_CLK_GPIO, ROTARY_DT_GPIO)

	// Create channels for specific types of hardware events
	redButtonEventChannel := make(chan string)
	blueButtonEventChannel := make(chan string)
	yellowButtonEventChannel := make(chan string)
	greenButtonEventChannel := make(chan string)
	blackButtonEventChannel := make(chan string)
	rotaryEncoderEventChannel := make(chan string)

	// Kick off monitoring goroutines
	go redButton.monitor(&redButtonEventChannel)
	go blueButton.monitor(&blueButtonEventChannel)
	go yellowButton.monitor(&yellowButtonEventChannel)
	go greenButton.monitor(&greenButtonEventChannel)
	go blackButton.monitor(&blackButtonEventChannel)
	go rotaryEncoder.monitor(&rotaryEncoderEventChannel)

	go func(hardwareChannel *chan string) {
		for {
			select {
			case event := <-redButtonEventChannel:
				if event == "High" {
					// Red button was pressed
					*hardwareChannel <- "RED"
				}
			case event := <-blueButtonEventChannel:
				if event == "High" {
					// Blue button was pressed
					*hardwareChannel <- "BLUE"
				}
			case event := <-yellowButtonEventChannel:
				if event == "High" {
					// Yellow button was pressed
					*hardwareChannel <- "YELLOW"
				}
			case event := <-greenButtonEventChannel:
				if event == "High" {
					// Green button was pressed
					*hardwareChannel <- "GREEN"
				}
			case event := <-blackButtonEventChannel:
				if event == "High" {
					// Black button was pressed
					*hardwareChannel <- "BLACK"
				}
			case event := <-rotaryEncoderEventChannel:
				if event == "CW" {
					// Rotary encoder was turned clockwise
					*hardwareChannel <- "CW"
				} else if event == "CCW" {
					// Rotary encoder was turned counterclockwise
					*hardwareChannel <- "CCW"
				} else if event == "High" {
					// Rotary encoder knob was pressed
					*hardwareChannel <- "KNOB"
				}
			}
		}
	}(hwEventChan)
}
