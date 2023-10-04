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

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func TraceRequest(ctx *gin.Context) {
	beforeRequest(ctx)

	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	ctx.Next()

	afterRequest(ctx, blw.Status(), blw.body.String())
}

func beforeRequest(ctx *gin.Context) {
	models.StartRequest(ctx)
	show.LogRequest(ctx)
}

func afterRequest(ctx *gin.Context, status int, body string) {
	models.StartResponse(status, body)
	show.LogResponse(ctx)
}
