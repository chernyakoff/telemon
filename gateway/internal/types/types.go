// Code generated by goctl. DO NOT EDIT.
package types

type MessageResponse struct {
	Message string `json:"message"`
}

type AccountUpsertRequest struct {
	Phone   string `json:"phone"`
	AppId   int    `json:"appid"`
	AppHash string `json:"apphash"`
	Session string `json:"session"`
}