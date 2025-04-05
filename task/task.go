package task

import (
	"context"
	"errors"
)

type Operate func(context.Context) error
type Task struct {
	Name    string
	Operate Operate
}

func NewTask(name string, operate Operate) *Task {
	return &Task{
		Name:    name,
		Operate: operate,
	}
}

func (t *Task) Validate() error {
	if t.Operate == nil {
		return errors.New("operate is nil")
	}

	if t.Name == "" {
		return errors.New("name is nil")
	}
	return nil
}
