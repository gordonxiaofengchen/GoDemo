package model

import (
	"database/sql/driver"
	"errors"
	"time"
)

const (
	FORMATTER = "2006-01-02 15:04:05"
)

type Time struct{
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(FORMATTER)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, FORMATTER)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	parsed, err := time.Parse(`"`+FORMATTER+`"`, string(data))
	*t = Time{Time:parsed}
	return err
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch value.(type){
	case time.Time:
		t.Time = value.(time.Time)
		return nil
	case string:
		parse, err := time.Parse(FORMATTER, value.(string))
		if err != nil {
			return err
		}else{
			t.Time = parse
			return nil
		}
	default:
		return errors.New("invalid type, only time.Time or string is valid")
	}
}

func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t Time) String() string{
	return t.Format(FORMATTER)
}