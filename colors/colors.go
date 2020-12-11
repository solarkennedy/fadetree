package colors

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
)

type Color struct {
	R, G, B uint8
}

func PrintColors(colors []Color, occasion string, day time.Time) {
	fmt.Printf("For %s (%s) our colors will be [", occasion, day)
	for _, c := range colors {
		PrintColorBlock(c)
	}
	fmt.Println("]")
}

func PrintColorBlock(c Color) {
	cprint := color.RGB(c.R, c.G, c.B) // fg color
	cprint.Print("█")
}

func GetDaysColors(day time.Time) []Color {
	var colors []string
	occasion := ""

	if day.IsZero() {
		return []Color{}
	}

	if TodayIs("January 1", day) {
		occasion = "New Years Day"
		colors = []string{
			"#ffd700",
			"#000000",
		}
	} else if TodayIs("January 18", day) {
		occasion = "Woman's March"
		colors = []string{
			"#ff4fc7",
			"#000000",
		}
	} else if TodayIs("January 20", day) {
		occasion = "MLK Day"
		colors = []string{
			"#ef3423",
			"#ffd102",
			"#2e9743",
			"#000000",
		}
	} else if TodayIs("January 25", day) {
		occasion = "Chinese New Year"
		colors = []string{
			"#000000",
			"#ff0000",
			"#be1800",
			"#ffd425",
			"#efc416",
		}
	} else if TodayIs("January 29", day) {
		occasion = "NEN Awards"
		colors = []string{
			"#97c1d1",
			"#7baa7c",
			"#fabb18",
		}
	} else if TodayIsRange("January 1", 31, day) {
		occasion = "January"
		colors = []string{
			"#0d0a5e",
			"#7d7b90",
			"#3c8d87",
			"#6d418b",
			"#7c4369",
			"#ffffff",
		}
	} else if TodayIsRange("February 1", 2, day) {
		occasion = "Red/Gold - 49ers at Superbowl"
		colors = []string{
			"#bb3431",
			"#ba966a",
		}
	} else if TodayIs("February 4", day) {
		occasion = "Emperor Norton's 200th Birthday"
		colors = []string{
			"#d4af37",
			"#000000",
		}
	} else if TodayIsRange("February 19", 5, day) {
		occasion = "Black History Month"
		colors = []string{
			"#ef3423",
			"#ffd102",
			"#2e9743",
			"#000000",
		}
	} else if TodayIsRange("February 1", 28, day) {
		occasion = "February"
		colors = []string{
			"#ec9cc0",
			"#f0b1d4",
			"#f6cff5",
			"#e3dbff",
			"#d2d5fd",
		}
	} else if TodayIs("March 8", day) {
		occasion = "International Women's Day"
		colors = []string{
			"#574a72",
			"#000000",
		}
	} else if TodayIs("March 17", day) {
		occasion = "Saint Patrick's Day"
		colors = []string{
			"#009959",
			"#000000",
		}
	} else if TodayIsRange("March 19", 2, day) {
		occasion = "Nowruz/Persian New Year"
		colors = []string{
			"#ff0000",
			"#00ff00",
			"#ffffff",
			"#000000",
		}
	} else if TodayIs("March 21", day) {
		occasion = "American Red Cross Day"
		colors = []string{
			"#ff0000",
			"#000000",
		}
	} else if TodayIs("March 24", day) {
		occasion = "World TB Day"
		colors = []string{
			"#990000",
			"#000000",
		}
	} else if TodayIs("March 25", day) {
		occasion = "National Cerebral Palsy Awareness Month"
		colors = []string{
			"#1f9e1f",
			"#000000",
		}
	} else if TodayIs("March 26", day) {
		occasion = "Colon Cancer Awareness Month"
		colors = []string{
			"#0080ff",
			"#000000",
			"#ffffff",
		}
	} else if TodayIsRange("March 1", 31, day) {
		occasion = "March"
		colors = []string{
			"#6f8bc7",
			"#54b495",
			"#4a925a",
			"#458d35",
			"#3f4b85",
			"#000000",
		}
	} else if TodayIsRange("April 1", 30, day) {
		occasion = "April"
		colors = []string{
			"#b895d6",
			"#c5b5e5",
			"#ced8ff",
			"#e8ffd2",
			"#f8f9ab",
		}
	} else if TodayIsRange("May 1", 31, day) {
		occasion = "May"
		colors = []string{
			"#1bec7b",
			"#02f68b",
			"#31ff96",
			"#37ccde",
			"#22cbf9",
		}
	} else if TodayIsRange("June 1", 30, day) {
		occasion = "June"
		colors = []string{
			"#ffef3f",
			"#70e0ff",
			"#e3a0f2",
			"#ff9adb",
			"#ccff00",
		}
	} else if TodayIsRange("July 1", 31, day) {
		occasion = "July"
		colors = []string{
			"#2cc950",
			"#264e5a",
			"#029ac9",
			"#05ba7d",
			"#337c54",
		}
	} else if TodayIs("August 8", day) {
		occasion = "Pantone 448c"
		colors = []string{
			"#4a412a",
			"#000000",
		}
	} else if TodayIsRange("August 1", 31, day) {
		occasion = "August"
		colors = []string{
			"#0287bc",
			"#0967a7",
			"#205d92",
			"#36cc6d",
			"#278342",
		}
	} else if TodayIsRange("September 1", 30, day) {
		occasion = "September"
		colors = []string{
			"#4eb081",
			"#04c1ba",
			"#088155",
			"#5daa8b",
			"#0398d3",
		}
	} else if TodayIsRange("October 1", 31, day) {
		occasion = "October"
		colors = []string{
			"#8d2f00",
			"#501400",
			"#590105",
			"#170803",
			"#000700",
		}
	} else if TodayIsRange("November 1", 30, day) {
		occasion = "November"
		colors = []string{
			"#c36400",
			"#751400",
			"#ffa500",
		}
	} else if TodayIs("December 1", day) {
		occasion = "World Aids Day"
		colors = []string{
			"#eb0000",
			"#000000",
		}
	} else if TodayIs("December 10", day) {
		occasion = "First night of Hannukah"
		colors = []string{
			"#005EB8",
			"#ffffff",
		}
	} else if TodayIsRange("December 20", 6, day) {
		occasion = "Christmas"
		colors = []string{
			"#ffffff",
			"#eb0000",
			"#02d92a",
			"#cbc967",
		}
	} else if TodayIs("December 26", day) {
		occasion = "Kwanza"
		colors = []string{
			"#ef3423",
			"#ffd102",
			"#2e9743",
			"#000000",
		}
	} else if TodayIs("December 29", day) {
		occasion = "Last night of Hannukah"
		colors = []string{
			"#005EB8",
			"#ffffff",
		}
	} else if TodayIsRange("December 1", 31, day) {
		occasion = "December"
		colors = []string{
			"#571614",
			"#eb0000",
			"#02d92a",
			"#092e05",
		}
	} else {
		occasion = "(No Occasion)"
		colors = []string{}
	}
	c := ConvertPalette(colors)
	PrintColors(c, occasion, day)
	return c
}

