package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成postID
	p.ID = snowflake.GenID()
	//保存到数据库
	return mysql.CreatePost(p)
	//3.返回
}
