package util

import "time"

type TimeWithoutTZ struct {
	time.Time
}

func (t *TimeWithoutTZ) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	tt, err := time.Parse("\"2006-01-02T15:04:05\"", string(data))
	*t = TimeWithoutTZ{tt}
	return err
}
