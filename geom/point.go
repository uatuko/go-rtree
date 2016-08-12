
package geom

import(
	"fmt"
)

type Point struct {
	X, Y float64
}


func ( p *Point ) String() ( string ) {
	return fmt.Sprintf( "Point{ [%.6f, %.6f] }", p.X, p.Y )
}

