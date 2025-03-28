package gotime

import (
	"fmt"
	"strconv"
	"unicode"
)

// ToMillis converts a time string (e.g., "1h", "30m", "24h") to milliseconds.
// Supported units: w (weeks), d (days), h (hours), m (minutes), s (seconds), S (milliseconds)
func ToMillis(time string) (int, error) {
	length := len(time)
	if length < 2 {
		return 0, ErrInvalidLength
	}

	unit := unicode.ToLower(rune(time[length-1]))
	multiplier, ok := timeUnitMultipliers[unit]
	if !ok {
		return 0, fmt.Errorf("%w: %c", ErrInvalidUnit, unit)
	}

	numStr := time[:length-1]
	number, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrInvalidFormat, numStr)
	}

	if number < 0 {
		return 0, ErrNegativeValue
	}

	return number * multiplier, nil
}
