package leetcode

/*
Given an integer, convert it to a roman numeral.
Input is guaranteed to be within the range from 1 to 3999.
*/

var romaMap = map[int]map[int]string{
	1:    {0: "", 1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"},
	10:   {0: "", 1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC"},
	100:  {0: "", 1: "C", 2: "CC", 3: "CCC", 4: "CD", 5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM"},
	1000: {0: "", 1: "M", 2: "MM", 3: "MMM"}}

func intToRoman(num int) string {
	return romaMap[1000][num/1000] + romaMap[100][(num%1000)/100] + romaMap[10][(num%100)/10] + romaMap[1][num%10]
}
