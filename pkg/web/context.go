package web

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	RequestContext *RequestContext
}

type HandlerFunc func(*Context)

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		context.RequestContext = &RequestContext{
			Host:  c.Request.Host,
			Token: c.GetHeader("Authorization"),
		}
		handler(context)
	}
}
