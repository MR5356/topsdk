package response

import (
	"github.com/MR5356/topsdk/ability375/domain"
)

type TaobaoTbkTpwdCreateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果对象
	*/
	Data domain.TaobaoTbkTpwdCreateMapData `json:"data,omitempty" `
}
