package rank

import (
	"context"
	"rank_list/dao/redis"
	"time"
)

func SetScoreForMouth(ctx context.Context, traceId, uid string, score float64, timestamp int64) error {
	scoreFloat := calculateScore(score, timestamp)
	tm := time.Unix(timestamp / 1000, 0)
	return redis.SetScore(redis.RankMouthKey(tm.Year(), int(tm.Month())), scoreFloat, uid)
}

