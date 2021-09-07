package request

type CreateCouponStruct struct {
	Name          string  `form:"name" binding:"required"`
	WaiverAmount  float32 `form:"waiverAmount" binding:"required"`
	Condition     float32     `form:"condition"`
	ConditionType int     `form:"conditionType"`
	Type          int     `form:"type" binding:"required"`
	CouponType    int     `form:"couponType"`
	Num           int     `form:"num" binding:"required"`
	Intro         string  `form:"intro"`
	StartTime     int64   `form:"startTime" binding:"required"`
	EndTime       int64   `form:"endTime" binding:"required"`
}
