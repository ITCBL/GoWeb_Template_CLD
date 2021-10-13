package models

// 定义请求参数的结构体
const ()

// ParamRegister 注册请求参数
type ParamRegister struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前用户的用户
	PostID string `json:"post_id" binding:"required"` // 帖子id
	// 因为int 默认就是0，如果填写的值存在0的整型类，不需要填写 binding:"required"，否则会报错。
	Direction int `json:"direction" binding:"oneof=1 0 -1"` // 赞成票（1）还是反对票（-1）,取消（0）
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct { // example 示例数据，将会在swagger 执行时,默认显示该值
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
