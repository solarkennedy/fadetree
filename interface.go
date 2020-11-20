package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brian-armstrong/gpio"
)

func guessGPIOPin() int {
	env := os.Getenv("FADETREE_MOTION_GPIO")
	if env == "" {
		return 0
	}
	pin, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}
	return pin
}

func (f *FadeTree) pollForMotion() {
	gpioPin := guessGPIOPin()
	if gpioPin != 0 {
		f.pollForMotionOnGPIO(gpioPin)
	} else {
		log.Print("No GPIO port detected. Skipping motion detection")
		f.Wakeness = 255
	}
}

func (f *FadeTree) pollForMotionOnGPIO(gpioPin int) {
	// https://openwrt.org/toh/tp-link/tl-wr703n#gpios
	p := gpio.NewInput(uint(gpioPin))
	f.Wakeness = 0
	f.WakenessRate = 0

	go func() {
		for {
			value, err := p.Read()
			if err != nil {
				fmt.Printf("Got error reading gpio pin %d: %s", gpioPin, err)
			}
			if f.MotionGPIOValue != value {
				f.MotionGPIOValue = value
				if value == 1 {
					fmt.Print("Motion Detected!")
					f.WakenessRate = 30
				} else {
					fmt.Print("Motion stopped")
					f.WakenessRate = -10
				}
			}
			time.Sleep(1 * time.Second)
			f.CalculateWakeness()
		}
	}()
}

func (f *FadeTree) CalculateWakeness() {
	w := int(f.Wakeness) + int(f.WakenessRate)
	if w < 0 {
		f.Wakeness = 0
	} else if w > 255 {
		f.Wakeness = 255
	} else {
		f.Wakeness = uint8(w)
	}
}
