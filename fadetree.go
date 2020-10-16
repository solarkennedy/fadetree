package main

import (
	"time"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
	"gobot.io/x/gobot/drivers/gpio"
)

var (
	NumJars       = 10
	NumLedsPerJar = 5
)

type Jar struct {
	Leds   []colors.Color
	Candle bool
}

type FadeTree struct {
	MotionSensor   *gpio.PIRMotionDriver
	Jars           []Jar
	OpcClient      *opc.Client
	MotionDetected bool
	Today          time.Time
	Sunrise        time.Time
	Sunset         time.Time
	ColorPalette   []colors.Color
	Brightness     uint8
}

func (f *FadeTree) MakeJars() {
	f.Jars = make([]Jar, NumJars)
	for c := range f.Jars {
		f.Jars[c].Leds = make([]colors.Color, NumLedsPerJar)
	}
}

func main() {
	var f FadeTree
	f.MakeJars()
	f.OpcClient = getOCClient()
	go f.pollForMotion()
	go runWebserver()
	f.runWatcher()
}
