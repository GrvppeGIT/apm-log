package apm

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

type OptionsAPM struct {
	engine      *gin.Engine
	serviceName string
	serverUrl   string
	secretKey   string
}

func StartAPM(opt OptionsAPM) {

	SetEnvAPM(opt.serviceName, opt.serverUrl, opt.secretKey)

	opt.engine.Use(apmgin.Middleware(opt.engine))

	// apm.set

}
