package main

func countDate(year, month, day int) int {
	if month == 1 {
		return day
	}
	if month == 2 {
		return 31 + day
	}

	monthDays := []int{31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	dayIndex := 31 + 28 + day
	//闰年二月份加1
	if (year%4 == 0 && year%100 != 0) || (year%100 == 0 && year%400 == 0) {
		dayIndex++
	}

	//汇总二月份后天数
	for i := 0; i < month-3; i++ {
		dayIndex += monthDays[i]
	}

	return dayIndex
}
