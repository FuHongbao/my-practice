package redis

import "github.com/go-redis/redis"

func SetScore(key string, score float64, member string) error {
	_, err := rdb.store.ZAdd(key, redis.Z{
		Score:  score,
		Member: member,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}


