package repositories

import (
	"github.com/astaxie/beego/orm"
	"royalty-service/models"
)

// VoucherRepository struct
type VoucherRepository struct {
	db orm.Ormer
}

// NewVoucherRepository is func for initiate VoucherRepository
func NewVoucherRepository(o orm.Ormer) VoucherRepository {
	return VoucherRepository{db: o}
}

// Insert func for insert voucher
func (repo VoucherRepository) Insert(v models.Voucher) (models.Voucher, error) {
	_, err := repo.db.Insert(&v)
	return v, err
}

// FindVoucherByID function for find voucher  by id
func (repo VoucherRepository) FindVoucherByID(id int64) (models.Voucher, error) {
	v := models.Voucher{}
	err := repo.db.QueryTable("vouchers").
		Filter("id", id).
		One(&v)

	return v, err
}
