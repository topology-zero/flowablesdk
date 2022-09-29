package timefmt

import "time"

// RFC3339ms RFC3339 标准,毫秒显示
const (
	RFC3339ms = "2006-01-02T15:04:05.999Z07:00"

	timeFmt = "2006-01-02 15:04:05"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFmt+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFmt)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFmt)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFmt)
}
