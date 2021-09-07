package request

type WxParam struct {
	Content WxContent `json:"content"`
	State   int       `json:"state"`
}

type WxContent struct {
	Guide int    `json:"guide"`
	Pic   string `json:"pic"`
}

type Banner struct {
	Pic      string `json:"pic"`
	Src      string `json:"src"`
	Type     string `json:"type"`
	LiveId   int    `json:"liveId"`
	LiveName string `json:"liveName"`
}

type Navigation struct {
	Id    int    `json:"id"`
	State int    `json:"state"`
	Pic   string `json:"pic"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Msg   string `json:"msg"`
}

type WxConfig struct {
	Id      int `form:"id"`
	State   int `form:"state"`
	Content struct {
		Guide int    `form:"guide"`
		Pic   string `form:"pic"`
	} `form:"content"`
}

type ShareConfig struct {
	Id       int    `form:"id" binding:"required"`
	Title    string `form:"title" binding:"required"`
	SubTitle string `form:"subTitle" binding:"required"`
	Thumb    string `form:"thumb" binding:"required"`
}

type ViewCountConfig struct {
	ID                     int `form:"id" binding:"required"`
	IsDisplayWatchNum      int `form:"isDisplayWatchNum"`
	CustomWatchTrendNumMin int `form:"customWatchTrendNumMin"`
	CustomWatchTrendNumMax int `form:"customWatchTrendNumMax"`
	CustomWatchLimitMax    int `form:"customWatchLimitMax"`
	CustomWatchNum         int `form:"customWatchNum"`
}

type NavigationConfig struct {
	ID              int    `form:"id" binding:"required"`
	Navigationstate int    `form:"navigationState"`
	Navigationstyle int    `form:"navigationStyle"`
	Content         string `form:"content"`
}

type NavInfo struct {
	ID    int         `json:"id"`
	State int         `json:"state"`
	Pic   string      `json:"pic"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Msg   interface{} `json:"msg"`
	Url   string      `json:"url"`
}

type MenuInfo struct {
	ID      int         `json:"id"`
	MenuId  int         `json:"menuId"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Sort    int         `json:"sort"`
	Content interface{} `json:"content"`
}

type MenuConfig struct {
	ID         int         `json:"id"`
	Aid        int         `json:"aid"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Multilevel int         `json:"multilevel"`
	Status     int         `json:"status"`
	Sort       int         `json:"sort"`
	Content    interface{} `json:"content"`
}

type SourceInfo struct {
	SourceId   int    `json:"sourceId"`
	IsSelect   int    `json:"isSelect"`
	Type       string `json:"type"`
	Style      int    `json:"style"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	CreateTime string `json:"createTime"`
	ID         int    `json:"id"`
}
