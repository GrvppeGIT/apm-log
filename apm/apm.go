package apm

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

type OptionsAPM struct {
	Engine      *gin.Engine
	ServiceName string
	ServerUrl   string
	SecretKey   string
}

func StartAPM(opt OptionsAPM) {

	SetEnvAPM(opt.ServiceName, opt.ServerUrl, opt.SecretKey)

	opt.Engine.Use(apmgin.Middleware(opt.Engine))

	// apm.set

}
