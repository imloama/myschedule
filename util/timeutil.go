package util

import (
	"time"
)

const Format_All = "2006-01-02 15:04:05"
const YYYY_MM_DD = "2006-01-02"

func Format(t time.Time) string {
	return t.Format(Format_All)
}

func Parse(t string) (time.Time, error) {
	return time.Parse(Format_All, t)
}

func Format_Y_M_D(t time.Time) string {
	return t.Format(YYYY_MM_DD)
}
