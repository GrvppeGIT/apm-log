package apmlog

import (
	"github.com/gin-gonic/gin"
)

type OptionsAPM struct {
	Engine      *gin.Engine
	ServiceName string
	ServerUrl   string
	SecretKey   string
}

var ApmMain ApmService

func StartAPM(opt OptionsAPM) {

	if !SetEnvAPM(opt.ServiceName, opt.ServerUrl, opt.SecretKey) {
		panic("error when defining APM environment variables")
	}

	ApmMain.Initialize()

	MainLog.Printer.Log("started APM server")

}
