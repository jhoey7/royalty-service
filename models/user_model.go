package models

import (
	"royalty-service/utils"
	"time"
)

// User struct for user
type User struct {
	PubID       string    `orm:"pk;column(pubid)" json:"pubId"`
	Email       string    `orm:"column(email);null" json:"email"`
	Mdn         string    `orm:"column(mdn);null" json:"mdn"`
	UserName    string    `orm:"column(username);null" json:"userName"`
	FullName    string    `orm:"column(full_name);null" json:"name"`
	Address     string    `orm:"column(address);null" json:"address"`
	DateOfBirth time.Time `orm:"column(date_of_birth);null" json:"dateOfBirth"`
	Gender      string    `orm:"column(gender);null" json:"gender"`
	CreateTs    time.Time `orm:"column(create_ts);null" json:"createTs"`
	CreateBy    string    `orm:"column(create_by);null" json:"createBy"`
	UpdateTs    time.Time `orm:"column(update_ts);null" json:"updateTs"`
	UpdateBy    string    `orm:"column(update_by);null" json:"updateBy"`
}

// TableName for users
func (u *User) TableName() string {

	return "users"
}

// NewUser is func for initialize User
func NewUser(req RegisterUserRequest) User {
	pubID := utils.NewV4().String()
	dob, _ := time.Parse("2006-01-02", req.DateOfBirth)
	return User{
		PubID:       pubID,
		Email:       req.Email,
		Mdn:         req.Mdn,
		UserName:    req.UserName,
		FullName:    req.FullName,
		Address:     req.Address,
		DateOfBirth: dob,
		Gender:      req.Gender,
		CreateTs:    time.Now(),
		CreateBy:    "SYSTEM",
		UpdateTs:    time.Now(),
	}
}

// RegisterUserRequest struct for register new user
type RegisterUserRequest struct {
	Email       string `json:"email"`
	Mdn         string `json:"mdn" valid:"Required"`
	UserName    string `json:"userName"`
	FullName    string `json:"name"`
	Address     string `json:"address"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
}
