package utils

import "time"

func GenerateTime(hour, minute int, loc *time.Location) time.Time {
	t := time.Now().UTC()
	t = time.Date(t.Year(), t.Month(), t.Day(), hour, minute, 0, 0, t.Location()).In(loc)
	return t
}
