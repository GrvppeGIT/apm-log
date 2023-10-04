package logger

import (
	"encoding/json"
	"log"

	"github.com/GrvppeGIT/apm-log/logger/models"
	"github.com/GrvppeGIT/apm-log/logger/utils"

	"github.com/gin-gonic/gin"
)

type Printer struct {
	log     models.Log
	dt      utils.DateTime
	context string
}

func (p *Printer) Initialize(log models.Log, dt utils.DateTime) {
	p.log = log
	p.dt = dt
	p.context = ""
}

func (p *Printer) SetContext(className string) {
	p.context = className
}

func (p *Printer) SetCnpj(cnpj string) {
	p.log.Cnpj = cnpj
}

func (p *Printer) Log(message string) {
	p.setArguments("info", message)
}

func (p *Printer) LogRequest(ctx *gin.Context) {
	p.setArgumentsRequest(ctx)
}

func (p *Printer) LogResponse(ctx *gin.Context) {
	p.setArgumentsResponse(ctx)
}

func (p *Printer) Debug(message string) {
	p.setArguments("debug", message)
}

func (p *Printer) Warn(message string) {
	p.setArguments("warn", message)
}

// func (p *Printer) Error(message string) {

// }

// ===========================
// next, private functions ==>
// ===========================

func (p *Printer) setArguments(level string, message string) {
	p.log.Timestamp = p.dt.GetDatetime()
	p.log.Message = message
	p.log.LogLevel = level
	p.log.LogLogger = p.context

	p.stdout()
}

func (p *Printer) setArgumentsRequest(ctx *gin.Context) {
	req := models.GetRequest()

	p.log.Timestamp = p.dt.GetDatetime()
	p.log.Cnpj = utils.GetCnpj(req.Auth)
	p.log.Message = "Request recebido!"
	p.log.LogLevel = "info"
	p.log.LogLogger = p.context

	p.log.SetHttpRequest(req.Body, req.Method, req.Referrer)

	p.stdout()
	p.log.ResetHttpRequest()
}

func (p *Printer) setArgumentsResponse(ctx *gin.Context) {
	res := models.GetResponse()

	p.log.Timestamp = p.dt.GetDatetime()
	p.log.Message = "Response recebido!"
	p.log.LogLogger = p.context
	p.log.LogLevel = p.getLevelHttp(res.Status)

	p.log.SetHttpResponse(res.Body, res.Status)

	p.stdout()
	p.resetArgumentsLog()
}

func (p *Printer) getLevelHttp(status int) string {
	if status >= 100 && status < 399 {
		return "info"
	}

	if status >= 400 && status < 499 {
		return "warn"
	}

	if status >= 500 && status < 699 {
		return "error"
	}

	return "info"

}

func (p *Printer) resetArgumentsLog() {
	p.context = ""

	p.log.Timestamp = ""
	p.log.Cnpj = ""
	p.log.Message = ""
	p.log.LogLevel = "info"
	p.log.LogLogger = p.context

	p.log.ResetHttpRequest()
	p.log.ResetHttpResponse()
}

func (p *Printer) stdout() {
	print, _ := json.Marshal(p.log)
	log.Println(string(print))
}
