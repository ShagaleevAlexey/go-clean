package time

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type CustomTime struct {
	time.Time
}

func CustomTimeFromTime(t time.Time) CustomTime {
	return CustomTime{
		Time: t,
	}
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-01-02T15:04:05.999999-07:00", string(b[1:len(b)-1]))
	if err != nil {
		t, err = time.Parse("2006-01-02T15:04:05.999999", string(b[1:len(b)-1]))
		if err != nil {
			return err
		}
	}

	ct.Time = t
	return nil
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time)
}

func Now() CustomTime {
	return CustomTimeFromTime(time.Now())
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *CustomTime) Scan(input interface{}) error {
	ct.Time = input.(time.Time)
	return nil
}
