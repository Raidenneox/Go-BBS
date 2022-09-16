package redis

//redis key
//将程序中运行时固定不变的一些key设置为常量

//Redis key 注意使用命名空间的方式区分不同的key，方便查询和拆分
const (
	Prefix             = "bluebell:"
	KeyPostTimeZSet    = "post:time"   //zset;帖子及发帖时间
	KeyScorePostZSet   = "post:score"  //zset;帖子及其分数
	KeyPostVotedZSetPF = "post:voted:" //zset;记录用户及其投票类型;参数是post id

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
)

//给redis key 加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
