package redis

import "web_app/models"

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis 获取ID
	//1. 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyScorePostZSet)
	}

	//确定查询的索引的起始点
	start := (p.Page - 1) * p.Size //第一页就从0开始
	end := start + p.Size - 1

	//3,ZREVRANGE按分数大从大到小的顺序查询指定数量大小的元素
	return rdb.ZRevRange(key, start, end).Result()
}
