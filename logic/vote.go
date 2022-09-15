package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

//投票功能
//1.用户投票的数据

//使用简化版的投票分数
//投一票就加432分,正常情况下随时间流逝
//60*60*24=86400;86400/200=432分,200张赞成票就可以给帖子续一天

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
