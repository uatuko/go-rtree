
package geom

import (
	"fmt"
	"math"
)

type Rect struct {
	Min, Max *Point
}


func NewRectFromRect( r *Rect ) ( *Rect ) {
	rr := &Rect{ &Point{ r.Min.X, r.Min.Y }, &Point{ r.Max.X, r.Max.Y } }
	return rr
}

func ( r *Rect ) Area() ( float64 ) {
	return ( r.Max.X - r.Min.X ) * ( r.Max.Y - r.Min.Y )
}

func ( r *Rect ) ContainsPoint( p *Point ) ( bool ) {
	return r.Min.X <= p.X &&
		r.Min.Y <= p.Y &&
		r.Max.X >= p.X &&
		r.Max.Y >= p.Y
}

func ( r *Rect ) ContainsRect( rect *Rect ) ( bool ) {
	return r.Min.X <= rect.Min.X &&
		r.Min.Y <= rect.Min.Y &&
		r.Max.X >= rect.Max.X &&
		r.Max.Y >= rect.Max.Y
}

func ( r *Rect ) IntersectionArea( rect *Rect ) ( float64 ) {
	ri := Rect{ &Point{}, &Point{} }
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

func ( r *Rect ) String() ( string ) {
	return fmt.Sprintf( "Rect{ [%.6f, %.6f] x [%.6f, %.6f] }", r.Min.X, r.Min.Y, r.Max.X, r.Max.Y )
}

func ( r *Rect ) Union( rect *Rect ) {
	r.Min.X = math.Min( r.Min.X, rect.Min.X )
	r.Min.Y = math.Min( r.Min.Y, rect.Min.Y )
	r.Max.X = math.Max( r.Max.X, rect.Max.X )
	r.Max.Y = math.Max( r.Max.Y, rect.Max.Y )
}

func ( r *Rect ) Unionn( rects []*Rect ) {
	for _, rect := range rects {
		r.Union( rect )
	}
}

