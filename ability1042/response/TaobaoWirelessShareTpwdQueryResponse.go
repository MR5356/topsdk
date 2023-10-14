package response

import "github.com/MR5356/topsdk/ability1042/domain"

type TaobaoWirelessShareTpwdQueryResponse struct {
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
	Data domain.TaobaoWirelessShareTpwdQueryData `json:"data,omitempty" `
}
