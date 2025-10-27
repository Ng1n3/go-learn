package main

import (
	"fmt"
	"time"
)

func main() {

	// Current local time
	fmt.Println(time.Now())

	// Specific Time
	specificTime := time.Date(2025, time.October, 23, 11, 10, 0, 0, time.UTC)
	fmt.Println("Specific Time: ", specificTime)

	// Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")      // Mon jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")         // Mon jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") // Mon jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("Monday 06-01-02 15-04-05"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time: ", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2025, time.October, 27, 13, 50, 20, 00, time.UTC)

	// Convert this to the specific time zone
	tLocal := t.In(loc)

	// Perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original time (UTC): ", t)
	fmt.Println("Original time (Local): ", tLocal)
	fmt.Println("Rounded Time (UTC): ", roundedTime)
	fmt.Println("Rounded Time (Local): ", roundedTimeLocal)

	fmt.Println("Truncated Time: ", t.Truncate(time.Hour))

	loc2, _ := time.LoadLocation("America/New_York")

	// Convert time to location
	tInNY := time.Now().In(loc2)
	fmt.Println("New York Time: ", tInNY)

	t1 := time.Date(2025, time.October, 27, 13, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, time.October, 27, 15, 0, 0, 0, time.UTC)

	duration := t2.Sub(t1)
	fmt.Println(duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))
}
