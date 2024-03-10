package services

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"royalty-service/models"
	"royalty-service/utils"
)

// UserProcessor interface for user process
type UserProcessor interface {
	FindByMdn(mdn string) (models.User, error)
	Insert(user models.User) (models.User, error)
}

// UserService struct
type UserService struct {
	Identifier    int64
	userProcessor UserProcessor
}

// NewUserService is func for initialize UserService
func NewUserService(up UserProcessor, i int64) UserService {
	return UserService{
		userProcessor: up,
		Identifier:    i,
	}
}

// Register is func for register new user
func (svc UserService) Register(b []byte) models.Response {
	request := models.RegisterUserRequest{}
	res := ConvertRequest(b, &request, svc.Identifier)
	if res.Code != utils.ErrNone {
		logs.Warn("[%d] Failed to convert request: %+v", svc.Identifier, res)
		return models.ResponseError(res.ErrorMessage, utils.ErrReqInvalid)
	}
	logs.Info("registerUser request: %+v", request)

	user, err := svc.userProcessor.FindByMdn(request.Mdn)
	if err != nil && !errors.Is(err, orm.ErrNoRows) {
		logs.Warn("[%d] Failed to find user by mdn: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	if user.PubID != "" {
		logs.Warn("[%d] User already exist: %s", svc.Identifier, user.Mdn)
		return models.ResponseError(utils.MsgUserAlreadyExist, utils.ErrUserAlreadyExist)
	}

	newUserReq := models.NewUser(request)
	user, err = svc.userProcessor.Insert(newUserReq)
	if err != nil {
		logs.Warn("[%d] Failed to register user: %s", svc.Identifier, err.Error())
		return models.ResponseError(utils.MsgErrDefault, utils.ErrDefault)
	}

	return models.ResponseSuccess(user)
}
