package main

import "testing"

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
		input        string
		preambleSize int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "given",
			args: args{
				input:        input,
				preambleSize: 5,
			},
			want: "127",
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
