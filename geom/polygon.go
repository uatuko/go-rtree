
package geom

import (
	"errors"
)

type Polygon struct {
	points []*Point
}


func NewPolygon( points []*Point ) ( *Polygon, error ) {
	if len( points ) < 3 {
		return nil, errors.New( "Need at least three points to create a polygon" )
	}

	poly := &Polygon{ points }
	return poly, nil
}

func ( poly *Polygon ) Mbr() ( *Rect )  {
	min, max := *poly.points[0], *poly.points[0]
	mbr := &Rect{ &min, &max }

	for i := 1; i < len( poly.points ); i++ {
		pt := poly.points[i]
		if pt.X < mbr.Min.X {
			mbr.Min.X = pt.X
		} else if pt.X > mbr.Max.X {
			mbr.Max.X = pt.X
		}

		if pt.Y < mbr.Min.Y {
			mbr.Min.Y = pt.Y
		} else if pt.Y > mbr.Max.Y {
			mbr.Max.Y = pt.Y
		}
	}

	return mbr
}

