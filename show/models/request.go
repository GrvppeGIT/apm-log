package models

import (
	"fmt"
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
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(body)
	return string(body)
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
