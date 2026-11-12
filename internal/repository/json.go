package repository

import (
	"context"

	"github.com/vadimptrrrr/dofordoing/internal/domain"
)

type Json struct {
	users map[string]*domain.User
}

func (j *Json) SaveUser(ctx context.Context, u *domain.User) error
func (j *Json) All(ctx context.Context) ([]*domain.User, error)
func (j *Json) ByName(ctx context.Context, name string) (*domain.User, error)
func (j *Json) DeleteUser(ctx context.Context, id string) error

func (j *Json) SaveTask(ctx context.Context, userID string, t domain.Task) error
func (j *Json) ByUser(ctx context.Context, userID string) ([]domain.Task, error)
func (j *Json) DeleteTask(ctx context.Context, userID, taskID string) error
