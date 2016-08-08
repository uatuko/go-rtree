
package geom

import (
	"testing"
)


func TestRectArea( t *testing.T ) {
	input := []Rect{
		Rect{ &Point{ 1, 1 }, &Point{ 3, 3 } },
		Rect{ &Point{ -2.56, 1.44 }, &Point{ 0.46, 1.45 } },
	}

	output := []float64{
		4,
		0.0302,
	}


	for idx, rect := range input {
		area := float64( int( rect.Area() * 1e6 ) ) / 1e6

		if area != output[ idx ] {
			t.Error( "Input: [", rect.Min, ",", rect.Max , "], expected: ", output[ idx ], ", got:", area )
		}
	}
}

func TestRectContainsRect( t *testing.T ) {
	input1 := []Rect{
		Rect{ &Point{ 1, 1 }, &Point{ 3, 3 } },
		Rect{ &Point{ 1, 1 }, &Point{ 3, 3 } },
		Rect{ &Point{ 1, 1 }, &Point{ 3, 3 } },
	}
	input2 := []*Rect{
		&Rect{ &Point{ 1, 1 }, &Point{ 3, 3 } },
		&Rect{ &Point{ 1.56, 2.22 }, &Point{ 1.99, 2.58 } },
		&Rect{ &Point{ -2.56, 1.44 }, &Point{ 0.46, 1.45 } },
	}
	output := []bool{
		true,
		true,
		false,
	}

	for idx, rect := range input1 {
		result := rect.ContainsRect( input2[ idx ] )
		if result != output[ idx ] {
			t.Error( "Input: [", rect, "contains", input2[ idx ], "], expected:", output[ idx ], ", got:", result )
		}
	}
}

func TestRectMargin( t *testing.T ) {
	input := []Rect{
		Rect{ &Point{ 1, 1 }, &Point{ 3, 3 } },
		Rect{ &Point{ -2.56, 1.44 }, &Point{ 0.46, 1.45 } },
	}

	output := []float64{
		4,
		3.03,
	}


	for idx, rect := range input {
		margin := float64( int( rect.Margin() * 1e6 ) ) / 1e6

		if margin != output[ idx ] {
			t.Error( "Input:", rect, ", expected: ", output[ idx ], ", got:", margin )
		}
	}
}

