package utilities

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomDate time.Time

const customDateFormat = "2006-01-02"

// UnmarshalJSON parses a date string in the custom format.
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1:len(str)-1] // Trim the quotes
	t, err := time.Parse(customDateFormat, str)
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}
	*cd = CustomDate(t)
	return nil
}

// MarshalJSON formats a CustomDate as a JSON string.
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(cd).Format(customDateFormat))), nil
}

// Scan implements the sql.Scanner interface.
func (cd *CustomDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*cd = CustomDate(v)
	case string:
		t, err := time.Parse(customDateFormat, v)
		if err != nil {
			return fmt.Errorf("invalid date format: %w", err)
		}
		*cd = CustomDate(t)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (cd CustomDate) Value() (driver.Value, error) {
	return time.Time(cd).Format(customDateFormat), nil
}
