
package geom

import (
	"math"
)

type Rect struct {
	Min, Max *Point
}


func ( r *Rect ) Area() ( float64 ) {
	return ( r.Max.X - r.Min.X ) * ( r.Max.Y - r.Min.Y )
}

func ( r *Rect ) ContainsRect( rect *Rect ) ( bool ) {
	return r.Min.X <= rect.Min.X &&
		r.Min.Y <= rect.Min.Y &&
		r.Max.X >= rect.Max.X &&
		r.Max.Y >= rect.Max.Y
}

func ( r *Rect ) Extend( rect *Rect ) {
	r.Min.X = math.Min( r.Min.X, rect.Min.X )
	r.Min.Y = math.Min( r.Min.Y, rect.Min.Y )
	r.Max.X = math.Max( r.Max.X, rect.Max.X )
	r.Max.Y = math.Max( r.Max.Y, rect.Max.Y )
}

func ( r *Rect ) Extendn( rects []*Rect ) {
	for _, rect := range rects {
		r.Extend( rect )
	}
}

func ( r *Rect ) IntersectionArea( rect *Rect ) ( float64 ) {
	ri := Rect{}
	ri.Min.X = math.Max( r.Min.X, rect.Min.X )
	ri.Min.Y = math.Max( r.Min.Y, rect.Min.Y )
	ri.Max.X = math.Min( r.Max.X, rect.Max.X )
	ri.Max.Y = math.Min( r.Max.Y, rect.Max.Y )

	return math.Max( 0, r.Max.X - r.Min.X ) * math.Max( 0, r.Max.Y - r.Min.Y )
}

func ( r *Rect ) Margin() ( float64 ) {
	// sum of edge lengths
	return ( r.Max.X - r.Min.X ) + ( r.Max.Y - r.Min.Y )
}

