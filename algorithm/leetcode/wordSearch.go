package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/word-search/#/description

Given a 2D board and a word, find if the word exists in the grid.
The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
For example,
Given board =

[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
word = "ABCCED", -> returns true,
word = "SEE", -> returns true,
word = "ABCB", -> returns false.
*/
//*************************************************************************

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	if len(board) == 0 {
		return false
	}
	rows := len(board)
	columns := len(board[0])
	firstChar := word[0]

	subFind := func(usedAee [][]bool, x int, y int, c byte) bool {
		return true
	}

	for i := 0; i < rows; i++ {
		x := i
		arr := board[x]
		if len(arr) == 0 {
			continue
		}
		for j := 0; j < columns; j++ {
			y := j
			if arr[y] == firstChar {
				charIndex := 1
				used := [][]bool{}
				used[i][j] = true
				for len(word) > charIndex {
					subFind(used, x, y, word[charIndex])
					charIndex++
				}
			}
		}
	}

	return false
}
