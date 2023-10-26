package dto

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}