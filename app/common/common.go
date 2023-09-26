package common

type Response struct {
	StatusCode int
	Message    string
}

type SuccessResponse struct {
	Response
	Data any
}

type FailResponse struct {
	Response
	Error error
}
