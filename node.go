
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
		n.mbr.Min.X = math.Min( n.mbr.Min.X, r.Min.X )
		n.mbr.Min.Y = math.Min( n.mbr.Min.Y, r.Min.Y )
		n.mbr.Max.X = math.Max( n.mbr.Max.X, r.Max.X )
		n.mbr.Max.Y = math.Max( n.mbr.Max.Y, r.Max.Y )
	}

	if n.parent != nil && !n.parent.mbr.ContainsRect( n.mbr ) {
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
	if len( n.children ) == 0 {
		return false
	}

	return true
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

	// TODO: choose split axis, choose split index, split
	var rects []*geom.Rect
	if n.isLeaf() {
		for _, item := range n.items {
			rects = append( rects, item.Mbr() )
		}
	} else {
		for _, child := range n.children {
			rects = append( rects, child.mbr )
		}
	}

	sx, sy := rects, rects
	RectSortBy( RectSortMinX ).Sort( sx )
	RectSortBy( RectSortMinY ).Sort( sy )
}

