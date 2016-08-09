
package rtree

import (
	"sort"
	"github.com/nukedzn/go-rtree/geom"
)

type rectSorter struct {
	rects  []*geom.Rect
	less   func ( r1, r2 *geom.Rect ) ( bool )
}


func ( s *rectSorter ) Len() ( int ) {
	return len( s.rects )
}

func ( s *rectSorter ) Swap( i, j int ) {
	s.rects[ i ], s.rects[ j ] = s.rects[ j ], s.rects[ i ]
}

func ( s *rectSorter ) Less( i, j int ) ( bool ) {
	return s.less( s.rects[ i ], s.rects[ j ] )
}



type RectSortBy func ( r1, r2 *geom.Rect ) ( bool )

func ( fn RectSortBy ) Sort( rects []*geom.Rect ) {
	s := &rectSorter{
		rects: rects,
		less: fn,
	}

	sort.Sort( s )
}


func RectSortMinX( r1, r2 *geom.Rect ) ( bool ) {
	return r1.Min.X < r2.Min.X
}

func RectSortMinY( r1, r2 *geom.Rect ) ( bool ) {
	return r1.Min.Y < r2.Min.Y
}


type itemSorter struct {
	items  []Item
	less   func ( i1, i2 Item ) ( bool )
}


func ( s *itemSorter ) Len() ( int ) {
	return len( s.items )
}

func ( s *itemSorter ) Swap( i, j int ) {
	s.items[ i ], s.items[ j ] = s.items[ j ], s.items[ i ]
}

func ( s *itemSorter ) Less( i, j int ) ( bool ) {
	return s.less( s.items[ i ], s.items[ j ] )
}



type ItemSortBy func ( i1, i2 Item ) ( bool )

func ( fn ItemSortBy ) Sort( items []Item ) {
	s := &itemSorter{
		items: items,
		less: fn,
	}

	sort.Sort( s )
}


func ItemSortMbrMinX( i1, i2 Item ) ( bool ) {
	return i1.Mbr().Min.X < i2.Mbr().Min.X
}

func ItemSortMbrMinY( i1, i2 Item ) ( bool ) {
	return i1.Mbr().Min.Y < i2.Mbr().Min.Y
}

