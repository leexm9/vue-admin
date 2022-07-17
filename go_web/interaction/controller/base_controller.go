package controller

import (
	"github.com/gin-gonic/gin"
	"go_web/pkg/resp"
	"net/http"
)

type BaseController struct {
}

type BaseResult struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var successCode = 200

func (r *BaseController) Success(c *gin.Context, data interface{}) {
	res := BaseResult{
		Code:    successCode,
		Success: true,
		Data:    data,
	}
	c.JSON(http.StatusOK, res)
}

func (r BaseController) Error(c *gin.Context, err resp.Err) {
	res := BaseResult{
		Code:    err.Code,
		Success: false,
		Msg:     err.Msg,
	}
	c.JSON(http.StatusOK, res)
}
