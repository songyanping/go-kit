package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusCode int

const (
	Error   StatusCode = 7
	Success StatusCode = 0

	//定义业务错误码
	UnknownError StatusCode = 100000
	UserNotFound StatusCode = 100001
)

type Response struct {
	Code StatusCode  `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code StatusCode, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(Success, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, "操作失败", c)
}

func FailCode(code StatusCode, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, message, c)
}

func FailWithMessageCode(code StatusCode, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Error, data, message, c)
}

func FailWithDetailedCode(code StatusCode, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Error,
		nil,
		message,
	})
}
