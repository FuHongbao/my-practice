package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const (
	mouthPrefix = "_rankM"
	dayPrefix   = "_rankD"
)
func RankMouthKey(yy, mm int) string {
	return fmt.Sprintf("%s:%d:%d", mouthPrefix, yy, mm)
}

func RankDayKey(yy, mm, dd int) string {
	return fmt.Sprintf("%s:%d:%d:%d", dayPrefix, yy, mm, dd)
}


type RDB struct {
	store 	*redis.Client
}
var rdb *RDB

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		MaxRetries:         2,
		DialTimeout:        time.Millisecond * 100,
		ReadTimeout:        time.Millisecond * 100,
		WriteTimeout:       time.Millisecond * 100,
	})
	if client == nil {
		panic("redis connect failed, conn is nil")
		return
	}
	rdb = &RDB{store: client}
	go healthCheck()
}

func healthCheck() {
	timeTick := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-timeTick.C:
			_, err := rdb.store.Ping().Result()
			if err != nil {
				client := redis.NewClient(&redis.Options{
					Addr:               "127.0.0.1:6379",
					MaxRetries:         2,
					DialTimeout:        time.Millisecond * 100,
					ReadTimeout:        time.Millisecond * 100,
					WriteTimeout:       time.Millisecond * 100,
				})
				if client != nil {
					rdb = &RDB{store: client}
				}
			}
		}
	}
}