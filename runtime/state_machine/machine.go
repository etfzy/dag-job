package state_machine

import (
	"context"

	graph_context "github.com/etfzy/dag-job/context"
	"github.com/etfzy/dag-job/graph"
	job_result "github.com/etfzy/dag-job/result"
	"github.com/etfzy/dag-job/runtime/singal"
	"github.com/etfzy/dag-job/task"

	"github.com/panjf2000/ants/v2"
)

type StateMachine struct {
	stateNodes map[string]*StateNode
	mTask      map[string]*task.Task
	graph      *graph.Graph
	antsPool   *ants.Pool
}

func NewStateMachine(g *graph.Graph, mTask map[string]*task.Task, pool *ants.Pool) *StateMachine {
	sm := &StateMachine{
		stateNodes: map[string]*StateNode{},
		mTask:      mTask,
		graph:      g,
		antsPool:   pool,
	}

	for name, node := range g.GetNodes() {
		sm.stateNodes[node.GetName()] = NewStateNode(name, mTask[name], node, sm)
	}

	return sm
}

func (e *StateMachine) Start(ctx context.Context) map[string]job_result.JobResult {
	grpahCtx := graph_context.NewGraphContext(ctx)

	startNodes := e.graph.IndependentNodes()

	singal := singal.NewSingal(len(e.mTask))

	for name, _ := range startNodes {
		rn := e.stateNodes[name]
		singal.Add()
		err := e.antsPool.Submit(func() {
			rn.run(grpahCtx, singal)
		})

		if err != nil {
			panic(err)
		}
	}

	singal.Wait()

	results := make(map[string]job_result.JobResult)
	for name, v := range e.stateNodes {
		jr := &job_result.JobResult{
			Name:  name,
			State: v.State.GetState(),
			Start: v.State.GetStart().Format("2006-01-02 15:04:05.000"),
			End:   v.State.GetEnd().Format("2006-01-02 15:04:05.000"),
		}

		if v.State.Error() != nil {
			jr.Error = v.State.Error().Error()
		}

		results[name] = *jr
	}

	return results
}
