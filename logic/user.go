package logic

import (
	"GoWeb_Template_CLD/common/jwt"
	"GoWeb_Template_CLD/common/snowflake"
	"GoWeb_Template_CLD/dao/mysql"
	"GoWeb_Template_CLD/models"
)

// 存放业务逻辑的代码

func Register(p *models.ParamRegister) (err error) {
	// 1.判断用户是否存在
	err = mysql.CheckoutUserExist(p.Username)
	if err != nil {
		return err
	}
	// 2.生成UUID
	genID := snowflake.GenID()
	// 构造user信息
	u := &models.User{
		UserID:   genID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存到数据库
	err = mysql.InsertUser(u)
	return
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.UserName,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
