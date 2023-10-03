package services

import (
	"encoding/json"
	"log"

	"github.com/GrvppeGIT/apm-log/show/models"
	"github.com/GrvppeGIT/apm-log/show/utils"

	"github.com/gin-gonic/gin"
)

type Show struct {
	show    models.Log
	dt      utils.DateTime
	context string
}

func (s *Show) Initialize(show models.Log, dt utils.DateTime) {
	s.show = show
	s.dt = dt
	s.context = ""
}

func (s *Show) SetContext(className string) {
	s.context = className
}

func (s *Show) SetCnpj(cnpj string) {
	s.show.Cnpj = cnpj
}

func (s *Show) Log(message string) {
	s.setArguments("info", message)
}

func (s *Show) LogRequest(ctx *gin.Context) {
	s.setArgumentsRequest(ctx)
}

func (s *Show) LogResponse(ctx *gin.Context) {
	s.setArgumentsResponse(ctx)
}

func (s *Show) Debug(message string) {
	s.setArguments("debug", message)
}

func (s *Show) Warn(message string) {
	s.setArguments("warn", message)
}

// func (s *Show) Error(message string) {

// }

// ===========================
// next, private functions ==>
// ===========================

func (s *Show) setArguments(level string, message string) {
	s.show.Timestamp = s.dt.GetDatetime()
	s.show.Message = message
	s.show.LogLevel = level
	s.show.LogLogger = s.context

	s.stdout()
}

func (s *Show) setArgumentsRequest(ctx *gin.Context) {
	req := models.GetRequest()

	s.show.Timestamp = s.dt.GetDatetime()
	s.show.Cnpj = utils.GetCnpj(req.Auth)
	s.show.Message = "O request chegou com sucesso!"
	s.show.LogLevel = "info"
	s.show.LogLogger = s.context

	s.show.SetHttpRequest(req.Body, req.Method, req.Referrer)

	s.stdout()
}

func (s *Show) setArgumentsResponse(ctx *gin.Context) {
	res := models.GetResponse()

	s.show.Timestamp = s.dt.GetDatetime()
	s.show.Message = "Retorno da requisição com sucesso! (response) - para mais detalhes, verificar http.response.body.content"
	s.show.LogLogger = s.context
	s.show.LogLevel = s.getLevelHttp(res.Status)

	s.show.SetHttpResponse(res.Body, res.Status)

	s.stdout()
	s.resetArgumentsLog()
}

func (s *Show) getLevelHttp(status int) string {
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

func (s *Show) resetArgumentsLog() {
	s.context = ""

	s.show.Timestamp = ""
	s.show.Cnpj = ""
	s.show.Message = ""
	s.show.LogLevel = "info"
	s.show.LogLogger = s.context

	s.show.ResetHttpRequest()
	s.show.ResetHttpResponse()
}

func (s *Show) stdout() {
	p, _ := json.Marshal(s.show)
	log.Println(string(p))
}
