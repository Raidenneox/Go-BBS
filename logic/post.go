package logic

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 1. 生成post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	err := mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err2 := redis.CreatePost(p.ID, p.CommunityID)
	if err2 != nil {
		fmt.Println("写入redis出错")
	}
	return err2
	// 3. 返回
}

// GetPostByID 根据帖子id查询帖子详情
func GetPostByID(pid int64) (data *models.APIPostDetail, err error) {
	//查询并组合接口想要的数据
	dataPost, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Error(err))
		return
	}
	//根据作者id查询作者信息
	user, err := mysql.GetUserByID(dataPost.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(dataPost.AuthorID) failed", zap.Error(err))
		return
	}
	//根据社区ID查询社区相关信息
	cmy, err := mysql.GetCommunityDetailByID(dataPost.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(dataPost.AuthorID) failed", zap.Error(err))
		return
	}
	//接口数据拼接
	data = &models.APIPostDetail{
		AuthorName:      user.Username,
		Post:            dataPost,
		CommunityDetail: cmy,
	}
	return
}

func GetPostList(page, size int64) (data []*models.APIPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	//有多少个帖子就有多少个详情
	data = make([]*models.APIPostDetail, 0, len(posts))

	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(dataPost.AuthorID) failed", zap.Error(err))
			continue
		}
		//根据社区ID查询社区相关信息
		cmy, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(dataPost.AuthorID) failed", zap.Error(err))
			continue
		}
		postDetail := &models.APIPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: cmy,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.APIPostDetail, err error) {
	//2.去redis查询ID值
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	//3.根据ID去MySQL数据库查询帖子详细信息
	//返回的数据还要按照给定的id顺序返回
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostInOrder success but no value existing")
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//提前查询好每篇帖子的投票数
	votaData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(dataPost.AuthorID) failed", zap.Error(err))
			continue
		}
		//根据社区ID查询社区相关信息
		cmy, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(dataPost.AuthorID) failed", zap.Error(err))
			continue
		}
		postDetail := &models.APIPostDetail{
			AuthorName:      user.Username,
			VoteNum:         votaData[idx],
			Post:            post,
			CommunityDetail: cmy,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.APIPostDetail, err error) {

	//2.去redis查询ID值
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	//3.根据ID去MySQL数据库查询帖子详细信息
	//返回的数据还要按照给定的id顺序返回
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostInOrder success but no value existing")
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//提前查询好每篇帖子的投票数
	votaData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(dataPost.AuthorID) failed", zap.Error(err))
			continue
		}
		//根据社区ID查询社区相关信息
		cmy, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(dataPost.AuthorID) failed", zap.Error(err))
			continue
		}
		postDetail := &models.APIPostDetail{
			AuthorName:      user.Username,
			VoteNum:         votaData[idx],
			Post:            post,
			CommunityDetail: cmy,
		}
		data = append(data, postDetail)
	}
	return
}
