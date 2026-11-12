package service

import (
	"context"
	"time"

	"github.com/vadimptrrrr/dofordoing/internal/domain"
)

type TaskService struct{}

func NewTaskService(repo TaskRepo, users UserRepo) *TaskService

func (s *TaskService) AddTask(ctx context.Context, userID, title string, deadline *time.Time) (domain.Task, error)
func (s *TaskService) Tasks(ctx context.Context, userID string) ([]domain.Task, error)
func (s *TaskService) CompleteTask(ctx context.Context, userID, taskID string) error
func (s *TaskService) WeekTasks(ctx context.Context, userID string, now time.Time) ([]domain.Task, error)
func (s *TaskService) DeleteTask(ctx context.Context, userID, taskID string) error
