package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())
	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/London")
	londonTime := t.In(loc)
	fmt.Println("London:", londonTime)

	myDate := "23 May 2017"
	d, _ := time.Parse("02 January 2006", myDate)
	fmt.Println(d)

	myDate1 := "23 May 2016"
	d1, _ := time.Parse("02 February 2006", myDate1)
	fmt.Println(d1)

	myDT := "Tuesday 23 May 2017 at 23:36"
	dt, _ := time.Parse("Monday 02 January 2006 at 15:04", myDT)
	fmt.Println(dt)
}