func ConvertPalette(input []string) []Color {
	output := []Color{}
	for _, c := range input {
		output = append(output, ParseHexColor(c))
	}
	return output
}

func TodayIs(input string, today time.Time) bool {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	return day == today.Day() && month == today.Month().String()
}

func TodayIsRange(input string, n int, today time.Time) bool {
	//BUG: Normalizes the date comparison within the same year,
	//so it can't really span year boundaries
	//	fmt.Printf("Is %s within %d days after %s?\n", today, n, input)
	input_date := parse_input_date(input, today)
	last_date := input_date.AddDate(0, 0, n)
	//	fmt.Printf("(Between %s and %s\n", input_date, last_date)
	result := (today.After(input_date) && today.Before(last_date)) || input_date == today
	//	fmt.Println(result)
	return result
}

func parse_input_date(input string, normalized_day time.Time) time.Time {
	s := strings.Split(input, " ")
	month := s[0]
	day, _ := strconv.Atoi(s[1])
	parsed := time.Date(normalized_day.Year(), MonthToMonth(month), day, 0, 0, 0, 0, normalized_day.Location())
	return parsed
}

func MonthToMonth(input string) time.Month {
	fake_date := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Now().Location())
	for i := 1; i <= 12; i++ {
		if fake_date.Month().String() == input {
			return fake_date.Month()
		}
		fake_date = fake_date.AddDate(0, 1, 0)
	}
	return time.Month(1)
}

func ParseHexColor(s string) (c Color) {
	if s[0] != '#' {
		log.Fatalf("Color hex must start with #: %s", s)
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		log.Fatalf("Can't interpret byte %b", b)
		return 0
	}
	c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
	c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
	c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	return
}
