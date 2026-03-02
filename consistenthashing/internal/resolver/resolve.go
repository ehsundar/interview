package resolver

import (
	"context"
)

func (r *Resolver) Resolve(ctx context.Context, key string) (string, error) {
	sh := hashKey(key)
	slot := int(sh) % (r.numSlots)
	n := r.findNode(slot)
	return n.target, nil
}

func (r *Resolver) findNode(slot int) node {
	if len(r.nodes) == 0 {
		return node{}
	}

	for i := len(r.nodes) - 1; i >= 0; i-- {
		if slot >= r.nodes[i].slot {
			return r.nodes[i]
		}
	}
	return r.nodes[len(r.nodes)-1]
}

func binarySearch(nodes []node, slot int) node {
	start, end := 0, len(nodes)
	for start < end {
		mid := start + (end-start)/2
		if nodes[mid].slot == slot {
			return nodes[mid]
		}
		if nodes[mid].slot < slot {
			start = mid + 1
		}
		if nodes[mid].slot > slot {
			end = mid - 1
		}
	}
	return nodes[start]
}
