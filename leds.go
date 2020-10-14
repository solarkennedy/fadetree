package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
)

func setRandomColors(jars []Jar) {
	for _, j := range jars {
		for _, led := range j.Leds {
			led.R = random(2, 255)
			led.G = random(2, 255)
			led.B = random(2, 255)
		}
	}
}

func setColorsFromPalette(jars []Jar, color_palette []colors.Color) {
	for _, j := range jars {
		for _, led := range j.Leds {
			c := color_palette[rand.Intn(len(color_palette))]
			led.R = c.R
			led.G = c.G
			led.B = c.B
		}
	}
}

func displayPattern(oc *opc.Client, jars []Jar, color_palette []colors.Color) {
	if len(color_palette) == 0 {
		setRandomColors(jars)
	} else {
		setColorsFromPalette(jars, color_palette)
	}
	Sync(oc, jars)
}

func (j Jar) TurnOff(oc *opc.Client) {
	for _, led := range j.Leds {
		// TODO set to zero
		led.B = 0
		led.G = 0
		led.R = 0
	}
}

func turnOffAllJars(oc *opc.Client, jars []Jar) {
	for _, j := range jars {
		j.TurnOff(oc)
	}
	Sync(oc, jars)
}

func Sync(oc *opc.Client, jars []Jar) {
	m := opc.NewMessage(0)
	counter := 0
	for jarCounter, j := range jars {
		for ledCounter, led := range j.Leds {
			counter = jarCounter * ledCounter
			m.SetPixelColor(counter, led.R, led.G, led.B)
			colors.PrintColorBlock(led)
		}
	}
	fmt.Println()
	m.SetLength(uint16(counter * 3))
	err := oc.Send(m)
	if err != nil {
		log.Println("couldn't send opc message", err)
	}
}

func getOCClient() *opc.Client {
	server := "fadetree:7890"
	oc := opc.NewClient()
	err := oc.Connect("tcp", server)
	if err != nil {
		log.Fatal("Could not connect to Fadecandy server", err)
	}
	return oc
}
