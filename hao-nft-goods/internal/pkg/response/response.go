package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess      = 0
	CodeError        = -1
	CodeInvalidParam = 4000
)

// Response 是接口统一的响应结构，Code 表示业务状态码。
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// JSON 使用 HTTP 200 返回业务结果；无数据时 data 为 {}。
func JSON(c *gin.Context, code int, message string, data any) {
	if data == nil {
		data = struct{}{}
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Success(c *gin.Context, data any) {
	JSON(c, CodeSuccess, "success", data)
}

func Error(c *gin.Context, code int, message string) {
	JSON(c, code, message, nil)
}
