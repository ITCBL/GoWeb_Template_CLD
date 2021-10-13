package controller

import (
	"GoWeb_Template_CLD/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
   "code": 10000, // 程序中的错误码
	"msg": xx,    // 提示信息
	"data":{},    // 响应数据
}
*/

type ResponseData struct {
	Code common.ResCode `json:"code"` // 程序中的错误码
	Msg  interface{}    `json:"msg"`  // 提示信息
	// todo 如果Data为空，不希望返回{data:null}。选择直接不返回，使用 omitempty(忽略空值的字段)
	Data interface{} `json:"data,omitempty"` // 响应数据
}

// 上面的结构体，也可以直接使用gin.H{}来定义，也可以

// ResponseError 错误响应
func ResponseError(c *gin.Context, code common.ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseErrorWithMsg 自定义的错误响应
func ResponseErrorWithMsg(c *gin.Context, code common.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: common.CodeSuccess,
		Msg:  common.CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
