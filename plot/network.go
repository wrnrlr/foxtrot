package plot

import "gioui.org/layout"

type Node struct {
	ID int
}

type Edge struct {
	A, B int
}

type Network struct {
	Nodes []Node
	Edges []Edge
}

func (n Network) Layout(gtx *layout.Context) {
	for _, e := range n.Edges {

	}
	for _, n := range n.Nodes {

	}
}
