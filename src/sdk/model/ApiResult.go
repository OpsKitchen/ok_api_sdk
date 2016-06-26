package model

type ApiResult struct {
	Success      *bool
	ErrorCode    string
	ErrorMessage string
	Data         interface{}
}
