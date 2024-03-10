package models

import (
	"fmt"
	"time"
)

// Transaction struct for transaction
type Transaction struct {
	ID                int64     `orm:"pk;auto;column(id)" json:"id"`
	InvoicePubID      string    `orm:"column(invoice_pubid)" json:"invoicePubId"`
	Amount            float64   `orm:"column(amount)" json:"amount"`
	ChargeAmount      float64   `orm:"column(charge_amount)" json:"chargeAmount"`
	UserPubID         string    `orm:"column(user_pubid)" json:"userPubId"`
	InvoiceType       string    `orm:"column(invoice_type)" json:"invoiceType"`
	CreateTs          time.Time `orm:"column(create_ts)" json:"createTs"`
	UpdateTs          time.Time `orm:"column(update_ts)" json:"updateTs"`
	IsEligibleVoucher bool      `orm:"column(is_eligible_voucher)" json:"isEligibleVoucher"`
}

// TableName for transaction
func (t *Transaction) TableName() string {
	return "transactions"
}

// CreateTrxReq request
type CreateTrxReq struct {
	Amount      float64 `json:"amount"`
	VoucherCode string  `json:"voucherCode"`
	UserPubID   string  `json:"userPubId"`
	InvoiceType string  `json:"invoiceType"`
	ReferenceID string  `json:"referenceId"`
}

// ToInsertReq func
func (ctr CreateTrxReq) ToInsertReq() Transaction {
	isEligible := false
	if ctr.InvoiceType == "TENANT" {
		isEligible = true
	}
	return Transaction{
		InvoicePubID:      fmt.Sprintf("%X", time.Now().UnixNano()),
		Amount:            ctr.Amount,
		ChargeAmount:      ctr.Amount,
		UserPubID:         ctr.UserPubID,
		CreateTs:          time.Now(),
		InvoiceType:       ctr.InvoiceType,
		IsEligibleVoucher: isEligible,
	}
}
