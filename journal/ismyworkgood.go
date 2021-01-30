package journal

import (
	"github.com/ppelayo/anxiety/graph"
)

var IsMyWorkGoodEnough graph.Thought = graph.Thought{Value: "IsMyWorkGoodEnough", First: true}
var WillAnyoneCare graph.Thought = graph.Thought{Value: "WillAnyoneCare"}
var DoesItMatter graph.Thought = graph.Thought{Value: "DoesItMatter"}
var WhatIfTheyDoCare graph.Thought = graph.Thought{Value: "WhatIfTheyDoCare"}
var WhatIfTheyDont graph.Thought = graph.Thought{Value: "WhatIfTheyDont"}

func isMyWorkGoodEnough(a *graph.Anxiety) {
	a.AddThought(&IsMyWorkGoodEnough)
	a.AddThought(&WillAnyoneCare)
	a.AddThought(&DoesItMatter)
	a.AddThought(&WhatIfTheyDoCare)
	a.AddThought(&WhatIfTheyDont)

	a.Connect(&IsMyWorkGoodEnough, &WillAnyoneCare)
	a.Connect(&WillAnyoneCare, &DoesItMatter)
	a.Connect(&DoesItMatter, &WhatIfTheyDoCare)
	a.Connect(&DoesItMatter, &WhatIfTheyDont)

	// Loop
	a.Connect(&WhatIfTheyDont, &IsMyWorkGoodEnough)
}
