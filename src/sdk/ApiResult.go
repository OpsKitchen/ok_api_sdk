package sdk

type ApiResult struct {
	Success   *bool
	ErrorCode string
	Data      interface{}
}