package store

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserIdDuplicated    = errors.New("user id duplicated")
	ErrImageNameDuplicated = errors.New("image name duplicated")
)
