
package rtree

import (
	"github.com/nukedzn/go-rtree/geom"
)

type node struct {
	parent   *node
	children []*node
	items    []*Item
	mbr      *geom.Rect
}

