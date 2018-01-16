package charts

const (
	TenkanPeriod       = 20
	KijunPeriod        = 60
	SenkouBPeriod      = 120
	CloudLeadingPeriod = 30
)

var OneDayInterval string = "1D"
var FourHourInterval string = "4h"
var OneHourInterval string = "1h"
var ThirtyMinuteInterval string = "30m"
var MinuteSecond = 60

var ValidIntervals = map[string]bool{
	OneDayInterval:       true,
	FourHourInterval:     true,
	OneHourInterval:      true,
	ThirtyMinuteInterval: true,
}

var IntervalMilliseconds = map[string]int{
	OneDayInterval:       MinuteSecond * 60 * 24,
	FourHourInterval:     MinuteSecond * 60 * 4,
	OneHourInterval:      MinuteSecond * 60,
	ThirtyMinuteInterval: MinuteSecond * 30,
}
