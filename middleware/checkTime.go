package middleware

import "time"

// 假设这是秒杀活动的开始时间和结束时间
var (
	startTime = time.Date(2024, 12, 12, 10, 0, 0, 0, time.Local)
	endTime   = time.Date(2024, 12, 12, 11, 0, 0, 0, time.Local)
)

// isInActivityTime 函数用于检查当前时间是否在活动时间区间内
func isInActivityTime() bool {
	now := time.Now()
	return now.After(startTime) && now.Before(endTime)
}
