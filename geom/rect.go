
package geom

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

func ( r *Rect ) Margin() ( float64 ) {
	// sum of edge lengths
	return ( r.Max.X - r.Min.X ) + ( r.Max.Y - r.Min.Y )
}

