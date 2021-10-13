package controller

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID" // 代码中不要出现莫名其妙的字符串，最好将这些字符串设置为Const常量。
var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey) // 代码中不要出现莫名其妙的字符串，最好将这些字符串设置为Const常量。
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	var (
		page     int64
		pageSize int64
		err      error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}

	// 限制pageSize大小,避免恶意查询
	if pageSize > 1000 {
		pageSize = 1000
	}
	return page, pageSize
}
