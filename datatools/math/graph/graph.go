package graph

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

type Node struct {
	value Item
}
