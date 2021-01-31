package graph

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Anxiety struct {
	FirstThoughts map[string]*Thought
	Thoughts      []*Thought
	Connections   map[Thought][]*Thought
	loop          bool
	lock          sync.RWMutex
}

func (a *Anxiety) AddThought(t *Thought) {
	a.lock.Lock()

	if a.FirstThoughts == nil {
		a.FirstThoughts = make(map[string]*Thought)
	}

	if t.First {
		a.FirstThoughts[t.Value] = t
	}

	a.Thoughts = append(a.Thoughts, t)
	a.lock.Unlock()
}

func (a *Anxiety) Connect(t1, t2 *Thought) {

	a.lock.Lock()
	if a.Connections == nil {
		a.Connections = make(map[Thought][]*Thought)
	}
	a.Connections[*t1] = append(a.Connections[*t1], t2)
	a.lock.Unlock()
}

func (a *Anxiety) RandomConnect() {
	a.loop = true
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(a.Thoughts)/4; i++ {
		randomThought0 := a.Thoughts[rand.Intn(len(a.Thoughts))]
		randomThought1 := a.Thoughts[rand.Intn(len(a.Thoughts))]
		a.Connect(randomThought0, randomThought1)
	}
}

func (a *Anxiety) String() {
	a.lock.RLock()
	s := ""
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

func (a *Anxiety) TraverseDepth(t *Thought) {
	a.lock.RLock()
	visitedThoughts := NewListMap()
	a.visit(t, visitedThoughts, 0)
	a.lock.RUnlock()
}

func (a *Anxiety) visit(n *Thought, visitedThoughts ListMap, level int) {

	if !a.loop && visitedThoughts.visited[n] > 0 {
		return
	} else if visitedThoughts.visited[n] > 634 {
		return
	}

	visitedThoughts.visit(n)

	// color.Set(pickColor(visitedThoughts.visited[n], a.loop))
	color.Set(pickColor(level, a.loop))
	// fmt.Printf("[%d] "+indent(level)+n.Value+"", visitedThoughts.visited[n])
	fmt.Printf(n.Value+"_%d     ", visitedThoughts.visited[n])
	color.Unset()

	rDuration := chance(2000, 500)
	time.Sleep(time.Duration(rDuration) * time.Millisecond)

	near := a.Connections[*n]
	for _, t := range near {
		a.visit(t, visitedThoughts, level+1)
	}
}

func (a *Anxiety) randomVisit(visitedThoughts ListMap, level int) {
	randomThought := a.Thoughts[chance(len(a.Thoughts), 0)]
	a.visit(randomThought, visitedThoughts, level)
}

func pickColor(n int, loop bool) color.Attribute {
	colors := map[int]color.Attribute{
		0: color.FgHiWhite,
		1: color.FgHiGreen,
		2: color.FgHiYellow,
		3: color.FgHiBlue,
		4: color.FgHiMagenta,
		5: color.FgHiCyan,
		6: color.FgHiRed,
	}
	return colors[n%len(colors)]
}

func indent(level int) string {
	str := ""
	for i := 0; i < level%100; i++ {
		str += "  "
	}
	return str
}

func chance(hi int, lo int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(hi-lo) + lo
}
