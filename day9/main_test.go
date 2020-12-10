package main

import (
	"testing"
)

const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func Test_run(t *testing.T) {
	type args struct {
		input        []int
		preambleSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given",
			args: args{
				input:        parseInput(input),
				preambleSize: 5,
			},
			want: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.input, tt.args.preambleSize); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_crack(t *testing.T) {
	type args struct {
		value int
		xmas  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given cracked example",
			args: args{
				value: 127,
				xmas:  parseInput(input),
			},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crack(tt.args.value, tt.args.xmas); got != tt.want {
				t.Errorf("crack() = %v, want %v", got, tt.want)
			}
		})
	}
}
