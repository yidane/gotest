package leetcode

import (
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'}, want: 'c'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'}, want: 'f'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'}, want: 'f'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'g'}, want: 'j'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'j'}, want: 'c'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'k'}, want: 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreatestLetter2(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'}, want: 'c'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'}, want: 'f'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'}, want: 'f'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'g'}, want: 'j'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'j'}, want: 'c'},
		{args: args{letters: []byte{'c', 'f', 'j'}, target: 'k'}, want: 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter2(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter2() = %v, want %v", got, tt.want)
			}
		})
	}
}
