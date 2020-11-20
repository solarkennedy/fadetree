package main

import (
	"time"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
)

var (
	NumJars       = 3
	NumLedsPerJar = 5
)

type Jar struct {
	Leds   []colors.Color
	Candle bool
}

type FadeTree struct {
	MotionGPIOValue uint
	Jars            []Jar
	OpcClient       *opc.Client
	MotionDetected  bool
	Today           time.Time
	ColorPalette    []colors.Color
	Wakeness        uint8
	WakenessRate    int
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
	f.runWatcher()
}
