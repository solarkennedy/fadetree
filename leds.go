package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
)

func setRandomColors(jars []Jar, brightness uint8) {
	for _, jar := range jars {
		setRandomColorsOnJar(jar, brightness)
	}
}

func setRandomColorsOnJar(jar Jar, brightness uint8) {
	for l := range jar.Leds {
		jar.Leds[l].R = random(0, int(brightness))
		jar.Leds[l].G = random(0, int(brightness))
		jar.Leds[l].B = random(0, int(brightness))
	}
}

func setColorsFromPalette(jars []Jar, color_palette []colors.Color, brightness uint8) {
	for _, jar := range jars {
		setColorsOnJar(jar, color_palette, brightness)
	}
}

func setColorsOnJar(jar Jar, color_palette []colors.Color, brightness uint8) {
	c := color_palette[rand.Intn(len(color_palette))]
	for l := range jar.Leds {
		jar.Leds[l].R = uint8(int(c.R) * int(brightness) / 255)
		jar.Leds[l].G = uint8(int(c.G) * int(brightness) / 255)
		jar.Leds[l].B = uint8(int(c.B) * int(brightness) / 255)
	}
}

func displayPattern(oc *opc.Client, jars []Jar, color_palette []colors.Color, brightness uint8) error {
	if len(color_palette) == 0 {
		setRandomColors(jars, brightness)
	} else {
		setColorsFromPalette(jars, color_palette, brightness)
	}
	return Sync(oc, jars)
}

func displayPatternOnJar(oc *opc.Client, jar Jar, color_palette []colors.Color, brightness uint8) {
	if len(color_palette) == 0 {
		setRandomColorsOnJar(jar, brightness)
	} else {
		setColorsOnJar(jar, color_palette, brightness)
	}
}

func (f *FadeTree) turnOffAllJars() {
	oc := f.OpcClient
	for i := range f.Jars {
		for l := range f.Jars[i].Leds {
			f.Jars[i].Leds[l].R = 0
			f.Jars[i].Leds[l].G = 0
			f.Jars[i].Leds[l].B = 0
		}
	}
	_ = Sync(oc, f.Jars)
}

func Sync(oc *opc.Client, jars []Jar) error {
	m := opc.NewMessage(0)
	LEDCursor := 0
	for _, j := range jars {
		for _, led := range j.Leds {
			m.SetPixelColor(LEDCursor, led.G, led.R, led.B)
			//m.SetPixelColor(counter, led.R, led.G, led.B)
			LEDCursor = LEDCursor + 1
		}
	}
	m.SetLength(uint16(LEDCursor * 3))
	err := oc.Send(m)
	if err != nil {
		return err
	}
	printStatus(jars)
	return nil
}

func printStatus(jars []Jar) {
	if isATTY() {
		for _, j := range jars {
			for _, led := range j.Leds {
				colors.PrintColorBlock(led)
			}
			fmt.Print("\t")
		}
	}
}

func isATTY() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	} else {
		return false
	}
}

func getOCClient() *opc.Client {
	server := "fadetree:7890"
	oc := opc.NewClient()
	err := oc.Connect("tcp", server)
	if err != nil {
		log.Printf("Could not connect to Fadecandy server %s", err)
		return nil
	}
	return oc
}
