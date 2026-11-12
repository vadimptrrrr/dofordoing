package domain

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tsks []Task `json:"tasks"`
}

func NewUser(name string) *User

func (u *User) AddTask(t Task) error
func (u *User) TaskByID(id string) (*Task, error)
func (u *User) HasTaskNamed(name string) bool
