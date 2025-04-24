package enums

type promotionStatusEnum struct {
	Active   uint8
	Inactive uint8
}

var PromotionStatus = promotionStatusEnum{
	Active:   1,
	Inactive: 0,
}
