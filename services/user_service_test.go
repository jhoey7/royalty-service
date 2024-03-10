package services

import (
	"errors"
	"royalty-service/models"
	"royalty-service/utils"
	"testing"
)

const expectedResponseCode = "Expected resp code to be %v but it was %v"

type fakeUserProcessor struct {
	respFindByMdn  models.User
	errFindByMdn   error
	respInsertUser models.User
	errInsertUser  error
}

func (f fakeUserProcessor) FindByMdn(mdn string) (models.User, error) {
	return f.respFindByMdn, f.errFindByMdn
}
func (f fakeUserProcessor) Insert(user models.User) (models.User, error) {
	return f.respInsertUser, f.errInsertUser
}

var (
	positiveRegisterUser = fakeUserProcessor{
		errFindByMdn:  nil,
		errInsertUser: nil,
	}

	negativeFindUserByMdn = fakeUserProcessor{
		errFindByMdn: errors.New("DB Error"),
	}

	existFindUserByMdn = fakeUserProcessor{
		errFindByMdn:  nil,
		respFindByMdn: models.User{PubID: "123"},
	}

	negativeInsertUser = fakeUserProcessor{
		errInsertUser: errors.New("DB Error"),
	}

	positiveRequest         = []byte(`{"mdn":"6281987876654", "password":"123", "confirmPassword": "123"}`)
	negativeRequest         = []byte(`I AM JSON`)
	passwordNotMatchRequest = []byte(`{"mdn":"6281987876654", "password":"123", "confirmPassword": "abc"}`)
)

func TestUserService_Register(t *testing.T) {
	cases := []struct {
		testName         string
		request          []byte
		userProcessor    UserProcessor
		expectedResponse int
	}{
		{
			testName:         "Positive: Success Flow",
			expectedResponse: utils.ErrNone,
			userProcessor:    positiveRegisterUser,
			request:          positiveRequest,
		},
		{
			testName:         "Negative Test: Invalid Request",
			request:          negativeRequest,
			expectedResponse: utils.ErrReqInvalid,
		},
		{
			testName:         "Negative Test: Failed Find Use By MDN",
			request:          positiveRequest,
			userProcessor:    negativeFindUserByMdn,
			expectedResponse: utils.ErrDefault,
		},
		{
			testName:         "Negative Test: User Already Exist",
			request:          positiveRequest,
			userProcessor:    existFindUserByMdn,
			expectedResponse: utils.ErrUserAlreadyExist,
		},
		{
			testName:         "Negative Test: Password Not Match",
			request:          passwordNotMatchRequest,
			userProcessor:    positiveRegisterUser,
			expectedResponse: utils.ErrPasswordNotMatch,
		},
		{
			testName:         "Negative Test: Failed Register User",
			request:          positiveRequest,
			userProcessor:    negativeInsertUser,
			expectedResponse: utils.ErrDefault,
		},
	}

	for _, c := range cases {
		svc := NewUserService(c.userProcessor, 123)
		resp := svc.Register(c.request)
		if resp.Code != c.expectedResponse {
			t.Errorf(expectedResponseCode, c.expectedResponse, resp.Code)
		}
	}
}
