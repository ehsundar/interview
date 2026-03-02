package resolver

import (
	"slices"
)

func (r *Resolver) AddTarget(target string) error {
	r.syncNodes()

	sortedByCoverage := slices.Clone(r.nodes)
	slices.SortStableFunc(sortedByCoverage, func(a, b node) int {
		// sort and reverse by coverage
		if a.coverage < b.coverage {
			return 1
		}
		if a.coverage > b.coverage {
			return -1
		}
		return 0
	})

	selectedNodes := sortedByCoverage[:r.virtualizationFactor+1]

	for _, s := range selectedNodes {
		newNode := node{
			target: target,
			slot:   (s.slot + s.coverage/2) % r.numSlots,
		}
		r.nodes = append(r.nodes, newNode)
	}

	r.syncNodes()
	return nil
}
