package errCode

import errs "github.com/cristiancll/go-errors"

const (
	INTERNAL errs.ErrorCode = iota + 1
	NOT_FOUND
	NOT_CHANGED
	ACCESS_DENIED
)
