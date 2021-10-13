package controller

import (
	"GoWeb_Template_CLD/common"
	"GoWeb_Template_CLD/logic"

	"GoWeb_Template_CLD/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// Register 处理注册请求的函数
func (u *UserController) Register(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamRegister) // new 返回一个指针类型
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Register with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok { // 如果不是validator.ValidationErrors类型,则不能翻译处理，直接原格式错误。
			ResponseError(c, common.CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, common.CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
		return // 调用c.json()后，必须return,否则会继续执行下去
	}
	// 2.业务处理
	if err := logic.Register(p); err != nil {
		zap.L().Error("logic.Register failed!", zap.Error(err))
		if errors.Is(err, common.ErrorUserExist) {
			ResponseError(c, common.CodeUserExist)
			return
		}
		ResponseError(c, common.CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, nil)
}

// Login 登录
func (u *UserController) Login(c *gin.Context) {
	// 获取参数和参数校验
	var p models.ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型。不是：直接返回错误信息；是：返回参数校验类型错误的翻译结果
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, common.CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, common.CodeInvalidParam, common.RemoveTopStruct(errs.Translate(common.Trans)))
		return
	}
	// 业务处理
	user, err := logic.Login(&p)
	if err != nil {
		// 一般在controller层进行错误日志打印（错误信息来自logic/dao等返回）
		zap.L().Error("logic.Login failed", zap.String("username", p.UserName), zap.Error(err)) // zap.String("username", p.UserName) 将用户名也打印出来，方便排查。
		if errors.Is(err, common.ErrorUserNotExist) {
			ResponseError(c, common.CodeUserNotExist)
			return
		}
		ResponseError(c, common.CodeInvalidPassword)
		return
	}
	// 返回响应
	ResponseSuccess(c, gin.H{
		// js支持的最大整数值：1>>53-1；go支持最大的整数值1>>63-1,避免失真。
		"user_id":   fmt.Sprintf("%d", user.UserID), // 转成string类型给前端
		"user_name": user.Username,
		"token":     user.Token,
	})
}
