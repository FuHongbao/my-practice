package rank

import (
	"context"
	"rank_list/api"
	"rank_list/dao/redis"
)

func GetRankInfoForMouth(ctx context.Context, traceId, uid string, yy, mm int, num int) ([]*api.RankInfoItem, error) {
	rankInfo, err :=  redis.GetItemsRangeN(redis.RankMouthKey(yy, mm), uid, num)
	if err != nil {
		return nil, err
	}
	if rankInfo == nil {
		return nil, nil
	}
	//去除小数点后的值
	for _, val := range rankInfo {
		val.Score = reductionScore(val.Score)
	}
	return rankInfo, nil
}
