package main

import (
	"fmt"
	"log"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"

	"go.bug.st/serial.v1"
)

func guessSerialPort() string {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		return ""
	}
	for _, port := range ports {
		fmt.Printf("Guessing serial port: %v\n", port)
		return port
	}
	return ""
}

func (f FadeTree) pollForMotion() {
	serialPort := guessSerialPort()
	if serialPort != "" {
		f.setupGobotAndRun(serialPort)
	} else {
		log.Print("No serial port detected. Skipping fermata interface")
	}
}

func (f FadeTree) setupGobotAndRun(serialPort string) {
	firmataAdaptor := firmata.NewAdaptor(serialPort)
	sensor := gpio.NewPIRMotionDriver(firmataAdaptor, "2")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		_ = sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Println(gpio.MotionDetected)
			f.MotionDetected = true
			_ = led.On()
		})
		_ = sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Println(gpio.MotionStopped)
			f.MotionDetected = false
			_ = led.Off()
		})
	}
	robot := gobot.NewRobot("motionBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor, led},
		work,
	)

	err := robot.Start()
	if err != nil {
		panic(err)
	}
	_ = firmataAdaptor.Finalize()
	_ = led.Connection().Finalize()
	_ = sensor.Connection().Finalize()
	os.Exit(1)
}
