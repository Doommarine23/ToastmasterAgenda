package main

import (
	"strconv"
	"time"
)

func getNextTuesday(t time.Time) time.Time {
	for int(t.Weekday()) != 2 { //todo figure out how to use the constant instead of special number 2! t.Weekday.Tuesday {
		t = t.AddDate(0, 0, 1)
	}

	return t
}

func AgendaDate(t time.Time) string {
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	year := strconv.Itoa(t.Year())

	return month + "." + day + "." + year
}

func DateWithSlashes(t time.Time) string {
	month := strconv.Itoa(int(t.Month()))
	day := strconv.Itoa(t.Day())
	year := strconv.Itoa(t.Year())

	return month + "/" + day + "/" + year
}

func AgendaMonthDayYear(t time.Time) string {
	day := strconv.Itoa(t.Day())
	month := t.Month().String()
	year := strconv.Itoa(t.Year())

	return month + " " + day + ", " + year
}

func prettyPrintTime(curTime time.Time, minToAdd int) (time.Time, string) {

	nextTime := time.Minute * time.Duration(minToAdd)
	curTime = curTime.Add(nextTime)
	hour, min, _ := curTime.Clock()

	return curTime, strconv.Itoa(hour) + ":" + strconv.Itoa(min)
}
