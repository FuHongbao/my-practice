package rank

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	group := r.Group("/rank")
	group.POST("/set_score", SetScore)
	group.POST("/get_rank_info", GetRankInfo)
}