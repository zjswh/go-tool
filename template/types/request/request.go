package request

type LoginStruct struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type RegisterStruct struct {
	Phone string `json:"phone"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
}

type GetInfoStruct struct {
	Id int `json:"id"`
	Phone string `json:"phone"`
	Username string `json:"username"`
}

type PageInfo struct {
	Page int `json:"page"`
	Num int `json:"num"`
}

type ProgrammeStruct struct {
	Id int `json:"id"`
	Name string `json:"name"`
	MssUrl string `json:"mssUrl"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	ChannelId int `json:"channelId"`
	Uin int64 `json:"uin"`
	Type int `json:"type"`
	Week int `json:"week"`
	STaskId int `json:"sTaskId"`
	ETaskId int `json:"eTaskId"`
}



