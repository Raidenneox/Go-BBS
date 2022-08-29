package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const salt = "nnkpassword"

//把每一步数据库操作封装成函数

//等待logic层根据业务需要调用

var (
	ErrorUserExist     = errors.New("用户已存在")
	ErrorUserNotExist  = errors.New("无此用户")
	ErrorWrongPassword = errors.New("用户名或密码错误")
)

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 像数据库中插入一条新的记录
func InsertUser(u *models.User) (err error) {
	//!!!不能在数据库中存储明文的密码,存加密后的
	u.Password = encryptPassword(u.Password)

	//执行SQL语句入库
	sqlStr := "insert into user(user_id,username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, u.UserID, u.Username, u.Password)
	return err
}

func encryptPassword(o string) string {
	h := md5.New()
	h.Write([]byte(salt))
	h.Sum([]byte(o))
	return hex.EncodeToString(h.Sum([]byte(o)))
}

func Login(user *models.User) (err error) {
	encryptedPwd := encryptPassword(user.Password)
	sqlStr1 := "select user_id,username,password from user where username = ?"
	err = db.Get(user, sqlStr1, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		//查询数据库失败
		return err
	}
	if encryptedPwd != user.Password {
		return ErrorWrongPassword
	}
	return

}
