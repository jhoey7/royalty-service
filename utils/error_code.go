package utils

// ErrReqInvalid list error
var (
	ErrReqInvalid               = 400
	ErrNone                     = 200
	ErrUserNotFound             = 404
	ErrUserAlreadyExist         = 406
	ErrDefault                  = 508
	ErrReferenceNotFound        = 407
	ErrVoucherNotFound          = 408
	ErrReferenceNotEligible     = 409
	ErrVoucherNotEligibleRedeem = 410

	MsgUserAlreadyExist         = "The mdn you specified is already in use"
	MsgUserNotFound             = "User Not Found"
	MsgReferenceNotFound        = "Reference ID not found"
	MsgErrDefault               = "Don't worry, we are handling this"
	MsgVoucherNotFound          = "Voucher Code not found."
	MsgReferenceNotEligible     = "Reference is not eligible to get voucher"
	MsgVoucherNotEligibleRedeem = "Voucher not eligible to redeem"
)
