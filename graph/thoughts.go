package graph

type Thought struct {
	Value string
	First bool
	Level int
}

func (t *Thought) GetLevel() *int {
	return &t.Level
}

func (t *Thought) setLevel(newLevel int) {
	*t = Thought{Value: t.Value, First: t.First, Level: newLevel}
}
