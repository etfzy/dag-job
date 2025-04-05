package runtime

import (
	"context"

	"github.com/etfzy/dag-job/graph"
	job_result "github.com/etfzy/dag-job/result"
	"github.com/etfzy/dag-job/runtime/state_machine"
	"github.com/etfzy/dag-job/task"

	"github.com/panjf2000/ants/v2"
)

type Runtime struct {
	graph *graph.Graph
	mTask map[string]*task.Task

	stateMachine *state_machine.StateMachine
}

var defaultAntsPool *ants.Pool

func init() {
	defaultAntsPool, _ = ants.NewPool(ants.DefaultAntsPoolSize)
}

func NewRuntime(g *graph.Graph, mTask map[string]*task.Task) *Runtime {
	r := &Runtime{
		graph:        g,
		mTask:        mTask,
		stateMachine: state_machine.NewStateMachine(g, mTask, defaultAntsPool),
	}

	return r
}

func (r *Runtime) Run(ctx context.Context) map[string]job_result.JobResult {
	return r.stateMachine.Start(ctx)
}
