package rank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rank_list/api"
	"rank_list/services/rank"
)

func GetRankInfo(ctx *gin.Context) {
	var reqParams api.GetRankInfoRequest
	if err := ctx.ShouldBindJSON(&reqParams); err != nil {
		ctx.JSON(http.StatusOK, &api.GetRankInfoResponse{
			Status:     0,
			ErrMsg:     "request bind json failed, please check request format!",
		})
		return
	}
	results, err := rank.GetRankInfoForMouth(ctx, reqParams.TraceId, reqParams.Uid, reqParams.DateYear, reqParams.DateMouth, reqParams.RangeNum)
	if err != nil {
		fmt.Println("rank.GetRankInfoForMouth failed, err==>", err)
		ctx.JSON(http.StatusInternalServerError, &api.GetRankInfoResponse{
			Status:     0,
			ErrMsg:     "Internal servers error, please retry",
		})
		return
	}
	if len(results) == 0 {
		ctx.JSON(http.StatusOK, &api.GetRankInfoResponse{
			Status:     0,
			ErrMsg:     "user's ranking info not exists",
		})
		return
	}
	var dest *api.RankInfoItem
	for _, val := range results {
		if val.Uid == reqParams.Uid {
			dest = val
		}
	}
	var resp = &api.GetRankInfoResponse{
		TraceId:      reqParams.TraceId,
		Status:       1,
		RankInfo:     dest,
		RankRange:    results,
	}
	ctx.JSON(http.StatusOK, resp)
	return
}
