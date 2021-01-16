package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
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

func inDebugMode() bool {
	return os.Getenv("VSCODE_PID") != ""
}

func random(min, max int) uint8 {
	if max-min <= 0 {
		return 0
	}
	xr := rand.Intn(max-min) + min
	return uint8(xr)
}

func shouldIBeOn() bool {
	if getEnvOverride() != "" {
		return true
	} else if inDebugMode() {
		return true
	} else {
		now := Now()
		hour := now.Hour()
		// TODO: Use sunrise,set
		//return (now.After(set) && hour <= 21) || (now.After(rise) && hour <= 7)
		return (hour >= 18 && hour <= 21) || (hour > 6 && hour <= 7)
	}
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

func imagesInMarkDown(colors []colors.Color) string {
	output := ""
	for _, c := range colors {
		colorHex := fmt.Sprintf("%x%x%x", c.R, c.G, c.B)
		output = fmt.Sprintf("%s![](https://raster.shields.io/badge/-%%20-%s?style=flat-square)", output, colorHex)

	}
	return output
}

func (f *FadeTree) setDailySettings() {
	f.Today = getToday()
	current := f.ColorPalette
	new, occasion := colors.GetDaysColors(f.Today)
	if !cmp.Equal(current, new) {
		sendPushNotification(
			fmt.Sprintf("New Colors for %s!", occasion),
			fmt.Sprintf("Picking new colors for %s: %s", f.Today, imagesInMarkDown(new)),
		)
		f.ColorPalette = new
	}
}

func (f *FadeTree) startDayTicker() {
	f.setDailySettings()
	ticker := time.NewTicker(time.Hour)
	for {
		t := <-ticker.C
		log.Printf("Got a DayTicker wakeup at %s", t)
		f.setDailySettings()
	}
}

func (f *FadeTree) setBrightness() {
	// TODO: Fade better?
}

func (f *FadeTree) startBrightnessTicker() {
	f.setBrightness()
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		f.setBrightness()
	}
}

func (f *FadeTree) runWatcher() {
	go f.startDayTicker()
	go f.startBrightnessTicker()

	for {
		r := randomJar(f.Jars)
		if shouldSetColorOnRandomJar() {
			displayPatternOnJar(f.OpcClient, r, f.ColorPalette, f.Wakeness)
			Sync(f.OpcClient, f.Jars)
		}
		//displayPattern(f.OpcClient, f.Jars, f.ColorPalette, f.Brightness)
		fmt.Printf("Brightness: %d\n", f.Wakeness)
		time.Sleep(time.Duration(1000 * time.Millisecond))
	}

}

func randomJar(jars []Jar) Jar {
	return jars[rand.Intn(len(jars))]
}

func shouldSetColorOnRandomJar() bool {
	return true
}
