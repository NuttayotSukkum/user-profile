package handler

type Code string

const (
	CodeSuccess       Code = "1000"
	CodeCreated       Code = "2000"
	CodeBadRequest    Code = "4000"
	CodeUnauthorized  Code = "8000"
	CodeForbidden     Code = "8001"
	CodeNotFound      Code = "4400"
	CodeConflict      Code = ""
	CodeInternalError Code = ""
)

type MessageResponse struct {
	Code Code   `json:"status"`
	Time string `json:"time"`
	Data any    `json:"data"`
}
