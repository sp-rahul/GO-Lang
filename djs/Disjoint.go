

package disjoint

type Element struct {
	parent *Element    // Parent element
	rank   int         // Rank (approximate depth) of the subtree with this element as root
	Data   interface{} // Arbitrary user-provided payload
}

func NewElement() *Element {
	s := &Element{}
	s.parent = s
	return s
}

func (e *Element) Find() *Element {
	for e.parent != e {
		e.parent = e.parent.parent
		e = e.parent
	}
	return e
}


func Union(e1, e2 *Element) {
	// Ensure the two Elements aren't already part of the same union.
	e1Root := e1.Find()
	e2Root := e2.Find()
	if e1Root == e2Root {
		return
	}

	// Create a union by making the shorter tree point to the root of the
	// larger tree.
	switch {
	case e1Root.rank < e2Root.rank:
		e1Root.parent = e2Root
	case e1Root.rank > e2Root.rank:
		e2Root.parent = e1Root
	default:
		e2Root.parent = e1Root
		e1Root.rank++
	}
}