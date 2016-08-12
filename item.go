
package rtree

import (
	"github.com/uditha-atukorala/go-rtree/geom"
)

type Item interface {
	Mbr() ( *geom.Rect )
}

