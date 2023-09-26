package controller

import (
	"application/src/usecase"
	"context"
	"github.com/gin-gonic/gin"
)

func GetContextFromGin(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, usecase.ClientAddress, c.ClientIP())
	ctx = context.WithValue(ctx, usecase.RequestID, c.Request.Header.Get("X-Request-ID"))
	return ctx
}

func ErrorHandlerToHTTPStatus(err error) (int, interface{}) {
	re, ok := err.(*usecase.ApplicationError)
	if ok {
		return re.StatusCode, re
	} else {
		return 500, struct {
			Code       int    `json:"code"`
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
			Err        error  `json:"error"`
		}{
			Code:       -1,
			StatusCode: 500,
			Message:    "unknown error",
			Err:        err,
		}
	}
}
