package db

var OneDayInterval string = "1D"
var FourHourInterval string = "4h"
var OneHourInterval string = "1h"
var ThirtyMinuteInterval string = "30m"

var ValidIntervals = map[string]bool{
	OneDayInterval:       true,
	FourHourInterval:     true,
	OneHourInterval:      true,
	ThirtyMinuteInterval: true,
}
