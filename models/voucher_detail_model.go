package models

import (
	"strings"
	"time"
)

// VoucherDetail struct for voucher details
type VoucherDetail struct {
	ID            int64     `orm:"pk;auto;column(id)" json:"id"`
	Voucher       *Voucher  `orm:"column(voucher_id);rel(fk)" json:"voucherId"`
	Code          string    `orm:"column(code)" json:"code"`
	Status        string    `orm:"column(status)" json:"status"`
	UserPubID     string    `orm:"column(user_pubid)" json:"userPubId"`
	GivenTs       time.Time `orm:"column(given_ts)" json:"givenTs"`
	RedeemTs      time.Time `orm:"column(redeem_ts)" json:"redeemTs"`
	TransactionId int64     `orm:"column(transaction_id)" json:"transactionId"`
}

// TableName for users
func (vd *VoucherDetail) TableName() string {
	return "vouchers_detail"
}

// NewCreateVoucherDtl function for create voucher detail
func NewCreateVoucherDtl(code string, v Voucher) VoucherDetail {
	return VoucherDetail{
		Code:    strings.ToUpper(code),
		Status:  "AVAILABLE",
		Voucher: &v,
	}
}
