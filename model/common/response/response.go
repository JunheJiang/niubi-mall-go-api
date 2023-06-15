package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ResultCode int         `json:"resultCode"`
	Data       interface{} `json:"data"`
	Msg        string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
	UNLOGIN = 416
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code, data, msg,
	})
}

// Ok  --- map[string]interface{}{} 空键值对
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// OkWithMsg --- interface{} obj任意类型 空接口 、 map[string]interface{}{}
func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "SUCCESS", c)
}
func OkWithDetail(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func FailWithMsg(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetail(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func UnLogin(data interface{}, c *gin.Context) {
	Result(UNLOGIN, data, "未登录！", c)
}
