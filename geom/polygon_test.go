
package geom

import (
	"testing"
)


func TestPolygonMbr( t *testing.T ) {
	input := []Polygon{
		Polygon{ []*Point{ &Point{ 1, 1 }, &Point{ 2, 3 }, &Point{ 3, 2 } } },
		Polygon{ []*Point{ &Point{ -1.05, 0.25 }, &Point{ -2, 0.02 }, &Point{ 0.65, 1.78 }, &Point{ 1.25, -1.56 } } },
	}

	output := []Rect{
		Rect{ Min: &Point{ 1, 1 }, Max: &Point{ 3, 3 } },
		Rect{ Min: &Point{ -2, -1.56 }, Max: &Point{ 1.25, 1.78 } },
	}


	for idx, poly := range input {
		mbr := poly.Mbr()

		if mbr.Min.X != output[ idx ].Min.X || mbr.Min.Y != output[ idx ].Min.Y {
			t.Error( "Input:", poly, ", expected (min): ", output[ idx ].Min, ", got (min): ", mbr.Min )
		}

		if mbr.Max.X != output[ idx ].Max.X || mbr.Max.Y != output[ idx ].Max.Y {
			t.Error( "Input:", poly, ", expected (max): ", output[ idx ].Max, ", got (max): ", mbr.Max )
		}
	}
}

func BenchmarkPolygonMbr( b *testing.B ) {
	b.ReportAllocs()
	poly, _ := NewPolygon( []*Point{ &Point{ -1.05, 0.25 }, &Point{ -2, 0.02 }, &Point{ 0.65, 1.78 }, &Point{ 1.25, -1.56 } } )
	for i := 0; i < b.N; i++ {
		poly.Mbr()
	}
}

