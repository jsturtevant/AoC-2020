package main

import "testing"

func Test_countValid(t *testing.T) {
	type args struct {
		rawInput []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				rawInput: []string{
					"1-3 a: abcde",
					"1-3 b: cdefg",
					"2-9 c: ccccccccc",
					"3-4 j: laksdfhh",
					"7-10 h: hhhhhhhhhhsd",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValid(tt.args.rawInput); got != tt.want {
				t.Errorf("countValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
