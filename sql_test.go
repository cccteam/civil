package civil

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDate_Value(t *testing.T) {
	t.Parallel()

	type fields struct {
		Year  int
		Month time.Month
		Day   int
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "valid date",
			fields: fields{
				Year:  2020,
				Month: 2,
				Day:   29,
			},
			want: "2020-02-29",
		},
		{
			name: "zero date",
			want: "0000-00-00",
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
			got, err := d.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.Value() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDate_Scan(t *testing.T) {
	t.Parallel()

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Date
		wantErr bool
	}{
		{
			name: "valid date string",
			args: args{
				value: "2020-02-29",
			},
			want: &Date{
				Year:  2020,
				Month: 2,
				Day:   29,
			},
		},
		{
			name: "valid date time.Time",
			args: args{
				value: time.Date(2020, time.February, 29, 0, 0, 0, 0, time.UTC),
			},
			want: &Date{
				Year:  2020,
				Month: 2,
				Day:   29,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := &Date{}
			if err := got.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Date.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.Scan() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTime_Value(t *testing.T) {
	t.Parallel()

	type fields struct {
		Hour       int
		Minute     int
		Second     int
		Nanosecond int
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "valid time",
			fields: fields{
				Hour:       3,
				Minute:     42,
				Second:     31,
				Nanosecond: 876,
			},
			want: "03:42:31.000000876",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := Time{
				Hour:       tt.fields.Hour,
				Minute:     tt.fields.Minute,
				Second:     tt.fields.Second,
				Nanosecond: tt.fields.Nanosecond,
			}
			got, err := tr.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Time.Value() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTime_Scan(t *testing.T) {
	t.Parallel()

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Time
		wantErr bool
	}{
		{
			name: "valid time string",
			args: args{
				value: "03:42:31.000000876",
			},
			want: &Time{
				Hour:       3,
				Minute:     42,
				Second:     31,
				Nanosecond: 876,
			},
		},
		{
			name: "valid time time.Time",
			args: args{
				value: time.Date(2020, time.February, 29, 3, 42, 31, 876, time.UTC),
			},
			want: &Time{
				Hour:       3,
				Minute:     42,
				Second:     31,
				Nanosecond: 876,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := &Time{}
			if err := got.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Time.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.Scan() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDateTime_Value(t *testing.T) {
	t.Parallel()

	type fields struct {
		Date Date
		Time Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "valid datetime",
			fields: fields{
				Date: Date{
					Year:  2020,
					Month: 2,
					Day:   29,
				},
				Time: Time{
					Hour:       3,
					Minute:     42,
					Second:     31,
					Nanosecond: 876,
				},
			},
			want: "2020-02-29T03:42:31.000000876",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dt := DateTime{
				Date: tt.fields.Date,
				Time: tt.fields.Time,
			}
			got, err := dt.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DateTime.Value() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDateTime_Scan(t *testing.T) {
	t.Parallel()

	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *DateTime
		wantErr bool
	}{
		{
			name: "valid datetime string",
			args: args{
				value: "2020-02-29T03:42:31.000000876",
			},
			want: &DateTime{
				Date: Date{
					Year:  2020,
					Month: 2,
					Day:   29,
				},
				Time: Time{
					Hour:       3,
					Minute:     42,
					Second:     31,
					Nanosecond: 876,
				},
			},
			wantErr: false,
		},
		{
			name: "valid datetime time.Time",
			args: args{
				value: time.Date(2020, time.February, 29, 3, 42, 31, 876, time.UTC),
			},
			want: &DateTime{
				Date: Date{
					Year:  2020,
					Month: 2,
					Day:   29,
				},
				Time: Time{
					Hour:       3,
					Minute:     42,
					Second:     31,
					Nanosecond: 876,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := &DateTime{}
			if err := got.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.Scan() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
