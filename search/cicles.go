package search

// HasCicle returns true if graph has a cicle
func HasCicle(g Graph) bool {
	vsize, _ := g.Size()
	visitor := new(cicleVisitor)
	visitor.visited = make([]bool, vsize)
	for v := 0; v < vsize; v++ {
		if checkCicle(g, v, visitor) {
			return true
		}
	}
	return false
}

type cicleVisitor struct {
	visited       []bool
	current, last int
	cicle         bool
}

func (cv *cicleVisitor) reset() {
	for i := range cv.visited {
		cv.visited[i] = false
	}
	cv.last = -1
	cv.current = -1
	cv.cicle = false
}

func (cv *cicleVisitor) Visit(v int) {
	if cv.visited[v] {
		cv.cicle = true
	}
	cv.visited[v] = true
	cv.last = cv.current
	cv.current = v
}

func (cv *cicleVisitor) Visited(v int) bool {
	return cv.last == v
}

func (cv *cicleVisitor) Continue() bool {
	return !cv.cicle
}

func checkCicle(g Graph, v int, visitor *cicleVisitor) bool {
	visitor.reset()
	DepthFirstSearch(g, v, visitor)
	return visitor.cicle
}
