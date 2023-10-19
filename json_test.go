package civil

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Date
		wantErr bool
	}{
		{
			name: "valid date",
			args: args{
				data: []byte(`"2021-01-01"`),
			},
			want: &Date{
				Year:  2021,
				Month: time.January,
				Day:   1,
			},
			wantErr: false,
		},
		{
			name: "invalid date",
			args: args{
				data: []byte(`"2021-01-32"`),
			},
			want:    &Date{},
			wantErr: true,
		},
		{
			name: "invalid json",
			args: args{
				data: []byte(`2021-01-30`),
			},
			want:    &Date{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Date{}
			if err := got.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.UnmarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	type fields struct {
		Year  int
		Month time.Month
		Day   int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid date",
			fields: fields{
				Year:  2021,
				Month: time.January,
				Day:   1,
			},
			want: []byte(`"2021-01-01"`),
		},
		{
			name: "year to large",
			fields: fields{
				Year:  20021,
				Month: time.January,
				Day:   1,
			},
			wantErr: true,
		},
		{
			name: "year to small",
			fields: fields{
				Year:  -1,
				Month: time.January,
				Day:   1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Year:  tt.fields.Year,
				Month: tt.fields.Month,
				Day:   tt.fields.Day,
			}
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.MarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Time
		wantErr bool
	}{
		{
			name: "valid time",
			args: args{
				data: []byte(`"12:03:10"`),
			},
			want: &Time{
				Hour:   12,
				Minute: 3,
				Second: 10,
			},
		},
		{
			name: "invalid time",
			args: args{
				data: []byte(`"12:00:60"`),
			},
			want:    &Time{},
			wantErr: true,
		},
		{
			name: "invalid json",
			args: args{
				data: []byte(`12:00:00`),
			},
			want:    &Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &Time{}
			if err := got.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Time.UnmarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Hour       int
		Minute     int
		Second     int
		Nanosecond int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid time",
			fields: fields{
				Hour:   12,
				Minute: 3,
				Second: 10,
			},
			want: []byte(`"12:03:10"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Time{
				Hour:       tt.fields.Hour,
				Minute:     tt.fields.Minute,
				Second:     tt.fields.Second,
				Nanosecond: tt.fields.Nanosecond,
			}
			got, err := tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Time.MarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDateTime_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *DateTime
		wantErr bool
	}{
		{
			name: "valid date time",
			args: args{
				data: []byte(`"2021-01-01T12:03:10"`),
			},
			want: &DateTime{
				Date: Date{
					Year:  2021,
					Month: time.January,
					Day:   1,
				},
				Time: Time{
					Hour:   12,
					Minute: 3,
					Second: 10,
				},
			},
		},
		{
			name: "invalid date format",
			args: args{
				data: []byte(`"2021-01-01 12:03:10"`),
			},
			want:    &DateTime{},
			wantErr: true,
		},
		{
			name: "invalid date",
			args: args{
				data: []byte(`"2021-01-32T12:03:10"`),
			},
			want:    &DateTime{},
			wantErr: true,
		},
		{
			name: "invalid time",
			args: args{
				data: []byte(`"2021-01-01T12:03:60"`),
			},
			want: &DateTime{
				Date: Date{
					Year:  2021,
					Month: time.January,
					Day:   1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &DateTime{}
			if err := got.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DateTime.UnmarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDateTime_MarshalJSON(t *testing.T) {
	t.Parallel()

	type fields struct {
		Date Date
		Time Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid date time",
			fields: fields{
				Date: Date{
					Year:  2021,
					Month: time.January,
					Day:   1,
				},
				Time: Time{
					Hour:   12,
					Minute: 3,
					Second: 10,
				},
			},
			want: []byte(`"2021-01-01T12:03:10"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &DateTime{
				Date: tt.fields.Date,
				Time: tt.fields.Time,
			}
			got, err := dt.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DateTime.MarshalJSON() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
