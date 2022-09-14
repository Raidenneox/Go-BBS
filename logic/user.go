package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在,存在没必要往下走
	if err := mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}
	//2.生成UID
	userID := snowflake.GenID()

	//构造一个User实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//4.保存进数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (u *models.User, err error) {
	u = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针
	if err := mysql.Login(u); err != nil {
		return nil, err
	}
	//生成JWT的Token
	token, err := jwt.GenToken(u.UserID, u.Password)
	if err != nil {
		return
	}
	u.Token = token
	return
}
