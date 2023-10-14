package domain


type TaobaoTbkDgOptimusMaterialSpCampaign struct {
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

func (s *TaobaoTbkDgOptimusMaterialSpCampaign) SetSpCid(v string) *TaobaoTbkDgOptimusMaterialSpCampaign {
    s.SpCid = &v
    return s
}
func (s *TaobaoTbkDgOptimusMaterialSpCampaign) SetSpName(v string) *TaobaoTbkDgOptimusMaterialSpCampaign {
    s.SpName = &v
    return s
}
func (s *TaobaoTbkDgOptimusMaterialSpCampaign) SetSpRate(v string) *TaobaoTbkDgOptimusMaterialSpCampaign {
    s.SpRate = &v
    return s
}
func (s *TaobaoTbkDgOptimusMaterialSpCampaign) SetSpLockStatus(v string) *TaobaoTbkDgOptimusMaterialSpCampaign {
    s.SpLockStatus = &v
    return s
}
func (s *TaobaoTbkDgOptimusMaterialSpCampaign) SetSpApplyLink(v string) *TaobaoTbkDgOptimusMaterialSpCampaign {
    s.SpApplyLink = &v
    return s
}
func (s *TaobaoTbkDgOptimusMaterialSpCampaign) SetSpStatus(v string) *TaobaoTbkDgOptimusMaterialSpCampaign {
    s.SpStatus = &v
    return s
}
