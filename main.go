package main

import (
	"github.com/ppelayo/anxiety/graph"
)

func createAnxiety() graph.Anxiety {

	// thinkAboutMistakes := graph.Thought{Value: "Do you want to think about past mistakes :(?", First: true}
	// doYouMissHer := graph.Thought{Value: "Do you miss her?", First: true}
	// shouldYouText := graph.Thought{Value: "Should you text her?", First: true}
	// whatWillSheSay := graph.Thought{Value: "What's she going to say?", First: true}
	// willSheSayGoodthings := graph.Thought{Value: "What if she says good things?", First: true}
	// willSheSayBadThings := graph.Thought{Value: "What if she says bad things?", First: true}
	// areYouGoingToBeHurt := graph.Thought{Value: "Are you going to be hurt?", First: true}
	// isItTimeToMoveOn := graph.Thought{Value: "Is it time to move on?", First: true}
	// whatsThePoint := graph.Thought{Value: "What's the point of trying then?", First: true}
	// tooMuchStress := graph.Thought{Value: "This is too much stress", First: true}
	// guessIllDoNothing := graph.Thought{Value: "Guess I'll do nothing", First: true}

	a := graph.Thought{Value: "a", First: true}
	b := graph.Thought{Value: "b", First: true}
	c := graph.Thought{Value: "c", First: true}
	d := graph.Thought{Value: "d", First: true}
	e := graph.Thought{Value: "e", First: true}
	f := graph.Thought{Value: "f", First: true}

	thoughts := []*graph.Thought{
		&a, &b, &c, &d, &e, &f,
	}

	anxiety := graph.Anxiety{Thoughts: thoughts}

	anxiety.Connect(&a, &d)
	anxiety.Connect(&d, &c)
	anxiety.Connect(&a, &b)
	anxiety.Connect(&a, &f)
	anxiety.Connect(&e, &f)
	anxiety.Connect(&d, &e)
	anxiety.Connect(&e, &a)
	// anxiety.Connect(&e, &a)
	// anxiety.Connect(&c, &a)
	// a.Connect(&nB, &nE)
	// a.Connect(&nC, &nE)
	// a.Connect(&nE, &nF)
	// a.Connect(&nD, &nA)

	return anxiety
}

func main() {
	a := createAnxiety()
	// a.String()
	a.TraverseDepth()
}
