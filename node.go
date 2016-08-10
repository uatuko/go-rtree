
package rtree

import (
	"math"
	"github.com/nukedzn/go-rtree/geom"
)

type node struct {
	parent   *node
	children []*node
	items    []Item
	mbr      *geom.Rect
	minEntries, maxEntries uint16
}


func NewNode( M uint16 ) ( *node ) {
	return &node{
		minEntries: uint16( math.Ceil( float64( M ) * 0.4 ) ),
		maxEntries: M,
	}
}

func ( n *node ) Mbr() ( *geom.Rect ) {
	return n.mbr
}

func ( n *node ) area() ( float64 ) {
	if n.mbr == nil {
		return 0
	}

	return n.mbr.Area()
}

func ( n *node ) areaCost( r *geom.Rect ) ( cost float64, area float64 ) {
	area = n.area()
	cost = n.mergedArea( r ) - area
	return
}

func ( n *node ) extend( r *geom.Rect ) {
	if n.mbr == nil {
		n.mbr = r
	} else {
		n.mbr.Extend( r )
	}

	if n.parent != nil {
		n.parent.extend( n.mbr )
	}
}

func ( n *node ) insert( item Item ) {
	r := item.Mbr()
	n.extend( r )
	n.items = append( n.items, item )
	n.split()
}

func ( n *node ) isLeaf() ( bool ) {
	return len( n.children ) == 0
}

func ( n *node ) mergedArea( r *geom.Rect ) ( float64 ) {
	if n.mbr == nil {
		return r.Area()
	}

	return math.Max( n.mbr.Max.X, r.Max.X ) - math.Min( n.mbr.Min.X, r.Min.X ) *
		math.Max( n.mbr.Max.Y, r.Max.Y ) - math.Min( n.mbr.Min.Y, r.Min.Y )
}

func ( n *node ) size() ( uint16 ) {
	return uint16( len( n.children ) + len( n.items ) )
}

func ( n *node ) split() {
	if n.size() < n.maxEntries {
		return
	}

	// sort
	var itemsX []Item
	if n.isLeaf() {
		itemsX = n.items
	} else {
		itemsX = make( []Item, len( n.children ) )
		for idx, child := range n.children {
			itemsX[ idx ] = child
		}
	}

	itemsY := make( []Item, len( itemsX ) )
	copy( itemsY, itemsX )
	ItemSortBy( ItemSortMbrMinX ).Sort( itemsX )
	ItemSortBy( ItemSortMbrMinY ).Sort( itemsY )

	// choose split axis and split index
	var splitAxis []Item
	splitIdx, minMargin := uint16( 0 ), math.Inf( 0 )
	for _, items := range [][]Item{ itemsX, itemsY } {
		var margin float64 = 0
		minOverlap, minArea := math.Inf( 0 ), math.Inf( 0 )
		bestIdx := n.minEntries

		for idx := n.minEntries; idx <= ( uint16( len( items ) ) - n.minEntries ); idx++ {
			left := items[ 0 ].Mbr()
			right := items[ len( items ) - 1 ].Mbr()

			var i uint16
			for i = 1; i < uint16( len( items ) - 1 ); i++ {
				if i < idx {
					left.Extend( items[ i ].Mbr() )
				} else {
					right.Extend( items[ i ].Mbr() )
				}
			}

			margin += left.Margin() + right.Margin()
			overlap := left.IntersectionArea( right )
			area := left.Area() + right.Area()

			if overlap < minOverlap || ( overlap == minOverlap && area < minArea ) {
				minOverlap = overlap
				minArea = area
			}
		}

		if margin < minMargin {
			splitIdx = bestIdx
			splitAxis = items
			minMargin = margin
		}
	}


	// split
	n.mbr = nil
	nr := NewNode( n.maxEntries )
	if n.isLeaf() {
		n.items = make( []Item, splitIdx )
		nr.items = make( []Item, uint16( len( splitAxis ) ) - splitIdx )
	} else {
		n.children = make( []*node, splitIdx )
		nr.children = make( []*node, uint16( len( splitAxis ) ) - splitIdx )
	}

	for idx, item := range splitAxis {
		if child, ok := item.( *node ); ok {
			if uint16( idx ) < splitIdx {
				n.children[ idx ] = child
				n.extend( child.Mbr() )
			} else {
				nr.children[ uint16( idx ) - splitIdx ] = child
				nr.extend( child.Mbr() )
			}
		} else {
			if uint16( idx ) < splitIdx {
				n.items[ idx ] = item
				n.extend( item.Mbr() )
			} else {
				nr.items[ uint16( idx ) - splitIdx ] = item
				nr.extend( item.Mbr() )
			}
		}
	}


	// special handling for root node
	parent := n.parent
	if parent == nil {
		parent = n
		n := NewNode( n.maxEntries )
		n.items, n.children = parent.items, parent.children
		n.mbr = parent.mbr
		n.parent = parent

		parent.items = []Item{}
		parent.children = []*node{ n }
	}

	// update parent
	nr.parent = parent
	nr.parent.children = append( nr.parent.children, nr )
	nr.parent.extend( nr.Mbr() )

	// propergate changes upwards
	parent.split()
}

