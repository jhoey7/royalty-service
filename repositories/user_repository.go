package repositories

import (
	"github.com/astaxie/beego/orm"
	"royalty-service/models"
)

// UserRepository struct
type UserRepository struct {
	db orm.Ormer
}

// NewUserRepository is func for initiate UserRepository
func NewUserRepository(o orm.Ormer) UserRepository {
	return UserRepository{db: o}
}

// FindByMdn is func for find user by mdn
func (repo UserRepository) FindByMdn(mdn string) (models.User, error) {
	user := models.User{}
	err := repo.db.QueryTable("users").
		Filter("mdn", mdn).
		One(&user)

	return user, err
}

// FindByPubID is func for find user by PubID
func (repo UserRepository) FindByPubID(pubID string) (models.User, error) {
	user := models.User{}
	err := repo.db.QueryTable("users").
		Filter("pubid", pubID).
		One(&user)

	return user, err
}

// Insert func for insert user
func (repo UserRepository) Insert(user models.User) (models.User, error) {
	_, err := repo.db.Insert(&user)
	return user, err
}
