package ability1042

import (
	"errors"
	"github.com/MR5356/topsdk"
	"github.com/MR5356/topsdk/ability1042/request"
	"github.com/MR5356/topsdk/ability1042/response"
	"github.com/MR5356/topsdk/util"
	"log"
)

type Ability1042 struct {
	Client *topsdk.TopClient
}

func NewAbility1042(client *topsdk.TopClient) *Ability1042 {
	return &Ability1042{client}
}

func (ability *Ability1042) TaobaoWirelessShareTpwdQuery(req *request.TaobaoWirelessShareTpwdQueryRequest) (*response.TaobaoWirelessShareTpwdQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1042 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.wireless.share.tpwd.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoWirelessShareTpwdQueryResponse{}
	if err != nil {
		log.Println("taobaoWirelessShareTpwdQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err

}
