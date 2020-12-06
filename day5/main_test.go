package main

import (
	"reflect"
	"testing"
)

const input = `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

func Test_findSeat(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "given set return largest name",
			args: args{
				input: input,
			},
			want: "BBFFBBFRLL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSeat(tt.args.input); !reflect.DeepEqual(got.Name(), tt.want) {
				t.Errorf("findSeat() = %v, want %v", got.Name(), tt.want)
			}
		})
	}
}

func Test_seat_ID(t *testing.T) {
	type fields struct {
		row string
		col string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "FBFBBFFRLR",
			fields: fields{
				row: "FBFBBFF",
				col: "RLR",
			},
			want: 357,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &seat{
				row: tt.fields.row,
				col: tt.fields.col,
			}
			if got := s.ID(); got != tt.want {
				t.Errorf("seat.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
