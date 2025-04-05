package job

import (
	"context"

	"github.com/etfzy/dag-job/graph"
	job_result "github.com/etfzy/dag-job/result"
	"github.com/etfzy/dag-job/runtime"
	"github.com/etfzy/dag-job/task"
)

type DAGJob struct {
	graph   *graph.Graph
	mTask   map[string]*task.Task
	isBuild bool
}

func NewDAGJob() *DAGJob {
	return &DAGJob{
		graph: graph.NewGraph("job"),
		mTask: map[string]*task.Task{},
	}
}

func (j *DAGJob) AddTask(t *task.Task) error {

	err := t.Validate()
	if err != nil {
		return err
	}

	err = j.graph.AddNode(t.Name)

	if err != nil {
		return err
	}

	j.mTask[t.Name] = t
	return nil
}

func (j *DAGJob) AddEdge(from, to string) error {
	return j.graph.AddEdge(from, to)
}

// 验证DAG
func (j *DAGJob) Build() error {
	return graph.HasCycle(j.graph.GetNodes())
}

func (j *DAGJob) Run(ctx context.Context) map[string]job_result.JobResult {
	r := runtime.NewRuntime(j.graph, j.mTask)
	return r.Run(ctx)
}
