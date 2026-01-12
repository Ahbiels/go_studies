package main

import (
	"fmt"
	"strconv"
	"switch/data/test/module"
	"time"
)

func main() {
	data := time.Now()
    today := data.Format(("Monday"))
	hours, _ := strconv.Atoi(data.Format(("15")))
	today_int := module.DayOfWeek(today)
	hours_int := module.HourOfTheDay(hours)
	fmt.Printf("Today is %v, day %v in the week, and it's %v\n\n", today, today_int, hours_int)
	

}