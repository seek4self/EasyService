package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// code message
const (
	ERROR   = 500
	SUCCESS = 199 + iota
	DBErr
	InternalErr
)

// codeMsg ...
var codeMsg map[int]string = map[int]string{
	SUCCESS:     "OK",
	DBErr:       "DB operating Error",
	InternalErr: "Service internal error",
}

var Resp = response{}

type response struct{}

// Data ...
type Data struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func (r response) Ok(c *gin.Context) {
	c.JSON(http.StatusOK, Data{SUCCESS, codeMsg[SUCCESS], struct{}{}})
}

func (r response) OkWithData(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Data{SUCCESS, codeMsg[SUCCESS], data})
}

func (r response) OkWithMsg(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Data{code, msg, struct{}{}})
}

func (r response) OkWithCode(code int, c *gin.Context) {
	c.JSON(http.StatusOK, Data{code, codeMsg[code], struct{}{}})
}

func (r response) BadRequest(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Data{400, msg, struct{}{}})
}
