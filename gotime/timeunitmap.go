package gotime

var timeUnitMultipliers = map[rune]int{
	'w': millisInWeek,   // weeks
	'd': millisInDay,    // days
	'h': millisInHour,   // hours
	'm': millisInMinute, // minutes
	's': millisInSecond, // seconds
	'S': 1,              // milliseconds
}
