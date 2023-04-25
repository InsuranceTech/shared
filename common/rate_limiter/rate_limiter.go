package rate_limiter

import (
	"fmt"
	"sync"
	"time"
)

type Limiter interface {
	Wait()
	WaitCancellable(c *CancellableLimiter) (Cancelled bool)
}

type limiter struct {
	tick    time.Duration
	count   uint
	entries []time.Time
	index   uint
	mutex   sync.Mutex
}

type dateLimiter struct {
	limiterType dateLimiterType
	count       uint
	index       uint
	dateHash    string
	mutex       sync.Mutex
}

type CancellableLimiter struct {
	cancelled bool
}

func (c *CancellableLimiter) Cancel() {
	c.cancelled = true
}

func (c *CancellableLimiter) IsCancelled() bool {
	return c.cancelled
}

type dateLimiterType int32

const (
	type_minute dateLimiterType = 1
	type_second dateLimiterType = 2
)

// var BinanceSpotWebRateLimit = NewLimiter("1m", 1100) // 1m/1200 limit
var BinanceSpotWebRateLimit = NewMinuteLimiter(1100) // 1m/1200 limit

func NewLimiter(durationStr string, count uint) Limiter {
	tick, err := time.ParseDuration(durationStr)
	if err != nil {
		panic("Error Parse duration : " + durationStr)
	}
	return NewLimiterDuration(tick, count)
}

func NewLimiterDuration(tick time.Duration, count uint) Limiter {
	l := limiter{
		tick:  tick,
		count: count,
		index: 0,
	}
	l.entries = make([]time.Time, count)
	before := time.Now().Add(-2 * tick)
	for i, _ := range l.entries {
		l.entries[i] = before
	}
	return &l
}

func (l *limiter) Wait() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	last := &l.entries[l.index]
	next := last.Add(l.tick)
	now := time.Now()
	if now.Before(next) {
		time.Sleep(next.Sub(now))
	}
	*last = time.Now()
	l.index = l.index + 1
	if l.index == l.count {
		l.index = 0
	}
}

func (l *dateLimiter) Wait() {
	l.WaitCancellable(nil)
}

func (l *limiter) WaitCancellable(c *CancellableLimiter) (Cancelled bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	last := &l.entries[l.index]
	next := last.Add(l.tick)
	now := time.Now()
	if now.Before(next) {
		nextTime := time.Now().Add(next.Sub(now))
		for time.Now().UnixMilli() < nextTime.UnixMilli() {
			if c.cancelled == true {
				return true
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
	*last = time.Now()
	l.index = l.index + 1
	if l.index == l.count {
		l.index = 0
	}
	return c.cancelled
}

func (l *dateLimiter) WaitCancellable(c *CancellableLimiter) (Cancelled bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.limiterType == type_minute {
		curDateHash := time.Now().Format(time.RFC822)
		if l.dateHash != curDateHash {
			l.dateHash = curDateHash
			l.index = 1
		}

		if l.index >= l.count {
			// Bir sonraki dakikaya kadar bekle
			sec := 60 - time.Now().Second() + 2
			fmt.Println("MIN: Rate limit Bekleniyor", sec, " saniye")
			time.Sleep(time.Duration(sec) * time.Second)
		}
	} else if l.limiterType == type_second {
		curDateHash := time.Now().Format(time.RFC850)
		if l.dateHash != curDateHash {
			l.dateHash = curDateHash
			l.index = 1
		}

		if l.index >= l.count {
			// Bir sonraki saniyeye kadar bekle
			ns := int(time.Second) - time.Now().Nanosecond() + int(time.Second/3)
			dr := time.Duration(int64(ns))
			if c == nil {
				fmt.Println("SEC: Rate limit Bekleniyor", dr)
				time.Sleep(dr)
			} else {
				nextTime := time.Now().Add(dr)
				for time.Now().UnixMilli() < nextTime.UnixMilli() {
					if c.cancelled == true {
						return true
					}
					time.Sleep(time.Millisecond * 10)
				}
			}
		}
	}

	l.index++
	return false
}

func NewMinuteLimiter(count uint) Limiter {
	l := dateLimiter{
		count:       count,
		limiterType: type_minute,
	}
	return &l
}

func NewSecondLimiter(count uint) Limiter {
	l := dateLimiter{
		count:       count,
		limiterType: type_second,
	}
	return &l
}
