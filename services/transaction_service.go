package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"royalty-service/models"
	"royalty-service/utils"
	"time"
)

// TransactionProcessor interface for transaction
type TransactionProcessor interface {
	Insert(v models.Transaction) (models.Transaction, error)
	FindTransactionByPubIDAndTypeAndUserPubID(pubID, invType, userPubID string) (models.Transaction, error)
	UpdateColumns(t models.Transaction, cols ...string) error
}

// UserFinder interface for find user
type UserFinder interface {
	FindByPubID(pubID string) (models.User, error)
}

// VoucherDetailProcessor interface
type VoucherDetailProcessor interface {
	FindGivenByCodeAndUserPubID(code, status string) (models.VoucherDetail, error)
	UpdateColumns(vd models.VoucherDetail, cols ...string) error
	FindAvailableVoucher() (models.VoucherDetail, error)
}

// VoucherProcessor interface
type VoucherProcessor interface {
	FindVoucherByID(id int64) (models.Voucher, error)
}

// TransactionService struct
type TransactionService struct {
	Identifier             int64
	trxProcessor           TransactionProcessor
	voucherDetailProcessor VoucherDetailProcessor
	voucherProcessor       VoucherProcessor
	userFinder             UserFinder
	o                      orm.Ormer
}

// NewTransactionService is func for initialize TransactionService
func NewTransactionService(tp TransactionProcessor, uf UserFinder, vdp VoucherDetailProcessor,
	vp VoucherProcessor, o orm.Ormer, i int64) TransactionService {
	return TransactionService{
		trxProcessor:           tp,
		userFinder:             uf,
		voucherDetailProcessor: vdp,
		voucherProcessor:       vp,
		o:                      o,
		Identifier:             i,
	}
}

// Create is func for create voucher
func (svc TransactionService) Create(b []byte) models.Response {
	request := models.CreateTrxReq{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("transaction request: %+v", request)

	var user models.User
	var err error
	if request.UserPubID != "" {
		user, err = svc.userFinder.FindByPubID(request.UserPubID)
		if err != nil {
			logs.Warn("[%d] Failed to find user: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgUserNotFound, utils.ErrUserNotFound)
		}
	}

	svc.o.Begin()
	trx, err := svc.trxProcessor.Insert(request.ToInsertReq())
	if err != nil {
		svc.o.Rollback()
		logs.Warn("[%d] Failed to insert voucher: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	if request.VoucherCode != "" {
		vd, err := svc.voucherDetailProcessor.FindGivenByCodeAndUserPubID(request.VoucherCode, user.PubID)
		if err != nil {
			logs.Warn("[%d] Failed to find Voucher Detail: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgVoucherNotFound, utils.ErrVoucherNotFound)
		}

		voucher, err := svc.voucherProcessor.FindVoucherByID(vd.Voucher.ID)
		if err != nil {
			logs.Warn("[%d] Failed to find Voucher: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgVoucherNotFound, utils.ErrVoucherNotFound)
		}

		now := time.Now()
		if now.After(voucher.StartTs) && now.Before(voucher.EndTs) {
			trx.ChargeAmount = trx.Amount - voucher.VoucherAmount
			trx.UpdateTs = time.Now()
			err = svc.trxProcessor.UpdateColumns(trx, "charge_amount", "update_ts")
			if err != nil {
				svc.o.Rollback()
				logs.Warn("[%d] Failed to update transaction voucher: %s", svc.Identifier, err.Error())
				return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
			}

			vd.Status = "REDEEMED"
			vd.RedeemTs = time.Now()
			err = svc.voucherDetailProcessor.UpdateColumns(vd, "status", "redeem_ts")
			if err != nil {
				svc.o.Rollback()
				logs.Warn("[%d] Failed to update voucher detail: %s", svc.Identifier, err.Error())
				return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
			}
		} else {
			logs.Warn("[%d] Voucher not eligible to redeem", svc.Identifier)
			return models.ResponseError(utils.MsgVoucherNotEligibleRedeem, utils.ErrVoucherNotEligibleRedeem)
		}
	}

	if request.InvoiceType == "SHOP" {
		if request.Amount >= beego.AppConfig.DefaultFloat("minTransactionAmount", 1000000) && request.ReferenceID != "" {
			trxTenant, err := svc.trxProcessor.FindTransactionByPubIDAndTypeAndUserPubID(request.ReferenceID, "TENANT", user.PubID)
			if err != nil {
				svc.o.Rollback()
				logs.Warn("[%d] Failed to reference invoice: %s", svc.Identifier, err.Error())
				return models.ResponseError(utils.MsgReferenceNotFound, utils.ErrReferenceNotFound)
			}

			if !trxTenant.IsEligibleVoucher {
				svc.o.Rollback()
				logs.Warn("[%d] Reference is not eligible to get voucher", svc.Identifier)
				return models.ResponseError(utils.MsgReferenceNotEligible, utils.ErrReferenceNotEligible)
			}

			totalVoucher := math.Floor(request.Amount / beego.AppConfig.DefaultFloat("minTransactionAmount", 1000000))
			for i := 0; i < int(totalVoucher); i++ {
				voucher, err := svc.voucherDetailProcessor.FindAvailableVoucher()
				if err != nil {
					logs.Warn("[%d] Failed to find Voucher: %s", svc.Identifier, err.Error())
				}

				if voucher.Code != "" {
					voucher.UserPubID = user.PubID
					voucher.Status = "GIVEN"
					voucher.GivenTs = time.Now()
					voucher.TransactionId = trx.ID
					err = svc.voucherDetailProcessor.UpdateColumns(voucher, "user_pubid", "status", "given_ts", "transaction_id")
					if err != nil {
						svc.o.Rollback()
						logs.Warn("[%d] Failed to update voucher detail: %s", svc.Identifier, err.Error())
						return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
					}
				}
			}

			trxTenant.IsEligibleVoucher = false
			trxTenant.UpdateTs = time.Now()
			err = svc.trxProcessor.UpdateColumns(trxTenant, "is_eligible_voucher", "update_ts")
			if err != nil {
				svc.o.Rollback()
				logs.Warn("[%d] Failed to update transaction tenant: %s", svc.Identifier, err.Error())
				return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
			}
		}
	}

	svc.o.Commit()
	return models.ResponseSuccess(trx)
}
