package repository

import (
	"context"
	"sync"

	"github.com/vadimptrrrr/dofordoing/internal/domain"
)

type Memory struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewMemory() *Memory

func (m *Memory) SaveUser(ctx context.Context, u *domain.User) error
func (m *Memory) All(ctx context.Context) ([]*domain.User, error)
func (m *Memory) ByName(ctx context.Context, name string) (*domain.User, error)
func (m *Memory) DeleteUser(ctx context.Context, id string) error

func (m *Memory) SaveTask(ctx context.Context, userID string, t domain.Task) error
func (m *Memory) ByUser(ctx context.Context, userID string) ([]domain.Task, error)
func (m *Memory) DeleteTask(ctx context.Context, userID, taskID string) error
