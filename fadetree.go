package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
)

var (
	NumJars       = 3
	NumLedsPerJar = 5
)

type Jar struct {
	Leds []colors.Color
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
	go f.setupExitHandler()
	f.runWatcher()
}

func (f *FadeTree) setupExitHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("Got %s. Shutting down", sig)
			os.Exit(0)
		}
	}()
}
