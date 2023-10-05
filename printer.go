package apmlog

import (
	"encoding/json"
	"log"

	"github.com/GrvppeGIT/apm-log/models"
	"github.com/GrvppeGIT/apm-log/utils"

	"github.com/gin-gonic/gin"
)

type Printer struct {
	log     models.Log
	dt      utils.DateTime
	context string
}

func (p *Printer) SetContext(className string) {
	p.context = className
}

func (p *Printer) SetCnpj(cnpj string) {
	p.log.Cnpj = cnpj
}

func (p *Printer) SetTracer(traceId string, transactionId string) {
	p.log.TraceId = traceId
	p.log.TransactionId = transactionId
}

func (p *Printer) ResetTracer() {
	p.log.ResetTracers()
}

func (p *Printer) Log(message string, context ...string) {
	p.setArguments("info", message, context)
}

func (p *Printer) Debug(message string, context ...string) {
	p.setArguments("debug", message, context)
}

func (p *Printer) Warn(message string, context ...string) {
	p.setArguments("warn", message, context)
}

// func (p *Printer) Error(message string) {

// }

// ===========================
// next, private functions ==>
// ===========================

func (p *Printer) initialize(log models.Log, dt utils.DateTime) {
	p.log = log
	p.dt = dt
	p.context = ""
}

func (p *Printer) logRequest(ctx *gin.Context) {
	p.setArgumentsRequest(ctx)
}

func (p *Printer) logResponse(ctx *gin.Context) {
	p.setArgumentsResponse(ctx)
}

func (p *Printer) setArguments(level string, message string, context []string) {
	p.log.Timestamp = p.dt.GetDatetime()
	p.log.Message = message
	p.log.LogLevel = level
	p.addContext(context)

	p.stdout()
}

func (p *Printer) setArgumentsRequest(ctx *gin.Context) {
	req := models.GetRequest()

	p.log.Timestamp = p.dt.GetDatetime()
	p.log.Cnpj = utils.GetCnpj(req.Auth)
	p.log.Message = "Request recebido!"
	p.log.LogLevel = "info"

	p.log.SetHttpRequest(req.Body, req.Method, req.Referrer)

	p.stdout()
	p.log.ResetHttpRequest()
}

func (p *Printer) setArgumentsResponse(ctx *gin.Context) {
	res := models.GetResponse()

	p.log.Timestamp = p.dt.GetDatetime()
	p.log.Message = "Response recebido!"
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

func (p *Printer) addContext(context []string) {
	if len(context) > 0 {
		p.log.LogLogger = context[0]
	} else if p.context != "" {
		p.log.LogLogger = p.context
	}
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
