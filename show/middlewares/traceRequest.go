package middlewares

import (
	"bytes"

	"github.com/GrvppeGIT/apm-log/show/models"
	"github.com/GrvppeGIT/apm-log/show/services"

	"github.com/gin-gonic/gin"
)

var show services.Show

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var blw *bodyLogWriter

func TraceRequest(ctx *gin.Context) {
	beforeRequest(ctx)
	ctx.Next()
	afterRequest(ctx)
}

func beforeRequest(ctx *gin.Context) {
	models.StartRequest(ctx)

	blw = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	show.LogRequest(ctx)

}

func afterRequest(ctx *gin.Context) {
	models.StartResponse(blw.Status(), blw.body.String())

	show.LogResponse(ctx)
}
