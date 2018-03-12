package main

import "testing"

func Test_countDate(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2017-1-1", args: args{year: 2017, month: 1, day: 1}, want: 1},
		{name: "2017-2-1", args: args{year: 2017, month: 2, day: 1}, want: 32},
		{name: "2017-3-1", args: args{year: 2017, month: 3, day: 1}, want: 60},
		{name: "2000-1-1", args: args{year: 2000, month: 1, day: 1}, want: 1},
		{name: "2000-3-1", args: args{year: 2000, month: 3, day: 1}, want: 61},
		{name: "1996-3-1", args: args{year: 2000, month: 3, day: 1}, want: 61},
		{name: "1998-3-1", args: args{year: 1998, month: 3, day: 1}, want: 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDate(tt.args.year, tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("countDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
