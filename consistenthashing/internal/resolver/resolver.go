package resolver

import (
	"fmt"
	"math"
)

type Resolver struct {
	targets              []string
	numSlots             int
	virtualizationFactor int

	nodes []node
}

type Option func(*Resolver)

func WithNumSlots(numSlots int) Option {
	if numSlots < 1 || numSlots > math.MaxUint16 {
		panic("numSlots out of range")
	}

	return func(r *Resolver) {
		r.numSlots = numSlots
	}
}

func WithVirtualizationFactor(virtualizationFactor int) Option {
	if virtualizationFactor < 0 {
		panic("virtualization factor must be positive")
	}

	return func(r *Resolver) {
		if virtualizationFactor > (r.numSlots/len(r.targets))-1 {
			panic("virtualization factor cannot be accommodated in current slot space")
		}
		r.virtualizationFactor = virtualizationFactor
	}
}

func NewResolver(targets []string, opts ...Option) Resolver {
	r := Resolver{
		targets:              targets,
		numSlots:             len(targets) * 4,
		virtualizationFactor: 1,
	}

	for _, opt := range opts {
		opt(&r)
	}

	totalNodeCount := len(targets) * (1 + r.virtualizationFactor)
	slotStep := float64(r.numSlots) / float64(totalNodeCount)
	slotOffset := float64(0)

	for i := 0; i < totalNodeCount; i++ {
		r.nodes = append(r.nodes, node{
			target: targets[i%len(r.targets)],
			slot:   int(math.Round(slotOffset + (slotStep * float64(i)))),
		})
	}

	r.syncNodes()

	return r
}

func (r *Resolver) PrintConfiguration() {
	for s := 0; s < r.numSlots; s++ {
		n := r.findNode(s)

		fmt.Printf("slot(%2d) -> %s\n", s, n.target)
	}
}
