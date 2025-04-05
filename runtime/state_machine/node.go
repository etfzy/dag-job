package state_machine

import (
	graph_context "github.com/etfzy/dag-job/context"
	"github.com/etfzy/dag-job/graph"
	"github.com/etfzy/dag-job/runtime/singal"

	"sync"

	"github.com/etfzy/dag-job/task"
)

type StateNode struct {
	trickLock sync.Mutex
	Name      string
	State     *state
	Task      *task.Task
	Node      *graph.Node
	Super     *StateMachine
}

func NewStateNode(name string, task *task.Task, node *graph.Node, super *StateMachine) *StateNode {
	return &StateNode{
		trickLock: sync.Mutex{},
		Name:      name,
		State:     newState(),
		Task:      task,
		Node:      node,
		Super:     super,
	}
}

func (r *StateNode) tricker() bool {
	r.trickLock.Lock()
	defer r.trickLock.Unlock()

	//为了保证可以重复执行，因此，只检查失败状态
	if r.State.IsFailed() || r.State.IsRunning() {
		return false
	}

	preNodes := r.Node.GetPreNodes()

	for name, _ := range preNodes {
		node := r.Super.stateNodes[name]
		if !node.State.IsSuccess() {
			return false
		}
	}

	//触发成功，设置为running状态
	r.State.Running()

	return true
}

func (s *StateNode) run(ctx *graph_context.GraphContext, sg *singal.Singal) {
	defer func() {
		sg.Done()
	}()

	if ctx.Context.Err() != nil {
		return
	}

	//触发不成功，就直接返回;
	//严格模式，有一个错误就全取消
	if !s.tricker() {
		return
	}

	err := s.Task.Operate(ctx.Context)

	//向上广播error，终结未执行节点
	if err != nil {
		ctx.Cancel(err)
		s.State.Failed(err)
		return
	}

	//状态Success
	s.State.Success()

	nextNodes := s.Node.GetNextNodes()

	for name, _ := range nextNodes {

		sg.Add()
		node := s.Super.stateNodes[name]
		err := node.Super.antsPool.Submit(func() {
			node.run(ctx, sg)
		})

		if err != nil {
			panic(err)
		}
	}

}
