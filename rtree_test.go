
package rtree

import (
	"github.com/nukedzn/go-rtree/geom"
	"testing"
)

type RtreeItem struct {
	id   int
	rect *geom.Rect
}

func ( item *RtreeItem ) Mbr() ( *geom.Rect) {
	return item.rect
}

func TestRtreeSearchC2( t *testing.T ) {
	tree := NewRtree( 9 )
	nodeA := NewNode( 9 )
	nodeB := NewNode( 9 )
	tree.root.children, nodeA.parent = append( tree.root.children, nodeA ), tree.root
	tree.root.children, nodeB.parent = append( tree.root.children, nodeB ), tree.root
	nodeA.insert( &RtreeItem{ 1, &geom.Rect{ &geom.Point{ 1, 1 }, &geom.Point{ 3, 3 } } } )
	nodeB.insert( &RtreeItem{ 2, &geom.Rect{ &geom.Point{ 4, 4 }, &geom.Point{ 6, 6 } } } )

	results, cost := tree.Search( &geom.Point{ 2, 2 } )
	if len( results ) != 1 {
		t.Error( "Expected search to return 1 result(s), got:", len( results ) )
	} else if results[0].( *RtreeItem ).id != 1 {
		t.Error( "Expected search to return id 1, got: ", results[0].( *RtreeItem ).id )
	}

	if cost != 2 {
		t.Error( "Expected cost of search to be 1, got:", cost )
	}
}


func BenchmarkRtreeSearchC0( b *testing.B ) {
	tree := NewRtree( 9 )
	nodeA := NewNode( 9 )
	nodeB := NewNode( 9 )
	tree.root.children, nodeA.parent = append( tree.root.children, nodeA ), tree.root
	tree.root.children, nodeB.parent = append( tree.root.children, nodeB ), tree.root
	nodeA.insert( &RtreeItem{ 1, &geom.Rect{ &geom.Point{ 1, 1 }, &geom.Point{ 3, 3 } } } )
	nodeB.insert( &RtreeItem{ 2, &geom.Rect{ &geom.Point{ 4, 4 }, &geom.Point{ 6, 6 } } } )

	b.ReportAllocs()
	p := &geom.Point{ 0, 0 }
	for i := 0; i < b.N; i++ {
		tree.Search( p )
	}
}

func BenchmarkRtreeSearchC2( b *testing.B ) {
	tree := NewRtree( 9 )
	nodeA := NewNode( 9 )
	nodeB := NewNode( 9 )
	tree.root.children, nodeA.parent = append( tree.root.children, nodeA ), tree.root
	tree.root.children, nodeB.parent = append( tree.root.children, nodeB ), tree.root
	nodeA.insert( &RtreeItem{ 1, &geom.Rect{ &geom.Point{ 1, 1 }, &geom.Point{ 3, 3 } } } )
	nodeB.insert( &RtreeItem{ 2, &geom.Rect{ &geom.Point{ 4, 4 }, &geom.Point{ 6, 6 } } } )

	b.ReportAllocs()
	p := &geom.Point{ 2, 2 }
	for i := 0; i < b.N; i++ {
		tree.Search( p )
	}
}

func BenchmarkRtreeSearchC3( b *testing.B ) {
	tree := NewRtree( 9 )
	nodeA := NewNode( 9 )
	nodeAA := NewNode( 9 )
	nodeAB := NewNode( 9 )
	tree.root.children, nodeA.parent = append( tree.root.children, nodeA ), tree.root
	nodeA.children, nodeAA.parent = append( nodeA.children, nodeAA ), nodeA
	nodeA.children, nodeAB.parent = append( nodeA.children, nodeAB ), nodeA
	nodeAA.insert( &RtreeItem{ 1, &geom.Rect{ &geom.Point{ 1, 1 }, &geom.Point{ 3, 3 } } } )
	nodeAB.insert( &RtreeItem{ 2, &geom.Rect{ &geom.Point{ 4, 4 }, &geom.Point{ 6, 6 } } } )

	b.ReportAllocs()
	p := &geom.Point{ 2, 2 }
	for i := 0; i < b.N; i++ {
		tree.Search( p )
	}
}

