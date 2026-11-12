package domain

import "time"

type Status string

const (
	StatusTodo    Status = "todo"
	StatusDone    Status = "done"
	StatusSkipped Status = "skipped"
)

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func NewTask(title string, deadline *time.Time) (Task, error)

func (t *Task) Complete(now time.Time) error
func (t Task) IsOverdue(now time.Time) bool
func (t Task) InWeek(from, to time.Time) bool
func (t Task) Validate() error
