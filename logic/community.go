package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查数据库并返回给前端
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
