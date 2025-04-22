package middleware

import (
	"sync"
	"time"
)

// Bucket 定义漏桶结构体
type Bucket struct {
	capacity int        //桶的容量
	water    int        //当前水量
	leakRate int        //漏桶速率（每秒允许的请求数）
	lastTime time.Time  //上一次漏水时间
	mutex    sync.Mutex //互斥锁
}

// NewBucket 创建漏桶
func NewBucket(capacity, leakRate int) *Bucket {
	return &Bucket{
		capacity: capacity,
		leakRate: leakRate,
		lastTime: time.Now(),
	}
}

func (b *Bucket) leak() {
	now := time.Now()
	elapsed := now.Sub(b.lastTime).Seconds()
	leaked := int(elapsed * float64(b.leakRate))
	b.water = max(0, b.water-leaked)
	b.lastTime = now
}

func (b *Bucket) Allow() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.leak()
	if b.water < b.capacity {
		b.water++
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
