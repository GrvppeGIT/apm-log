package models

type Response struct {
	Body   string
	Status int
}

var response Response

func StartResponse(statusCode int, body string) {
	response = Response{Body: body, Status: statusCode}
}

func GetResponse() Response {
	return response
}
