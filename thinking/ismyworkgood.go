package thinking

import (
	"github.com/ppelayo/anxiety/graph"
)

var IsMyWorkGoodEnough graph.Thought = graph.Thought{Value: "IsMyWorkGoodEnough", First: true}
var WillAnyoneCare graph.Thought = graph.Thought{Value: "WillAnyoneCare"}
var DoesItMatter graph.Thought = graph.Thought{Value: "DoesItMatter"}
var WhatIfTheyDoCare graph.Thought = graph.Thought{Value: "WhatIfTheyDoCare"}
var WhatIfTheyDont graph.Thought = graph.Thought{Value: "WhatIfTheyDont"}

// func IsMyWorkGood(a *graph.Anxiety, loop bool) {

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
