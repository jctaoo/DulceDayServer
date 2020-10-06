package common

type BaseResponse struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
}
