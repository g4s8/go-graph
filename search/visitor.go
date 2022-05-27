package search

// Visitor for search functions
type Visitor interface {
	// Visit graph vertex
	Visit(int)
	// Visited returns true if vertex was visited
	Visited(int) bool
	// Continue search
	Continue() bool
}

type visitor []bool

// NewFullVisitor creates visitor for full search
func NewFullVisitor(g Graph) Visitor {
	v, _ := g.Size()
	bfs := make(visitor, v)
	return bfs
}

func (v visitor) Visit(id int) {
	v[id] = true
}

func (v visitor) Visited(id int) bool {
	return v[id]
}

func (v visitor) Continue() bool {
	return true
}

func (v visitor) Back() {
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

func (v *ShortPathVisitor) Visit(id int) {
	v.visited[id] = true
}

func (v *ShortPathVisitor) Visited(id int) bool {
	return v.visited[id]
}

func (v *ShortPathVisitor) Continue() bool {
	return !v.visited[v.v]
}

// Success returns true of expected vertex was visited
func (v *ShortPathVisitor) Success() bool {
	return v.visited[v.v]
}
