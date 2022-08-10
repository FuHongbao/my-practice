package rank

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rank_list/api"
	"rank_list/services/rank"
)

func SetScore(ctx *gin.Context) {
	var reqParams api.SetScoreRequest
	if err := ctx.ShouldBindJSON(&reqParams); err != nil {
		ctx.JSON(http.StatusOK, &api.SetScoreResponse{
			Status:     0,
			ErrMsg:     "request bind json failed, please check request format!",
		})
		return
	}
	if err := rank.SetScoreForMouth(ctx, reqParams.TraceId, reqParams.Uid, reqParams.Score, reqParams.TimeStamp); err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.SetScoreResponse{
			TraceId:    reqParams.TraceId,
			Status:     0,
			ErrMsg:     "Internal servers error, please retry",
		})
		return
	}
	ctx.JSON(http.StatusOK, &api.SetScoreResponse{
		TraceId:    reqParams.TraceId,
		Status:  	1,
	})
	return
}


