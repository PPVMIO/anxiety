package graph

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Anxiety struct {
	Thoughts    []*Thought
	Connections map[Thought][]*Thought
	lock        sync.RWMutex
}

type Thought struct {
	Value string
	First bool
	Level int
}

func (t *Thought) GetLevel() *int {
	return &t.Level
}

func (t *Thought) setLevel(newLevel int) {
	// level := t.GetLevel()
	// *level = newLevel

	*t = Thought{Value: t.Value, First: t.First, Level: newLevel}
}

func (a *Anxiety) AddThought(n *Thought) {
	a.lock.Lock()
	a.Thoughts = append(a.Thoughts, n)
	a.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (a *Anxiety) Connect(n1, n2 *Thought) {

	a.lock.Lock()
	// n2.setLevel(n1.Level + 1)
	if a.Connections == nil {
		a.Connections = make(map[Thought][]*Thought)
	}
	a.Connections[*n1] = append(a.Connections[*n1], n2)
	// a.Connections[*n2] = append(a.Connections[*n2], n1)
	a.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (a *Anxiety) String() {
	a.lock.RLock()
	s := ""
	time.Sleep(1 * time.Second)
	for i := 0; i < len(a.Thoughts); i++ {
		s += a.Thoughts[i].Value + " -> "
		near := a.Connections[*a.Thoughts[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].Value + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	a.lock.RUnlock()
}

func (a *Anxiety) TraverseDepth() {
	a.lock.RLock()
	n := a.Thoughts[0]
	visited := NewListMap()
	a.visit(n, visited)
	a.lock.RUnlock()
}

func (a *Anxiety) visit(n *Thought, visited ListMap) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(2)
	time.Sleep(time.Duration(r) * time.Second)
	visited.Set(n, true)
	for _, v := range visited.keys {
		// fmt.Printf(v.Value+"%d", v.Level)
		fmt.Printf(fmt.Sprintf(v.Value+"%d ", v.Level))
	}
	// fmt.Println(fmt.Sprintf(n.Value+"%d", n.Level))
	fmt.Println()
	near := a.Connections[*n]
	for _, t := range near {
		a.visit(t, visited)
	}
}

// Traverse implements the BFS traversing algorithm
func (a *Anxiety) Traverse() {
	a.lock.RLock()
	q := ThoughtQueue{}
	q.New()
	n := a.Thoughts[0]
	q.Enqueue(*n)

	prevThoughts := NewListMap()
	// visited := make(map[*Thought]bool)
	level := 0
	for {
		if q.IsEmpty() {
			break
		}

		level++
		thought := q.Dequeue()
		near := a.Connections[*thought]

		// a.String()
		prevThoughts.Set(thought, true)

		for i := 0; i < len(near); i++ {
			j := near[i]
			thinkAbout(j, prevThoughts)

			q.Enqueue(*j)

			// if !visited[j] {
			// 	visited[j] = true
			// }
		}
		// visited[thought] = true

	}
	a.lock.RUnlock()
}

func thinkAbout(t *Thought, visited ListMap) {

	time.Sleep(1 * time.Second)
	for k := range visited.keys {
		// keys = append(keys, k)
		indent := ""
		for i := 0; i < visited.keys[k].Level; i++ {
			indent += "  "
		}
		fmt.Println(fmt.Sprintf(indent+visited.keys[k].Value+"%d ", visited.keys[k].Level))
	}
	fmt.Println()
	// fmt.Printf("Recently Visited %v ", t.Value)

}

type ListMap struct {
	m    map[*Thought]bool
	keys []*Thought
}

func NewListMap() ListMap {
	return ListMap{
		m:    make(map[*Thought]bool, 0),
		keys: make([]*Thought, 0),
	}
}

func (p *ListMap) Set(t *Thought, b bool) {
	p.m[t] = b
	p.keys = append(p.keys, t)
}
