package civil

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-playground/errors/v5"
)

// UnmarshalJSON implements encoding/json Unmarshaler interface
func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("date should be a string, got %s", data)
	}

	val, err := ParseDate(s)
	if err != nil {
		return errors.Wrapf(err, "invalid date: %s", s)
	}
	*d = val

	return nil
}

// MarshalJSON implements encoding/json Marshaler interface
func (d *Date) MarshalJSON() ([]byte, error) {
	if y := d.Year; y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, fmt.Errorf("Date.MarshalJSON: year '%v' outside of range [0,9999]", y)
	}

	b := make([]byte, 0, 12)
	b = append(b, '"')
	b = append(b, d.String()...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements encoding/json Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("time should be a string, got %s", data)
	}

	val, err := ParseTime(s)
	if err != nil {
		return errors.Wrap(err, "invalid time")
	}
	*t = val

	return nil
}

// MarshalJSON implements encoding/json Marshaler interface
func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 20)
	b = append(b, '"')
	b = append(b, t.String()...)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON implements encoding/json Unmarshaler interface
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	tIdx := bytes.IndexAny(data, "Tt")
	if tIdx < 10 {
		return fmt.Errorf("data is not valid DateTime value: %s", string(data))
	}

	dataDate := data[:tIdx+1]
	dataDate[tIdx] = byte('"')
	dataTime := data[tIdx:]
	dataTime[0] = byte('"')

	if err := dt.Date.UnmarshalJSON(dataDate); err != nil {
		return errors.Wrapf(err, "date prefix (%s) in '%s' could not be converted", dataDate, data)
	}

	if err := dt.Time.UnmarshalJSON(dataTime); err != nil {
		return errors.Wrapf(err, "time suffix (%s) in '%s' could not be converted", dataTime, data)
	}

	return nil
}

// MarshalJSON implements encoding/json Marshaler interface
func (dt *DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 31)
	b = append(b, '"')
	b = append(b, dt.String()...)
	b = append(b, '"')

	return b, nil
}
