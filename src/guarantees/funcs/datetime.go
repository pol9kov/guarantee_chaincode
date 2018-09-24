package funcs

import (
	"time"
	"strconv"
)

func GetTime() string {
	then :=  time.Date(
		2018, 1, 1, 0, 0, 0, 0, time.UTC)

	now := time.Now()

	diff := now.Sub(then)
	result := diff.Nanoseconds()

	return strconv.FormatInt(result, 10)
}