package civil

import "time"

func (d Date) Format(s string) string {
	return d.In(time.UTC).Format(s)
}

func (dt DateTime) Format(s string) string {
	return dt.In(time.UTC).Format(s)
}
