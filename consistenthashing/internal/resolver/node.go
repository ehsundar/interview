package resolver

import (
	"fmt"
)

type node struct {
	target   string
	slot     int
	coverage int
}

func (n node) String() string {
	return fmt.Sprintf("%s(%d)", n.target, n.slot)
}
