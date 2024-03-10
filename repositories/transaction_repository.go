package repositories

import (
	"github.com/astaxie/beego/orm"
	"royalty-service/models"
)

// TransactionRepository struct
type TransactionRepository struct {
	db orm.Ormer
}

// NewTransactionRepository is func for initiate TransactionRepository
func NewTransactionRepository(o orm.Ormer) TransactionRepository {
	return TransactionRepository{db: o}
}

// Insert func for insert voucher
func (repo TransactionRepository) Insert(t models.Transaction) (models.Transaction, error) {
	id, err := repo.db.Insert(&t)
	t.ID = id
	return t, err
}

// FindTransactionByPubIDAndTypeAndUserPubID is func to find transaction
func (repo TransactionRepository) FindTransactionByPubIDAndTypeAndUserPubID(pubID, invType, userPubID string) (models.Transaction, error) {
	trx := models.Transaction{}
	err := repo.db.QueryTable("transactions").
		Filter("invoice_pubid", pubID).
		Filter("invoice_type", invType).
		Filter("user_pubid", userPubID).
		One(&trx)

	return trx, err
}

// UpdateColumns function for update transaction data using certain columns
func (repo TransactionRepository) UpdateColumns(t models.Transaction, cols ...string) error {
	_, err := repo.db.Update(&t, cols...)
	return err
}
