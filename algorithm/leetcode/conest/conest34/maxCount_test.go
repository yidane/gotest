package conest34

import "testing"

func Test_maxCount(t *testing.T) {
	type args struct {
		m   int
		n   int
		ops [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{m: 26, n: 17, ops: [][]int{{20, 10}, {26, 11}, {2, 11}, {4, 16}, {2, 3}, {23, 13}, {7, 15}, {11, 11}, {25, 13}, {11, 13}, {13, 11}, {13, 16}, {26, 17}}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCount(tt.args.m, tt.args.n, tt.args.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
