package ViewModel

type ResultData struct {
	CreateTime int64 `json:"createTime"`
	UserId int `json:"userId"`
	Token string `json:"token"`
	RealName string `json:realName`
}

type Vresult struct {
	Code int `json:"code"`
	Data ResultData `json:"data"`
}

type Verror struct {
	Code int `json:"code"`
	ThrowType string `json:"throwType"`
	Message string `json:"message"`
}