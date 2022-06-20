package account

import "errors"

var (
	ERR_INVALID_ID           = errors.New("err_invalid_id")
	ERR_INVALID_ACCOUNT_TYPE = errors.New("err_invalid_account_type")
	ERR_INVALID_NONCE        = errors.New("err_invalid_nonce")
	ERR_NONCE_TOO_LOW        = errors.New("err_nonce_too_low")
	ERR_INVALID_SIGNATURE    = errors.New("err_invalid_signature")
	ERR_INVALID_TRANSACTION  = errors.New("err_invalid_transaction")
)
