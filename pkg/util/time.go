package util

import "time"

// order 处理完成时间段str转Time
func StringToTime(start string, end string) (time.Time, time.Time) {
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", start+" 00:00:00", time.Local)
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", end+" 23:59:59", time.Local)
	return startTime, endTime
}

// 字符串转时间格式
func OneStingToTime(str string) time.Time  {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	return t
}