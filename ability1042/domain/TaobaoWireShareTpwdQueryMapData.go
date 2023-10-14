package domain

type TaobaoWirelessShareTpwdQueryData struct {
	Content     *string `json:"content,omitempty" `
	Title       *string `json:"title,omitempty" `
	Price       *string `json:"price,omitempty" `
	PicUrl      *string `json:"pic_url,omitempty" `
	Suc         *bool   `json:"suc,omitempty" `
	URL         *string `json:"url,omitempty" `
	NativeUrl   *string `json:"native_url,omitempty" `
	ThumbPicUrl *string `json:"thumb_pic_url,omitempty" `
}

func (s *TaobaoWirelessShareTpwdQueryData) SetContent(v string) *TaobaoWirelessShareTpwdQueryData {
	s.Content = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetTitle(v string) *TaobaoWirelessShareTpwdQueryData {
	s.Title = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetPrice(v string) *TaobaoWirelessShareTpwdQueryData {
	s.Price = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetPicUrl(v string) *TaobaoWirelessShareTpwdQueryData {
	s.PicUrl = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetSuc(v bool) *TaobaoWirelessShareTpwdQueryData {
	s.Suc = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetURL(v string) *TaobaoWirelessShareTpwdQueryData {
	s.URL = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetNativeUrl(v string) *TaobaoWirelessShareTpwdQueryData {
	s.NativeUrl = &v
	return s
}

func (s *TaobaoWirelessShareTpwdQueryData) SetThumbPicUrl(v string) *TaobaoWirelessShareTpwdQueryData {
	s.ThumbPicUrl = &v
	return s
}
