package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}

func AddHoursDBFormat(hours int) string {
	return AddHours(hours).Format(apiDbLayout)
}

func AddHours(hours int) time.Time {
	timein := add(hours, 0, 0)
	return timein
}

func add(hours, mins, secs int) time.Time {
	timein := GetNow().Add(time.Hour*time.Duration(hours) +
		time.Minute*time.Duration(mins) +
		time.Second*time.Duration(secs))
	return timein
}

func IsExpired(expire int64) bool {
	now := GetNow().Unix()
	return expire < now
}
