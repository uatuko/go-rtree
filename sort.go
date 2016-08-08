
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

