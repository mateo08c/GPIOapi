package api

import "encoding/json"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{code, message}
}

func (e *Error) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}
