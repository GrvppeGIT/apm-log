package logger

import (
	"bytes"

	"github.com/GrvppeGIT/apm-log/logger/models"

	"github.com/gin-gonic/gin"
)

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
	MainLog.Printer.logRequest(ctx)
}

func afterRequest(ctx *gin.Context, status int, body string) {
	models.StartResponse(status, body)
	MainLog.Printer.logResponse(ctx)
}
