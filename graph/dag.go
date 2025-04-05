package graph

import "errors"

type Graph struct {
	name   string
	mNodes map[string]*Node
}

func NewGraph(name string) *Graph {
	return &Graph{
		name:   name,
		mNodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(name string) error {
	if _, ok := g.mNodes[name]; !ok {
		g.mNodes[name] = newNode(name)

		return nil
	} else {
		return errors.New("repeat name")
	}
}

func (g *Graph) AddEdge(from, to string) error {
	fromNode, ok := g.mNodes[from]
	if !ok {
		return errors.New("from node not exist")
	}

	toNode, ok := g.mNodes[to]
	if !ok {
		return errors.New("to node not exist")
	}

	fromNode.addNextNode(toNode)
	toNode.addPreNode(fromNode)
	return nil
}

func (g *Graph) IndependentNodes() map[string]*Node {
	results := map[string]*Node{}

	for name, node := range g.mNodes {
		if len(node.pres) == 0 {
			results[name] = g.mNodes[name]
		}
	}

	return results
}

func (g *Graph) GetNode(name string) *Node {
	return g.mNodes[name]
}

func (g *Graph) GetNodes() map[string]*Node {
	return g.mNodes
}
