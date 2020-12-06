package main

import "testing"

const input = `abc

a
b
c

ab
ac

a
a
a
a

b`

func Test_sumOfAnswers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given",
			args: args{input: input},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfAnswers(tt.args.input); got != tt.want {
				t.Errorf("sumOfAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}
