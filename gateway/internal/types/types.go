// Code generated by goctl. DO NOT EDIT.
package types

type MessageResponse struct {
	Message string `json:"message"`
}

type AccountUpsertRequest struct {
	Phone   int64 `json:"Phone"`
	AppId   int64  `json:"AppId"`
	AppHash string `json:"AppHash"`
	Session string `json:"Session"`
}
