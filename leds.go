package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
)

func displayPattern(oc *opc.Client, leds_len int, color_palette []colors.Color) {
	m := opc.NewMessage(0)
	led_grouping := 1
	if len(color_palette) == 0 {
		for i := 0; i < leds_len; i++ {
			m.SetLength(uint16(leds_len * 3))
			m.SetPixelColor(i, random(2, 255), random(2, 255), random(2, 255))
		}
	} else {
		for i := 0; i < leds_len; i += led_grouping {
			c := color_palette[rand.Intn(len(color_palette))]
			for j := i; j < (i + led_grouping); j++ {
				m.SetLength(uint16(leds_len * 3))
				m.SetPixelColor(j, c.R, c.G, c.B)
				colors.PrintColorBlock(c)
			}
		}
	}
	err := oc.Send(m)
	fmt.Println()
	if err != nil {
		log.Println("couldn't send color", err)
	}

}

func turnOff(oc *opc.Client, leds_len int) {
	m := opc.NewMessage(0)
	for i := 0; i < leds_len; i++ {
		m.SetLength(uint16(leds_len * 3))
		m.SetPixelColor(i, 0, 0, 0)
	}
	err := oc.Send(m)
	if err != nil {
		log.Println("couldn't send color", err)
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
