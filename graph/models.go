package graph

type ListMap struct {
	visited map[*Thought]int
	keys    []*Thought
}

func NewListMap() ListMap {
	return ListMap{
		visited: make(map[*Thought]int, 0),
		keys:    make([]*Thought, 0),
	}
}

func (p *ListMap) visit(t *Thought) {
	p.visited[t]++
	p.keys = append(p.keys, t)
}
