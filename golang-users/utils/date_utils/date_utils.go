package date_utils

import (
	"log"
	"strconv"
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

func AddHoursToUnixToString(in int) string {
	return strconv.FormatInt(AddHours(in).Unix(), 10)
}
func StringToUnixDate(in string) (int64, error) {
	return strconv.ParseInt(in, 10, 64)
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

func IsExpiredFromString(in string) bool {
	expire, err := StringToUnixDate(in)
	if err != nil {
		log.Println("Conversion Error " + err.Error())
		return true
	}
	now := GetNow().Unix()
	return expire < now
}
