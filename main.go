package gotime

import (
	"errors"
	"strconv"
	"unicode"
)

func toMillis(time string) (int, error) {
	if len(time) < 2 {
		return 0, errors.New("invalid time string")
	}

	numStr := time[:len(time)-1]
	unit := time[len(time)-1]

	number, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, errors.New("failed to parse numeric part of the time string")
	}

	switch unicode.ToLower(rune(unit)) {
	case 'w':
		return number * 7 * 24 * 60 * 60 * 1000, nil
	case 'd':
		return number * 24 * 60 * 60 * 1000, nil
	case 'h':
		return number * 60 * 60 * 1000, nil
	case 'm':
		return number * 60 * 1000, nil
	case 's':
		return number * 1000, nil
	default:
		return 0, errors.New("invalid time unit in the input string")
	}
}
