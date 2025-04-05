package graph_context

import "context"

type GraphContext struct {
	Context context.Context
	cancel  context.CancelCauseFunc
}

func NewGraphContext(pctx context.Context) *GraphContext {
	ctx, cancel := context.WithCancelCause(pctx)
	return &GraphContext{
		Context: ctx,
		cancel:  cancel,
	}
}

func (g *GraphContext) Cancel(err error) {
	g.cancel(err)
}
