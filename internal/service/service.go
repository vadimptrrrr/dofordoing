package service

import (
	"context"

	"github.com/vadimptrrrr/dofordoing/internal/domain"
)

type UserRepo interface {
	SaveUser(ctx context.Context, u *domain.User) error
	All(ctx context.Context) ([]*domain.User, error)
	ByName(ctx context.Context, name string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type TaskRepo interface {
	SaveTask(ctx context.Context, userID string, t domain.Task) error
	ByUser(ctx context.Context, userID string) ([]domain.Task, error)
	DeleteTask(ctx context.Context, userID, taskID string) error
}
