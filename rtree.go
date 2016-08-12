
package rtree

import (
	"math"
	"github.com/nukedzn/go-rtree/geom"
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

func ( tree *RTree ) Search( p *geom.Point ) ( []Item, int )  {
	n, nodes := tree.root, []*node{}
	results := []Item{}
	cost := 0

	for n != nil && n.Mbr().ContainsPoint( p ) {
		cost++
		if n.isLeaf() {
			for _, item := range n.items {
				if item.Mbr().ContainsPoint( p ) {
					results = append( results, item )
				}
			}
		} else {
			for _, child := range n.children {
				if child.Mbr().ContainsPoint( p ) {
					nodes = append( nodes, child )
				}
			}
		}

		if n = nil; len( nodes ) > 0 {
			n, nodes = nodes[0], nodes[1:]
		}
	}

	return results, cost
}

func ( tree *RTree ) chooseLeaf( n *node, item Item ) ( *node ) {
	r := item.Mbr()

	for !n.isLeaf() {
		var chosen *node
		var minCost, minArea float64 = math.Inf( 0 ), math.Inf( 0 )

		for _, child := range n.children {
			cost, area := child.insertionCost( r )
			if cost < minCost {
				minCost = cost
				chosen  = child
				if area < minArea {
					minArea = area
				}
			} else if cost == minCost && area < minArea {
				minArea = area
				chosen  = child
			}
		}

		n = chosen
	}

	return n
}

