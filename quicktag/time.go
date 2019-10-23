package quicktag

import (
	"time"
)

type myTime time.Time

// MarshalJSON MarshalJSON
func (mt myTime) MarshalJSON() ([]byte, error) {
	t := time.Time(mt)
	return []byte(t.Format("\"2006-01-02 15:04:05\"")), nil
}

// UnmarshalJSON UnmarshalJSON
func (mt *myTime) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation("\"2006-01-02 15:04:05\"", string(data), time.Local)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}
