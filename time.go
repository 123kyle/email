package email

import (
	"fmt"
	"time"
)

var utcFormat = "2006-01-02 15:04:05 -0700"

type UTCTime struct {
	time.Time
}

func (u UTCTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, u.Format(utcFormat))), nil
}

func (t *UTCTime) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.Parse(fmt.Sprintf(
		`"%s"`, utcFormat), string(data))
	return err
}
