package helpers

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JSONTime time.Time

func CreateDateTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := time.Now().In(loc)

	res := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d +0300", date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second())
	datetime, err := time.Parse("2006-01-02 15:04:05 -0700", string(res))
	PanicIfError(err)
	return datetime

}

func CreateDate() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	return now.Format("2006-01-02")

}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	tm, err := time.Parse("2006-01-02", s)
	PanicIfError(err)
	*t = JSONTime(tm)
	return nil
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t))
}

/* null time */

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
