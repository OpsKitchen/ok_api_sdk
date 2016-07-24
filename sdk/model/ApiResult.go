package model

import "encoding/json"

type ApiResult struct {
	Success      bool        `json:"success"`
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func (r *ApiResult) ConvertDataTo(outPointer interface{}) error {
	responseDataBytes, _ := json.Marshal(r.Data)
	return json.Unmarshal(responseDataBytes, outPointer)
}
