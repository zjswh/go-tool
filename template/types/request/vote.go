package request

type CreateVoteStruct struct {
	Topic string `json:"topic"`
	VoteIntro string `json:"vote_intro"`
	VoteType int `json:"vote_type"`
	Banner string `json:"banner"`
	VoteWay int `json:"vote_way"`
	RefreshStatus int `json:"refresh_status"`
	VoteOption int `json:"vote_option"`
	VoteChooseNum int `json:"vote_choose_num"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
	IsRank int `json:"is_rank"`
	Content []VoteContentStruct `json:"content"`
}

type VoteInfoStruct struct {
	ID int  `form:"id"`
	Topic string  `form:"topic"`
	Banner string  `form:"banner"`
	AccessId int `form:"access_id"`
	Aid int `form:"aid"`
	Url string  `form:"url"`
	StartTime string  `form:"start_time"`
	EndTime string  `form:"end_time"`
	IsRank int  `form:"is_rank"`
	VoteChooseNum int  `form:"vote_choose_num"`
	VoteIntro string  `form:"vote_intro"`
	VoteOption int  `form:"vote_option"`
	VoteType int  `form:"vote_type"`
	VoteWay int  `form:"vote_way"`
	RefreshStatus int  `form:"refresh_status"`
	Content string `form:"content"`
}

type VoteContentStruct struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Intro string `json:"intro"`
	VoteId int `json:"vote_id"`
	Pic string `json:"pic"`
	VideoUrl string `json:"videoUrl"`
	VideoCoverImg string `json:"videoCoverImg"`
}

type VoteOption struct {
	OptionId int `form:"optionId"`
	UserId int `form:"userId"`
	Type string `form:"type"`
	Num int `form:"num"`
	Page int `form:"page"`
	Phone string `form:"phone"`
	Name string `form:"name"`
	StartTime string `form:"startTime"`
	EndTime string `form:"endTime"`
	Filename string `form:"filename"`
}
