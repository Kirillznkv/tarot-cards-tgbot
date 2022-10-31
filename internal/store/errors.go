package store

import "errors"

var (
	ErrUserIdDuplicated    = errors.New("user id duplicated")
	ErrImageNameDuplicated = errors.New("image name duplicated")
)
