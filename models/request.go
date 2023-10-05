package models

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Auth     string
	Body     string
	Method   string
	Referrer string
}

var request Request

func StartRequest(ctx *gin.Context) {
	request = Request{Auth: getToken(ctx), Method: ctx.Request.Method, Referrer: getReferrer(ctx), Body: getBody(ctx)}
}

func GetRequest() Request {
	return request
}

func getReferrer(ctx *gin.Context) string {
	schema := "http"

	if ctx.Request.TLS != nil {
		schema = "https"
	}

	return schema + "://" + ctx.Request.Host + ctx.Request.URL.Path
}

func getBody(ctx *gin.Context) string {
	bodyCopy := new(bytes.Buffer)
	_, err := io.Copy(bodyCopy, ctx.Request.Body)

	if err != nil {
		return ""
	}

	bodyData := bodyCopy.Bytes()
	ctx.Request.Body = io.NopCloser(bytes.NewReader(bodyData))

	return string(bodyData)

}

func getToken(ctx *gin.Context) string {
	const Bearer_schema = "Bearer "
	header := ctx.GetHeader("Authorization")

	if header == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return ""
	}

	return header[len(Bearer_schema):]

}
