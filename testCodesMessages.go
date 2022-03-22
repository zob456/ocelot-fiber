package ocelotFiber

import "errors"

const (
	BadRequestCode = 400
	NotAuthorizedCode = 401
	NotFoundCode = 404
	InternalServerErrCode = 500
)
var (
	BadRequestErr = errors.New(PublicBadRequest)
	NotAuthorizedErr = errors.New(PublicNotAuthorized)
	NotFoundErr = errors.New(PublicNotFound)
	TestSqlErr = errors.New("sql: no rows in result set")
	InternalServerErr = errors.New(PublicInternalServerError)
)
