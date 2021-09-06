package appConst

const (

	//redis
	CouponInfoKey = "new_live_coupon_info"
	RecordInfoKey = "coupon_record_info"

	DefaultTimeFormat = "2006-01-02 15:04:05"
)

var ApiMap = map[string][]string{
	"/v1/coupon/market/Discount/create":         []string{"创建优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/Discount/update":         []string{"修改优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/Discount/delete":         []string{"删除优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/Discount/saveInclude":    []string{"配置优惠券", "活动管理-优惠券"},
	"/v1/coupon/market/DiscountDetail/writeOff": []string{"核销优惠券", "活动管理-优惠券"},
}
