package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
	"rank_list/api"
	"strconv"
)

//拉取目标前后各n个对象的排名的情况，共2n + 1个结果
func GetItemsRangeN(key string, member string, n int) ([]*api.RankInfoItem, error) {
	keys := []string{key, member}
	args := []string{
		strconv.Itoa(n),
	}
	resp, err := rdb.store.Eval(itemsRangeNScript, keys, args).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	data := cast.ToSlice(resp)
	var startPos int64
	var uidArr []string
	var scoresArr []float64
	for _, item := range data {
		switch val := item.(type) {
		case int64:
			startPos = val + 1
		case int:
			startPos = int64(val) + 1
		case []interface{}:
			for ind, obj := range val {
				if strObj, ok := obj.(string); ok {
					if ind % 2 == 0 {
						uidArr = append(uidArr, strObj)
					} else {
						floatObj, err := strconv.ParseFloat(strObj, 10)
						if err != nil {
							return nil, err
						}
						scoresArr = append(scoresArr, floatObj)
					}
				}
			}
		default:
			return nil, errors.New("zrevrange from redis failed, data type is not support")
		}
	}
	if len(scoresArr) == 0 || len(scoresArr) != len(uidArr) {
		return nil, errors.New("zrevrange from redis failed, data is wrong")
	}
	var results []*api.RankInfoItem
	for ind, uid := range uidArr {
		result := &api.RankInfoItem{
			Uid:     uid,
			RankPos: startPos + int64(ind),
			Score:   scoresArr[ind],
		}
		results = append(results, result)
	}
	return results, nil
}

const itemsRangeNScript = `local n = tonumber(ARGV[1])
local rank = tonumber(redis.call("zrevrank", KEYS[1], KEYS[2]))
if rank == nil then
	return nil
end

local left = 0
if rank - n > left then
	left = rank - n
end
local right = rank + n
local listItems = redis.call("zrevrange", KEYS[1], left, right, "withscores")
local ret = {}
table.insert(ret, listItems)
table.insert(ret, left)
return ret`

/*

for ind, val in pairs(listItems)  do
	table.insert(ret, val)
end
*/