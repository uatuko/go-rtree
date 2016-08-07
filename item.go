
package rtree

import (
	"github.com/nukedzn/go-rtree/geom"
)

type Item interface {
	Mbr() ( *geom.Rect )
}

