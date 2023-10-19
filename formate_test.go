package civil

import (
	"testing"
	"time"
)

func TestDate_Format(t *testing.T) {
	t.Parallel()

	type fields struct {
		Year  int
		Month time.Month
		Day   int
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "valid date",
			fields: fields{
				Year:  2021,
				Month: time.February,
				Day:   3,
			},
			args: args{
				s: "2006-01-02",
			},
			want: "2021-02-03",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			if got := d.Format(tt.args.s); got != tt.want {
				t.Errorf("Date.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_Format(t *testing.T) {
	t.Parallel()

	type fields struct {
		Date Date
		Time Time
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "valid date time",
			fields: fields{
				Date: Date{
					Year:  2021,
					Month: time.February,
					Day:   3,
				},
				Time: Time{
					Hour:   12,
					Minute: 3,
					Second: 10,
				},
			},
			args: args{
				s: "2006-01-02T15:04:05",
			},
			want: "2021-02-03T12:03:10",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := DateTime{
				Date: tt.fields.Date,
				Time: tt.fields.Time,
			}
			if got := d.Format(tt.args.s); got != tt.want {
				t.Errorf("DateTime.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
