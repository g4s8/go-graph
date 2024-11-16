package search

// Visitor for search functions
type Visitor interface {
	// Visit graph vertex. Returns next visitor or nil to stop search.
	Visit(int) Visitor
	// Visited returns true if vertex was visited
	Visited(int) bool
}

type visitor []bool

// NewFullVisitor creates visitor for full search
func NewFullVisitor(g Graph) Visitor {
	s, _ := g.Size()
	return make(visitor, s)
}

func (v visitor) Visit(id int) Visitor {
	v[id] = true
	return v
}

func (v visitor) Visited(id int) bool {
	return v[id]
}

func (v visitor) reset() {
	for i := range v {
		v[i] = false
	}
}

// ShortPathVisitor stops search when found expected vertex
type ShortPathVisitor struct {
	v       int
	visited []bool
}

// NewShortPathVisitor creates new short path visitor for expected vertex
func NewShortPathVisitor(g Graph, vertex int) *ShortPathVisitor {
	v, _ := g.Size()
	var spv ShortPathVisitor
	spv.visited = make([]bool, v)
	spv.v = vertex
	return &spv
}

func (v *ShortPathVisitor) Visit(id int) Visitor {
	v.visited[id] = true
	if id == v.v {
		return nil
	}
	return v
}

func (v *ShortPathVisitor) Visited(id int) bool {
	return v.visited[id]
}

// Success returns true of expected vertex was visited
func (v *ShortPathVisitor) Success() bool {
	return v.visited[v.v]
}
