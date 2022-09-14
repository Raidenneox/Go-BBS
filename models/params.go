package models

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票参数
type ParamVoteData struct {
	//UserID 从请求中直接获取当前用户
	PostID    string `json:"post_id" binding:"required"`                       //防止前端数据失真
	Direction int8   `json:"direction,string" binding:"required,oneof=1 -1 0"` //帖子赞成票(1)还是反对票(-1)取消投票(0)
}
