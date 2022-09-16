package models

import "time"

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`                            // 帖子id
	AuthorID    int64     `json:"author_id" db:"author_id"`                          // 作者id
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` // 社区id
	Status      int32     `json:"status" db:"status"`                                // 帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               // 帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           // 帖子内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      // 帖子创建时间
}

type APIPostDetail struct {
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"` //投票数/分数 后期可以应用其它算法来进行计算
	*Post                               //帖子信息
	*CommunityDetail `json:"community"` //分类信息
}
