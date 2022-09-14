package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成postID
	p.ID = snowflake.GenID()
	//保存到数据库
	return mysql.CreatePost(p)
	//3.返回
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
