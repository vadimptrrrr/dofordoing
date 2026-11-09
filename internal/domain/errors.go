package domain

import "errors"

var (
	ErrEmptyTitle   = errors.New("[TASK] Not found title from task")
	ErrTaskNotFound = errors.New("[TASK]")
	ErrUserNotFound = errors.New("[USER]")
	ErrUserExists   = errors.New("[USER]")
)
