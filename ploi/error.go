package ploi

type ErrorCode int

const (
	ErrorCodeValidationFailed int = 422
	ErrorCodeTooManyAttepts   int = 429
	ErrorCodeInternalError    int = 500
	ErrirCodeMaintaince       int = 503
)
