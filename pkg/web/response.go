package web

import (
	"gomall-center/pkg/e"
)

type JsonResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *Context) Response(code int) {
	c.Context.JSON(e.SUCCESS, JsonResult{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}

func (c *Context) ResponseData(code int, data interface{}) {
	c.Context.JSON(e.SUCCESS, JsonResult{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}

func (c *Context) OK() {
	c.Context.JSON(e.SUCCESS, JsonResult{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
	})
}

func (c *Context) Success(code int) bool {
	// 业务代码只有返回200属于正常
	return code == 200
}
