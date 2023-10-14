package domain


type TaobaoTbkDgMaterialOptionalSpCampaign struct {
    /*
        定向计划活动ID     */
    SpCid  *string `json:"sp_cid,omitempty" `

    /*
        定向计划名称     */
    SpName  *string `json:"sp_name,omitempty" `

    /*
        定向佣金率     */
    SpRate  *string `json:"sp_rate,omitempty" `

    /*
        定向是否锁佣，0=不锁佣 1=锁佣     */
    SpLockStatus  *string `json:"sp_lock_status,omitempty" `

    /*
        定向计划申请链接     */
    SpApplyLink  *string `json:"sp_apply_link,omitempty" `

    /*
        定向计划是否可用 1-可用 0-不可用     */
    SpStatus  *string `json:"sp_status,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalSpCampaign) SetSpCid(v string) *TaobaoTbkDgMaterialOptionalSpCampaign {
    s.SpCid = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalSpCampaign) SetSpName(v string) *TaobaoTbkDgMaterialOptionalSpCampaign {
    s.SpName = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalSpCampaign) SetSpRate(v string) *TaobaoTbkDgMaterialOptionalSpCampaign {
    s.SpRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalSpCampaign) SetSpLockStatus(v string) *TaobaoTbkDgMaterialOptionalSpCampaign {
    s.SpLockStatus = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalSpCampaign) SetSpApplyLink(v string) *TaobaoTbkDgMaterialOptionalSpCampaign {
    s.SpApplyLink = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalSpCampaign) SetSpStatus(v string) *TaobaoTbkDgMaterialOptionalSpCampaign {
    s.SpStatus = &v
    return s
}
