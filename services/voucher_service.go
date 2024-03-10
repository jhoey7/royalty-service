package services

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"royalty-service/models"
	"royalty-service/utils"
)

// VoucherCreator interface for crate voucher
type VoucherCreator interface {
	Insert(v models.Voucher) (models.Voucher, error)
}

// VoucherDetailCreator interface for create voucher details
type VoucherDetailCreator interface {
	Insert(vd models.VoucherDetail) (models.VoucherDetail, error)
}

// VoucherService struct
type VoucherService struct {
	Identifier        int64
	voucherCreator    VoucherCreator
	voucherDtlCreator VoucherDetailCreator
	o                 orm.Ormer
}

// NewVoucherService is func for initialize VoucherService
func NewVoucherService(vc VoucherCreator, vdc VoucherDetailCreator, o orm.Ormer, i int64) VoucherService {
	return VoucherService{
		voucherCreator:    vc,
		voucherDtlCreator: vdc,
		o:                 o,
		Identifier:        i,
	}
}

// Create is func for create voucher
func (svc VoucherService) Create(b []byte) models.Response {
	request := models.CreateVoucherReq{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("voucherCreate request: %+v", request)

	svc.o.Begin()
	voucher, err := svc.voucherCreator.Insert(request.ToInsertReq())
	if err != nil {
		svc.o.Rollback()
		logs.Warn("[%d] Failed to insert voucher: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	vDtls := request.TotalBudget / request.VoucherAmount
	var vds []*models.VoucherDetail
	for i := 0; i < int(vDtls); i++ {
		vCode := utils.GenerateRandomString(7)
		dtlReq := models.NewCreateVoucherDtl(vCode, voucher)
		vd, err := svc.voucherDtlCreator.Insert(dtlReq)
		if err != nil {
			svc.o.Rollback()
			logs.Warn("[%d] Failed to insert voucher detail: %s", svc.Identifier, err.Error())
			return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
		}

		vds = append(vds, &vd)
	}

	svc.o.Commit()

	voucher.VoucherDetail = vds
	return models.ResponseSuccess(voucher)
}
