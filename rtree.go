
package rtree


type RTree struct {
	root *node
}


func NewRtree() ( *RTree ) {
	tree := &RTree{ root: &node{} }
	return tree
}

func ( tree *RTree ) Insert( item *Item )  {
	leaf := tree.chooseLeaf( tree.root, item )
	leaf.items = append( leaf.items, item )

	// TODO: check overflow and balance tree?
}

func ( tree *RTree ) Search()  {

}

func ( tree * RTree ) chooseLeaf( n *node, item *Item ) ( *node ) {
	return nil
}

