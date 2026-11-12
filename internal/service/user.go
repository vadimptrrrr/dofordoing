package service

import (
	"context"

	"github.com/vadimptrrrr/dofordoing/internal/domain"
)

type UserService struct {
	uRepo UserRepo
}

func NewUserService(repo UserRepo) *UserService

func (s *UserService) AddUser(ctx context.Context, name string) (*domain.User, error)
func (s *UserService) Users(ctx context.Context) ([]*domain.User, error)
func (s *UserService) UserByName(ctx context.Context, name string) (*domain.User, error)
