package service

import (
	"context"

	"github.com/vadimptrrrr/dofordoing/internal/domain"
)

type UserRepo interface {
	Save(ctx context.Context, u *domain.User) error
	All(ctx context.Context) ([]*domain.User, error)
	ByName(ctx context.Context, name string) (*domain.User, error)
	Delete(ctx context.Context, id string) error
}

type TaskRepo interface {
	Save(ctx context.Context, userID string, t domain.Task) error
	ByUser(ctx context.Context, userID string) ([]domain.Task, error)
	Delete(ctx context.Context, userID, taskID string) error
}
