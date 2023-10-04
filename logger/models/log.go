package models

type Log struct {
	Timestamp               string `json:"@timestamp,omitempty"` //ok
	Cnpj                    string `json:"cnpj,omitempty"`       //ok
	Message                 string `json:"message,omitempty"`    //ok
	LogLevel                string `json:"log.level,omitempty"`  //ok
	LogLogger               string `json:"log.logger,omitempty"` //ok
	TraceId                 string `json:"trace.id,omitempty"`
	TransactionId           string `json:"transaction.id,omitempty"`
	ServiceName             string `json:"service.name,omitempty"`               //ok
	ServiceVersion          string `json:"service.version,omitempty"`            //ok
	HttpRequestBodyContent  string `json:"http.request.body.content,omitempty"`  //ok
	HttpRequestMethod       string `json:"http.request.method,omitempty"`        //ok
	HttpRequestReferrer     string `json:"http.request.referrer,omitempty"`      //ok
	HttpResponseBodyContent string `json:"http.response.body.content,omitempty"` //ok
	HttpResponseStatusCode  int    `json:"http.response.status_code,omitempty"`  //ok
	ErrorMessage            string `json:"error.message,omitempty"`
	ErrorStackTrace         string `json:"error.stack_trace,omitempty"`
}

func (l *Log) SetCnpj(cnpj string) {
	l.Cnpj = cnpj
}

// next, functions with more than one param ******

func (l *Log) SetService(name string, version string) {
	l.ServiceName = name
	l.ServiceVersion = version
}

func (l *Log) SetHttpRequest(content string, method string, referrer string) {
	l.HttpRequestBodyContent = content
	l.HttpRequestMethod = method
	l.HttpRequestReferrer = referrer
}

func (l *Log) ResetHttpRequest() {
	l.HttpRequestBodyContent = ""
	l.HttpRequestMethod = ""
	l.HttpRequestReferrer = ""
}

func (l *Log) SetHttpResponse(content string, status_code int) {
	l.HttpResponseBodyContent = content
	l.HttpResponseStatusCode = status_code
}

func (l *Log) ResetHttpResponse() {
	l.HttpResponseBodyContent = ""
	l.HttpResponseStatusCode = 0
}

func (l *Log) SetError(message string, stack_trace string) {
	l.ErrorMessage = message
	l.ErrorStackTrace = stack_trace
}
