package convert

import (
	"strconv"
	"time"
)

const (
	//ISO8601 is format for ISO8601 dates. Like RFC3339, but without timezone
	ISO8601 = "2006-01-02T15:04:05"
)

//ToBool tries to convert string into bool type
func ToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

//ToInt tries to convert string into int type
func ToInt(value string) (int, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	return int(i), err
}

//ToFloat tries to convert string into float64 type
func ToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

//ToDate tries to convert string into time.Time type
func ToDate(value string) (time.Time, error) {
	//is serial format?
	serial, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return time.Unix(int64((serial-25569)*86400), 0), nil
	}

	return time.Parse(ISO8601, value)
}
