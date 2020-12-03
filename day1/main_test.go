package main

import "testing"

func Test_find2020(t *testing.T) {
	type args struct {
		expenseReport map[int]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given test",
			args: args{
				expenseReport: map[int]bool{
					1721: true,
					979:  true,
					366:  true,
					299:  true,
					675:  true,
					1456: true,
				},
			},
			want: 514579,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find2020(tt.args.expenseReport); got != tt.want {
				t.Errorf("find2020() = %v, want %v", got, tt.want)
			}
		})
	}
}
