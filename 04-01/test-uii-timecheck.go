package main

import (
	"fmt"
	"time"
)

func IsRangeTime(dateTimeStr, startRangeStr, endRangeStr string) bool {
	layout := "2006-01-02 15:04:05"

	dateTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		panic("Invalid date and time format for dateTimeStr")
	}

	startRange, err := time.Parse(layout, startRangeStr)
	if err != nil {
		panic("Invalid date and time format for startRangeStr")
	}

	endRange, err := time.Parse(layout, endRangeStr)
	if err != nil {
		panic("Invalid date and time format for endRangeStr")
	}

	//cek range date time
	result := dateTime.After(startRange) && dateTime.Before(endRange)
	return result
}

func main() {
	dateTime := "2023-06-26 10:30:00"
	startRange := "2023-06-26 08:00:00"
	endRange := "2023-06-26 12:00:00"

	if IsRangeTime(dateTime, startRange, endRange) {
		fmt.Println("Tanggal dan waktu berada dalam rentang tanggal dan waktu")
	} else {
		fmt.Println("Tanggal dan waktu tidak berada dalam rentang tanggal dan waktu")
	}

	dateTime = "2023-06-26 15:45:30"
	startRange = "2023-06-27 18:00:00"
	endRange = "2023-06-27 20:00:00"

	if IsRangeTime(dateTime, startRange, endRange) {
		fmt.Println("Tanggal dan waktu berada dalam rentang tanggal dan waktu")
	} else {
		fmt.Println("Tanggal dan waktu tidak berada dalam rentang tanggal dan waktu")
	}
}
