package api


type SetScoreRequest struct {
	TraceId			string		`json:"trace_id"`
	Uid  			string		`json:"uid"`
	Score			float64     `json:"score"`
	TimeStamp       int64		`json:"timestamp"`
}

type SetScoreResponse struct {
	TraceId			string		`json:"trace_id"`
	Status 			int		    `json:"status"`

	ErrMsg 			string		`json:"err_msg,omitempty"`
}

type GetRankInfoRequest struct {
	TraceId			string		`json:"trace_id"`
	Uid  			string		`json:"uid"`
	RangeNum 		int  		`json:"range_num"`  //查询前后各多少个人
	DateYear	    int		    `json:"date_year"`
	DateMouth	    int  		`json:"date_mouth"`
	TimeStamp       int64		`json:"timestamp"`
}

type GetRankInfoResponse struct {
	TraceId			string			`json:"trace_id"`
	Status 			int		    	`json:"status"`
	RankInfo		*RankInfoItem 	`json:"rank_info"`
	RankRange       []*RankInfoItem	`json:"rank_range"`

	ErrMsg 			string			`json:"err_msg,omitempty"`
}

type RankInfoItem struct {
	Uid  			string		`json:"uid"`
	RankPos 		int64 		`json:"rank_pos"`
	Score 			float64   	`json:"score"`
}