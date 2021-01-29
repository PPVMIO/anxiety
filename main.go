package main

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ppelayo/anxiety/graph"
	"github.com/ppelayo/anxiety/thinking"
)

// func createAnxiety() graph.Anxiety {

// 	/* REAL Version */

// 	// thinkAboutMistakes := graph.Thought{Value: "Do you want to think about past mistakes :(?", First: true}
// 	// doYouMissHer := graph.Thought{Value: "Do you miss her?", First: true}
// 	// shouldYouText := graph.Thought{Value: "Should you text her?", First: true}
// 	// whatWillSheSay := graph.Thought{Value: "What's she going to say?", First: true}
// 	// willSheSayGoodthings := graph.Thought{Value: "What if she says good things?", First: true}
// 	// willSheSayBadThings := graph.Thought{Value: "What if she says bad things?", First: true}
// 	// areYouGoingToBeHurt := graph.Thought{Value: "Are you going to be hurt?", First: true}
// 	// isItTimeToMoveOn := graph.Thought{Value: "Is it time to move on?", First: true}
// 	// whatsThePoint := graph.Thought{Value: "What's the point of trying then?", First: true}
// 	// tooMuchStress := graph.Thought{Value: "This is too much stress", First: true}
// 	// guessIllDoNothing := graph.Thought{Value: "Guess I'll do nothing", First: true}

// 	// thoughts := []*graph.Thought{
// 	// 	&thinkAboutMistakes, &doYouMissHer, &shouldYouText, &whatWillSheSay, &willSheSayGoodthings, &willSheSayBadThings,
// 	// }

// 	// anxiety := graph.Anxiety{Thoughts: thoughts}

// 	// anxiety.Connect(&thinkAboutMistakes, &doYouMissHer)
// 	// anxiety.Connect(&thinkAboutMistakes, &whatWillSheSay)
// 	// anxiety.Connect(&doYouMissHer, &shouldYouText)
// 	// anxiety.Connect(&shouldYouText, &whatWillSheSay)
// 	// anxiety.Connect(&whatWillSheSay, &willSheSayGoodthings)
// 	// anxiety.Connect(&whatWillSheSay, &willSheSayBadThings)
// 	// anxiety.Connect(&willSheSayBadThings, &thinkAboutMistakes)

// 	/* Generic Version */

// 	a := graph.Thought{Value: "why think about thought 1?", First: true}
// 	b := graph.Thought{Value: "when you can think about thought 2?", First: true}
// 	c := graph.Thought{Value: "but don't forget about thought 3?", First: true}
// 	d := graph.Thought{Value: "and then you have to worry about 4?", First: true}
// 	e := graph.Thought{Value: "omfg", First: true}
// 	f := graph.Thought{Value: "this is too much", First: true}

// 	thoughts := []*graph.Thought{
// 		&a, &b, &c, &d, &e, &f,
// 	}

// 	anxiety := graph.Anxiety{Thoughts: thoughts}

// 	anxiety.Connect(&a, &d)
// 	anxiety.Connect(&d, &c)
// 	anxiety.Connect(&a, &b)
// 	anxiety.Connect(&a, &f)
// 	anxiety.Connect(&e, &f)
// 	anxiety.Connect(&d, &e)
// 	anxiety.Connect(&e, &a)
// 	anxiety.Connect(&c, &a)

// 	return anxiety
// }

func main() {

	anxiety := graph.Anxiety{}
	thinking.BuildAnxiety(&anxiety)
	firstThoughts := make([]string, 0)
	for _, t := range anxiety.FirstThoughts {
		firstThoughts = append(firstThoughts, t.Value)
	}

	questions := []*survey.Question{
		{
			Name: "choice",
			Prompt: &survey.Select{
				Message: "Choose a thought:",
				Options: firstThoughts,
				Default: firstThoughts[0],
			},
		},
	}
	answers := struct {
		Choice string `survey:"choice"`
	}{}
	err := survey.Ask(questions, &answers)
	handleErr(err)

	shouldLoop := false
	err = survey.AskOne(&survey.Confirm{
		Message: "Loop Mode On",
	}, &shouldLoop)
	handleErr(err)

	if shouldLoop {
		anxiety.RandomConnect()
		anxiety.RandomConnect()
		anxiety.RandomConnect()
	}

	firstThought := anxiety.FirstThoughts[answers.Choice]
	anxiety.TraverseDepth(firstThought)
	// anxiety.String()

}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func countLeadingSpace(line string) int {
	i := 0
	for _, runeValue := range line {
		if runeValue == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
