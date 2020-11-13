package main

import (
	"fmt"
	"log"

	"github.com/brian-armstrong/gpio"
)

func guessGPIOPin() int {
	return 0
}

func (f *FadeTree) pollForMotion() {
	gpioPin := guessGPIOPin()
	if gpioPin != 0 {
		f.pollForMotionOnGPIO(gpioPin)
	} else {
		log.Print("No serial port detected. Skipping fermata interface")
	}
}

func (f *FadeTree) pollForMotionOnGPIO(gpioPin int) {
	// https://openwrt.org/toh/tp-link/tl-wr703n#gpios
	watcher := gpio.NewWatcher()
	watcher.AddPin(14)
	defer watcher.Close()

	go func() {
		for {
			pin, value := watcher.Watch()
			fmt.Printf("read %d from gpio %d\n", value, pin)
		}
	}()

}
