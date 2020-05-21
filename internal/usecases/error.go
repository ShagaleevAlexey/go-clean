package usecases

import "encoding/json"

type AppError struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

func NewAppError(code int, message string) *AppError {
	return &AppError{code, message}
}

func (e *AppError) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

func (e *AppError) Errorb() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		return []byte(err.Error())
	}

	return data
}

var (
	ErrUserNotFound         = NewAppError(20001, "User not found")
	ErrUserUnauthorizeError = NewAppError(20002, "User unauthorized")
	TokenIsInvalidError     = NewAppError(20003, "Token is invalid")
	TokenGenerationError    = NewAppError(20004, "Token generation error")
)
