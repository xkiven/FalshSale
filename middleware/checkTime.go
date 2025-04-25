package middleware

import (
	"FlashSale/models"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义 Redis 中活动时间的键
const (
	ActivityStartTimeKey = "activity:start_time:%d"
	ActivityEndTimeKey   = "activity:end_time:%d"
	TimeFormat           = "2006-01-02T15:04:05Z07:00"
)

// 从 Redis 获取活动时间
func getActivityTimeFromRedis(ctx context.Context, rdb *redis.Client, key string) (string, error) {
	result, err := rdb.Get(ctx, key).Result()
	if err != nil && err.Error() != "redis: nil" {
		log.Printf("Failed to get activity time from Redis: %v", err)
		return "", err
	}
	return result, nil
}

// 设置活动时间到 Redis
func setActivityTimeToRedis(ctx context.Context, rdb *redis.Client, key string, value string, duration time.Duration) error {
	err := rdb.Set(ctx, key, value, duration).Err()
	if err != nil {
		log.Printf("Failed to set activity time to Redis: %v", err)
		return err
	}
	return nil
}

// 解析时间字符串
func parseTime(timeStr string) (time.Time, error) {
	t, err := time.Parse(TimeFormat, timeStr)
	if err != nil {
		log.Printf("Failed to parse time: %v", err)
		return time.Time{}, err
	}
	return t, nil
}

// IsInActivityTime 函数用于检查当前时间是否在特定活动的时间区间内
func IsInActivityTime(ctx context.Context, db *gorm.DB, rdb *redis.Client, activityID int32) (bool, error) {
	// 生成特定活动的 Redis 键
	startTimeKey := fmt.Sprintf(ActivityStartTimeKey, activityID)
	endTimeKey := fmt.Sprintf(ActivityEndTimeKey, activityID)

	// 尝试从 Redis 中获取活动的开始时间和结束时间
	startTimeStr, err := getActivityTimeFromRedis(ctx, rdb, startTimeKey)
	if err != nil {
		return false, err
	}
	endTimeStr, err := getActivityTimeFromRedis(ctx, rdb, endTimeKey)
	if err != nil {
		return false, err
	}

	// 如果 Redis 中没有缓存活动时间，从数据库中查询
	if startTimeStr == "" || endTimeStr == "" {
		var activity models.Activity
		result := db.Where("id = ?", activityID).First(&activity)
		if result.Error != nil {
			log.Printf("Failed to get activity from database: %v", result.Error)
			return false, result.Error
		}

		startTimeStr = activity.StartTime
		endTimeStr = activity.EndTime

		// 将活动时间缓存到 Redis 中，并设置合理的过期时间
		cacheDuration := time.Hour * 1 // 例如设置为 1 小时
		err = setActivityTimeToRedis(ctx, rdb, startTimeKey, startTimeStr, cacheDuration)
		if err != nil {
			return false, err
		}
		err = setActivityTimeToRedis(ctx, rdb, endTimeKey, endTimeStr, cacheDuration)
		if err != nil {
			return false, err
		}
	}

	// 解析活动的开始时间和结束时间
	startTime, err := parseTime(startTimeStr)
	if err != nil {
		return false, err
	}
	endTime, err := parseTime(endTimeStr)
	if err != nil {
		return false, err
	}

	log.Printf("Activity start time: %s, end time: %s", startTime, endTime)

	// 获取当前时间
	now := time.Now()

	// 检查当前时间是否在活动时间区间内
	return now.After(startTime) && now.Before(endTime), nil
}
