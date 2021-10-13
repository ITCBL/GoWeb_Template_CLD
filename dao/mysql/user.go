package mysql

import (
	"GoWeb_Template_CLD/common"
	"GoWeb_Template_CLD/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// 把每一步数据库操作封装成函数，待logic层根据业务需求调用
const secret = "hello"

// CheckoutUserExist 查询用户是否存在
func CheckoutUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return
	}
	if count > 0 {
		/*todo 利用全局的常量
		return errors.New("用户已存在")*/
		return common.ErrorUserExist
	}
	return
}

// InsertUser 新增用户数据
func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}
func encryptPassword(oldPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(oldPassword)))
}

// Login 用户登录
func Login(user *models.User) (err error) {
	oldPassword := user.Password // 保存登录密码（未加密）

	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	// 先判断是不是数据库为空
	if err == sql.ErrNoRows {
		// 一般情况下，不会直接返回“用户名不存在”，而是返回“用户名或密码错误”。
		// 因为直接返回“用户名不存在”，别人就知道该用户名未注册等，可能进行攻击之类的。
		/*todo 利用全局的常量
		return errors.New("用户不存在")*/
		return common.ErrorUserNotExist
	}
	// 查询数据库失败
	if err != nil {
		return err
	}
	// 判断密码是否正确

	if encryptPassword(oldPassword) != user.Password {
		/*todo 利用全局的常量
		return errors.New("密码错误")*/
		return common.ErrorInvalidPassword
	}
	return
}
func GetUserById(userId int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id=?`
	err = db.Get(user, sqlStr, userId)
	return
}
