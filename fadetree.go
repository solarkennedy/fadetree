package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kellydunn/go-opc"
	"github.com/kelvins/sunrisesunset"
	"github.com/solarkennedy/fadetree/colors"
)

func Now() time.Time {
	pst, _ := time.LoadLocation("America/Los_Angeles")
	now := time.Now().In(pst)
	return now
}

func getEnvOverride() string {
	return os.Getenv("FADECANDYCAL_DATE")
}

func random(min, max int) uint8 {
	xr := rand.Intn(max-min) + min
	return uint8(xr)
}

func shouldIBeOn() bool {
	if getEnvOverride() != "" {
		return true
	} else {
		now := Now()
		hour := now.Hour()
		//		rise, set := getSunriseSunset()
		//return (now.After(set) && hour <= 21) || (now.After(rise) && hour <= 7)
		return (hour >= 18 && hour <= 21) || (hour > 6 && hour <= 7)
	}
}

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

func parseOverride(input string) time.Time {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	today := Now()
	parsed := time.Date(today.Year(), colors.MonthToMonth(month), day, 0, 0, 0, 0, today.Location())
	fmt.Printf("Parsed env override '%s' as '%s'\n", input, parsed)
	return parsed
}

func getToday() time.Time {
	override := getEnvOverride()
	if override != "" {
		return parseOverride(override)
	} else {
		return Now()
	}
}

func getUTCOffset(now time.Time) float64 {
	offset, err := strconv.Atoi(now.Format("-0700"))
	if err != nil {
		panic(err)
	}
	return float64(offset / 100)
}

func getSunriseSunset(now time.Time) (time.Time, time.Time) {
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	p := sunrisesunset.Parameters{
		Latitude:  37.774929,
		Longitude: -122.419418,
		UtcOffset: getUTCOffset(now),
		Date:      today,
	}
	sunrise, sunset, err := p.GetSunriseSunset()
	if err != nil {
		panic(err)
	}
	sunrise_today := time.Date(now.Year(), now.Month(), now.Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, now.Location())
	sunset_today := time.Date(now.Year(), now.Month(), now.Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), 0, now.Location())
	fmt.Println(" Sunrise:", sunrise_today.Format("3:04PM"), " / Sunset:", sunset_today.Format("3:04PM"))
	return sunrise_today, sunset_today
}

func main() {
	leds_len := 64
	oc := getOCClient()
	getSunriseSunset(time.Now())
	go runWebserver()

	go pollForMotion()

	for {
		today := getToday()
		color_palette := colors.GetDaysColors(today)
		if shouldIBeOn() {
			displayPattern(oc, leds_len, color_palette)
		} else {
			turnOff(oc, leds_len)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
