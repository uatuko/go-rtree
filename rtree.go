
package rtree

import (
	"math"
)

type RTree struct {
	root *node
}


func NewRTree( maxNodeEntries uint16 ) ( *RTree ) {
	tree := &RTree{
		root: NewNode( maxNodeEntries ),
	}
	return tree
}

func ( tree *RTree ) Insert( item Item )  {
	leaf := tree.chooseLeaf( tree.root, item )
	leaf.insert( item )
}

func ( tree *RTree ) Search()  {

}

func ( tree *RTree ) chooseLeaf( n *node, item Item ) ( *node ) {
	r := item.Mbr()

	for !n.isLeaf() {
		var chosen *node
		var minCost, minArea float64 = math.Inf( 0 ), math.Inf( 0 )

		for _, child := range n.children {
			cost, area := child.areaCost( r )
			if cost < minCost {
				minCost = cost
				chosen  = child
				if area < minArea {
					minArea = area
				}
			} else if cost == minCost {
				if area < minArea {
					minArea = area
					chosen  = child
				}
			}
		}

		if chosen != nil {
			n = chosen
		} else {
			n = n.children[0]
		}
	}

	return n
}

