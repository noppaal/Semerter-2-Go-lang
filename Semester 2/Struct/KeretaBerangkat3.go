package main

import (
	"fmt"
)

type Time struct {
	Hour   int
	Minute int
}

type TrainSchedule struct {
	TrainNumber string
	Origin      string
	Departure   Time
	Destination string
	Arrival     Time
}

func main() {
	// Input data
	var schedules [3]TrainSchedule
	for i := 0; i < 3; i++ {
		var (
			trainNum, origin, dest string
			depHour, depMinute     int
			arrHour, arrMinute     int
		)
		fmt.Scan(&trainNum, &origin, &depHour, &depMinute, &dest, &arrHour, &arrMinute)
		schedules[i] = TrainSchedule{
			TrainNumber: trainNum,
			Origin:      origin,
			Departure:   Time{Hour: depHour, Minute: depMinute},
			Destination: dest,
			Arrival:     Time{Hour: arrHour, Minute: arrMinute},
		}
	}

	// Mencari jadwal kereta api terlama
	longestTrain := schedules[0].TrainNumber
	longestDuration := calculateDuration(schedules[0].Departure, schedules[0].Arrival)

	for _, schedule := range schedules[1:] {
		duration := calculateDuration(schedule.Departure, schedule.Arrival)
		if duration > longestDuration {
			longestTrain = schedule.TrainNumber
			longestDuration = duration
		}
	}

	// Output hasil
	fmt.Println(longestTrain)
}

func calculateDuration(departure Time, arrival Time) int {
	depMinutes := departure.Hour*60 + departure.Minute
	arrMinutes := arrival.Hour*60 + arrival.Minute
	return arrMinutes - depMinutes
}
