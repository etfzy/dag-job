package job

import "github.com/etfzy/dag-job/task"

func NewTask(name string, operate task.Operate) *task.Task {
	return &task.Task{
		Name:    name,
		Operate: operate,
	}
}
