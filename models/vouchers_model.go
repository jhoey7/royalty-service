package models

import (
	"time"
)

// Voucher struct for vouchers
type Voucher struct {
	ID            int64            `orm:"pk;auto;column(id)" json:"id"`
	StartTs       time.Time        `orm:"column(start_ts)" json:"startTs"`
	EndTs         time.Time        `orm:"column(end_ts)" json:"endTs"`
	TotalBudget   float64          `orm:"column(total_budget)" json:"totalBudget"`
	VoucherAmount float64          `orm:"column(voucher_amount)" json:"voucherAmount"`
	CreateTs      time.Time        `orm:"column(create_ts)" json:"createTs"`
	UpdateTs      time.Time        `orm:"column(update_ts)" json:"updateTs"`
	VoucherDetail []*VoucherDetail `orm:"reverse(many);null"`
}

// TableName for users
func (v *Voucher) TableName() string {
	return "vouchers"
}

// CreateVoucherReq request
type CreateVoucherReq struct {
	StartTs       string  `json:"startTs"`
	EndTs         string  `json:"endTs"`
	TotalBudget   float64 `json:"totalBudget"`
	VoucherAmount float64 `json:"voucherAmount"`
}

// ToInsertReq func
func (cvr CreateVoucherReq) ToInsertReq() Voucher {
	layout := "2006-01-02 15:04:05"
	startTs, _ := time.Parse(layout, cvr.StartTs)
	endTs, _ := time.Parse(layout, cvr.EndTs)
	return Voucher{
		StartTs:       startTs,
		EndTs:         endTs,
		TotalBudget:   cvr.TotalBudget,
		VoucherAmount: cvr.VoucherAmount,
		CreateTs:      time.Now(),
	}
}
