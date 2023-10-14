package request

type TaobaoWirelessShareTpwdQueryRequest struct {
	PasswordContent *string `json:"password_content" required:"true"`
}

func (s *TaobaoWirelessShareTpwdQueryRequest) SetPasswordContent(v string) *TaobaoWirelessShareTpwdQueryRequest {
	s.PasswordContent = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if s.PasswordContent != nil {
		paramMap["password_content"] = *s.PasswordContent
	}
	return paramMap
}

func (s *TaobaoWirelessShareTpwdQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
