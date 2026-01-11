package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule int

const weekSize = 7
const teenthDay = 13

const (
	First  WeekSchedule = 1
	Second              = First + weekSize
	Third               = First + weekSize*2
	Fourth              = First + weekSize*3
	Last                = First - weekSize
	Teenth              = teenthDay
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {

	defaultHour := 12
	defaultTick := 0

	if wSched == Last {
		month++
	}

	t := time.Date(year, month, int(wSched), defaultHour, defaultTick, defaultTick, defaultTick, time.UTC)

	return t.Day() + int(wDay-t.Weekday()+weekSize)%weekSize

}
