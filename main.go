package main

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ppelayo/anxiety/graph"
	"github.com/ppelayo/anxiety/journal"
)

func main() {

	anxiety := graph.Anxiety{}
	journal.BuildAnxiety(&anxiety)
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

	shouldSpiral := false
	err = survey.AskOne(&survey.Confirm{
		Message: "Spiral Mode",
	}, &shouldSpiral)
	handleErr(err)

	if shouldSpiral {
		anxiety.RandomConnect()
	}

	firstThought := anxiety.FirstThoughts[answers.Choice]
	anxiety.TraverseDepth(firstThought)

}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
