package thinking

import (
	"github.com/ppelayo/anxiety/graph"
)

var Test graph.Thought = graph.Thought{Value: "here is a thought that might be looped?"}

// willAnyoneCare := graph.Thought{Value: "will anyone care?"}
// doesItMatter := graph.Thought{Value: "does it matter?"}
// whatIfTheyDoCare := graph.Thought{Value: "what if they do care?"}
// whatIfTheyDont := graph.Thought{Value: "what if they don't care?"}

// func WhoseDogIsIt(a *graph.Anxiety, loop bool) {

// 	a.AddThought(&isMyWorkGoodEnough)
// 	a.AddThought(&willAnyoneCare)
// 	a.AddThought(&doesItMatter)
// 	a.AddThought(&whatIfTheyDoCare)
// 	a.AddThought(&whatIfTheyDont)

// 	a.Connect(&isMyWorkGoodEnough, &willAnyoneCare)
// 	a.Connect(&willAnyoneCare, &doesItMatter)
// 	a.Connect(&doesItMatter, &whatIfTheyDoCare)
// 	a.Connect(&doesItMatter, &whatIfTheyDont)
// }
