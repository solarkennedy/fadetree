package main

import (
	"github.com/kellydunn/go-opc"
	"gobot.io/x/gobot/drivers/gpio"
)

var (
	NumJars = 10
)

type Jar struct {
	NumLeds int
}

type FadeTree struct {
	MotionSensor   *gpio.PIRMotionDriver
	Jars           []Jar
	NumLeds        int
	OpcClient      *opc.Client
	MotionDetected bool
}

func main() {
	var f FadeTree
	f.Jars = make([]Jar, NumJars)
	f.NumJars = 5
	f.OpcClient = getOCClient()
	go f.pollForMotion()
	go runWebserver()
	f.runWatcher()
}
