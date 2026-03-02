package resolver

import (
	"slices"
)

func (r *Resolver) syncNodes() {
	alreadySorted := slices.IsSortedFunc(r.nodes, func(a, b node) int {
		if a.slot < b.slot {
			return -1
		}
		if a.slot > b.slot {
			return 1
		}
		return 0
	})
	if !alreadySorted {
		slices.SortStableFunc(r.nodes, func(a, b node) int {
			if a.slot < b.slot {
				return -1
			}
			if a.slot > b.slot {
				return 1
			}
			return 0
		})
	}
	for i := len(r.nodes) - 1; i >= 1; i-- {
		r.nodes[i-1].coverage = r.nodes[i].slot - r.nodes[i-1].slot
	}

	r.nodes[len(r.nodes)-1].coverage = r.nodes[0].slot + (r.numSlots - r.nodes[len(r.nodes)-1].slot)
}
