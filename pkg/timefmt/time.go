package timefmt

import "time"

const (
	timeFmt = "2006-01-02 15:04:05"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFmt)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFmt)
	b = append(b, '"')
	return b, nil
}

func ParseInLocation(layout, value string, loc *time.Location) (*Time, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return nil, err
	}
	pt := Time(t)
	return &pt, nil
}
