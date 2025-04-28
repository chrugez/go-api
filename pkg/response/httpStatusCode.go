package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParamInvalid  = 20003 // Email is invalid
	ErrCodeInvalidToken  = 30001 // Token is invalid
	ErrCodeUserHasExists = 50001 // User has already existed
)

// message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrCodeInvalidToken:  "Token is invalid",
	ErrCodeUserHasExists: "User has already existed",
}