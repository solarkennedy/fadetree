package main

import (
	"fmt"

	"github.com/kellydunn/go-opc"
	"github.com/solarkennedy/fadetree/colors"
	"gobot.io/x/gobot/drivers/gpio"
)

var (
	NumJars       = 10
	NumLedsPerJar = 5
)

type Jar struct {
	Leds []colors.Color
}

type FadeTree struct {
	MotionSensor   *gpio.PIRMotionDriver
	Jars           []Jar
	OpcClient      *opc.Client
	MotionDetected bool
}

func (f FadeTree) MakeJars() {
	f.Jars = make([]Jar, NumJars)
	//for _, jar := range f.Jars {
	//jar.Leds = make([]colors.Color, NumLedsPerJar)
	//jar.Leds = append(jar.Leds, colors.Color{1, 1, 1})
	//jar.Leds[0] = colors.Color{1, 1, 1}
	//fmt.Println(f.Jars)
	//}
	fmt.Println(f.Jars)
	fmt.Println(len(f.Jars))
}

func main() {
	var f FadeTree
	f.MakeJars()
	fmt.Println(f.Jars)
	fmt.Println(len(f.Jars))
	f.OpcClient = getOCClient()
	go f.pollForMotion()
	go runWebserver()
	f.runWatcher()
}
