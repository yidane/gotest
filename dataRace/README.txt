避免Data Race的两种方式

1、加锁 详细解决方式如 dataRace2
2、使用chan chan方式 思路就是让数据操作按照顺序执行 详细解决方式如 dataRace3