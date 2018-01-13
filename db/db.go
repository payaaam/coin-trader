package db

var OneDayInterval string = "1D"
var FourHourInterval string = "4h"
var OneHourInterval string = "1h"
var ThirtyMinuteInterval string = "30m"
var MinuteMillisecond = 60000

var ValidIntervals = map[string]bool{
	OneDayInterval:       true,
	FourHourInterval:     true,
	OneHourInterval:      true,
	ThirtyMinuteInterval: true,
}

var IntervalMilliseconds = map[string]int{
	OneDayInterval:       MinuteMillisecond * 60 * 24,
	FourHourInterval:     MinuteMillisecond * 60 * 4,
	OneHourInterval:      MinuteMillisecond * 60,
	ThirtyMinuteInterval: MinuteMillisecond * 30,
}
