package common

import (
	"fmt"
	"strings"
)

func ErrDB(err error) *ErrorRes {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *ErrorRes {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrCannotListEntity(entity string, err error) *ErrorRes {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
