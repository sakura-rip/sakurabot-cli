package database

import "gorm.io/gorm"

type ChargeType string

const (
	ChargeType_PAYPAY         ChargeType = "paypay"
	ChargeType_LINEPAY        ChargeType = "linepay"
	ChargeType_LINEPAY_CHARGE ChargeType = "linepay-c"
	ChargeType_PAYPAL         ChargeType = "paypal"
	ChargeType_BTC            ChargeType = "btc"
	ChargeType_LITE_COIN      ChargeType = "ltc"
	ChargeType_KYASH          ChargeType = "kyash"
	ChargeType_MERUPAY        ChargeType = "perupay"
	ChargeType_AMAZON_GIFT    ChargeType = "amazon"
)

type Charge struct {
	*gorm.Model

	Amount int
	Type   ChargeType
	UserId int
}
