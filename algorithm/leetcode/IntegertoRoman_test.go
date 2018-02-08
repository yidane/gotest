package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{num: 2}, want: "II"},
		{name: "2", args: args{num: 12}, want: "XII"},
		{name: "3", args: args{num: 4}, want: "IV"},
		{name: "4", args: args{num: 29}, want: "XXIX"},
		{name: "5", args: args{num: 49}, want: "XLIX"},
		{name: "6", args: args{num: 6}, want: "VI"},
		{name: "7", args: args{num: 101}, want: "CI"},
		{name: "8", args: args{num: 110}, want: "CX"},
		{name: "9", args: args{num: 1001}, want: "MI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}

	// romaCache()
	// t.Error()
}
