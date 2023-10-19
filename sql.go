package civil

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Value implements the database/sql/driver valuer interface.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan implements the database/sql scanner interface.
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		val, err := ParseDate(v)
		if err != nil {
			return err
		}
		*d = val
	case []byte:
		val, err := ParseDate(string(v))
		if err != nil {
			return err
		}
		*d = val
	case time.Time:
		val := DateOf(v)
		*d = val
	default:
		return fmt.Errorf("'%v' could not be converted into a valid type", value)
	}

	return nil
}

// Value implements the database/sql/driver valuer interface.
func (t Time) Value() (driver.Value, error) {
	return t.String(), nil
}

// Scan implements the database/sql scanner interface.
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		val, err := ParseTime(v)
		if err != nil {
			return err
		}
		*t = val
	case []byte:
		val, err := ParseTime(string(v))
		if err != nil {
			return err
		}
		*t = val
	case time.Time:
		val := TimeOf(v)
		*t = val
	default:
		return fmt.Errorf("'%v' could not be converted into a valid type", value)
	}

	return nil
}

// Value implements the database/sql/driver valuer interface.
func (dt DateTime) Value() (driver.Value, error) {
	return dt.String(), nil
}

// Scan implements the database/sql scanner interface.
func (dt *DateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		val, err := ParseDateTime(v)
		if err != nil {
			return err
		}
		*dt = val
	case []byte:
		val, err := ParseDateTime(string(v))
		if err != nil {
			return err
		}
		*dt = val
	case time.Time:
		val := DateTimeOf(v)
		*dt = val
	default:
		return fmt.Errorf("'%v' could not be converted into a valid type", value)
	}

	return nil
}
