package sdk

type ApiResult struct {
	Success   *bool
	ErrorCode string
	Data      interface{}
}

func (r *ApiResult) IsSuccess() bool {
	return r.Success
}

func (r *ApiResult) GetErrorCode() string {
	return r.ErrorCode
}

func (r *ApiResult) GetData() interface{} {
	return r.Data
}