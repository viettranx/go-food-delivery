package common

import "errors"

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

type ErrorRes struct {
	RootErr error  `json:"root_err"`
	Message string `json:"message"`
	Log     string `json:"log"`
	Key     string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *ErrorRes {
	return &ErrorRes{
		RootErr: root,
		Message: msg,
		Log:     log,
		Key:     key,
	}
}

func NewCustomError(root error, msg string, key string) *ErrorRes {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *ErrorRes) RootError() error {
	if err, ok := e.RootErr.(*ErrorRes); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *ErrorRes) Error() string {
	return e.RootError().Error()
}
