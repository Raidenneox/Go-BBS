package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

/*投票的几种情况
direction=1时,有两种情况:
	1.之前没有投过票,现在投赞成票 +1 +432
	2.之前投反对票,现在改投赞成票 +2 +864

direction=0时,有两种情况:
	1.之前投反对,现在要取消 +1 +432
	2.之前投赞成,现在要取消 -1 -432

direction=-1时,有两种情况:
	1.之前没有投过票,现在投反对票 -1 -432
	2.之前投赞成票,现在改投反对票 -2 -432

投票的限制:
每个帖子自发表之日起一个星期之内允许用户投票,超过一个星期就不允许投票了
	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2.到期之后删除那个保存的KeyPostVoteZSet
*/

const (
	oneWeekInSeconds = 7 * 24 * 60 * 60
	scorePerVote     = 432 //每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID int64) error {

	pipeline := rdb.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyScorePostZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制
	//去redis取贴子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2与3需要放进一个pipeline事务中

	//2.更新分数
	//先查之前的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值
	pipeline := rdb.TxPipeline()

	pipeline.ZIncrBy(getRedisKey(KeyScorePostZSet), op*diff*scorePerVote, postID)
	//3.记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, //当前用户投的是赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
