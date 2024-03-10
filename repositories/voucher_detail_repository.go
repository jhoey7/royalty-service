package repositories

import (
	"github.com/astaxie/beego/orm"
	"royalty-service/models"
)

// VoucherDetailRepository struct
type VoucherDetailRepository struct {
	db orm.Ormer
}

// NewVoucherDetailRepository is func for initiate VoucherDetailRepository
func NewVoucherDetailRepository(o orm.Ormer) VoucherDetailRepository {
	return VoucherDetailRepository{db: o}
}

// Insert func for insert voucher
func (repo VoucherDetailRepository) Insert(vd models.VoucherDetail) (models.VoucherDetail, error) {
	_, err := repo.db.Insert(&vd)
	return vd, err
}

// FindGivenByCodeAndUserPubID function for find given voucher by status and userPubID
func (repo VoucherDetailRepository) FindGivenByCodeAndUserPubID(code, userPubID string) (models.VoucherDetail, error) {
	vd := models.VoucherDetail{}
	err := repo.db.QueryTable("vouchers_detail").
		Filter("code", code).
		Filter("status", "GIVEN").
		Filter("user_pubid", userPubID).
		One(&vd)

	return vd, err
}

// FindAvailableVoucher function for find available voucher
func (repo VoucherDetailRepository) FindAvailableVoucher() (models.VoucherDetail, error) {
	vd := models.VoucherDetail{}
	err := repo.db.QueryTable("vouchers_detail").
		Filter("status", "AVAILABLE").
		One(&vd)

	return vd, err
}

// UpdateColumns function for update voucher detail data using certain columns
func (repo VoucherDetailRepository) UpdateColumns(vd models.VoucherDetail, cols ...string) error {
	_, err := repo.db.Update(&vd, cols...)
	return err
}
