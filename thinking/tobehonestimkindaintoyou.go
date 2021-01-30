package thinking

import (
	"github.com/ppelayo/anxiety/graph"
)

var ToBeHonestImKindaIntoYou graph.Thought = graph.Thought{Value: "ToBeHonestImKindaIntoYou", First: true}
var LikeYoureReallyCool graph.Thought = graph.Thought{Value: "LikeYoureReallyCool"}
var AndWeDontTalkMuch graph.Thought = graph.Thought{Value: "AndWeDontTalkMuch"}
var WellIGuessWeKindaDo graph.Thought = graph.Thought{Value: "WellIGuessWeKindaDo"}
var ButItCanBeKindaDumb graph.Thought = graph.Thought{Value: "ButItCanBeKindaDumb"}
var ItJustALotOfMemes graph.Thought = graph.Thought{Value: "ItJustALotOfMemes"}
var ButAtLeastWeShareTheSameTasteInMemes graph.Thought = graph.Thought{Value: "ButAtLeastWeShareTheSameTasteInMemes"}
var AndThatsImportantIn2021Right graph.Thought = graph.Thought{Value: "AndThatsImportantIn2021Right"}
var AndYoureAttractive graph.Thought = graph.Thought{Value: "AndYoureAttractive"}
var ButItsKindaFuckedUpNo graph.Thought = graph.Thought{Value: "ButItsKindaFuckedUpNo"}
var WhoIveBeenWith graph.Thought = graph.Thought{Value: "WhoIveBeenWith"}
var WhoYouveBeenFriendsWith graph.Thought = graph.Thought{Value: "WhoYouveBeenFriendsWith"}
var Sigh graph.Thought = graph.Thought{Value: "Sigh"}
var ItsAMess graph.Thought = graph.Thought{Value: "ItsAMess"}
var ImAMess graph.Thought = graph.Thought{Value: "ImAMess"}
var MaybeItsJustNotWorthIt graph.Thought = graph.Thought{Value: "MaybeItsJustNotWorthIt"}
var ButItCouldBe graph.Thought = graph.Thought{Value: "ButItCouldBe"}
var WhyNot graph.Thought = graph.Thought{Value: "WhyNot"}
var IGuessTheresALotOfReasonsLol graph.Thought = graph.Thought{Value: "IGuessTheresALotOfReasonsLol"}

func toBeHonestImKindaIntoYouAnxiety(a *graph.Anxiety) {
	// to be honest I'm kinda into you
	a.AddThought(&ToBeHonestImKindaIntoYou)
	a.AddThought(&LikeYoureReallyCool)
	a.AddThought(&AndWeDontTalkMuch)
	a.AddThought(&WellIGuessWeKindaDo)
	a.AddThought(&ButItCanBeKindaDumb)
	a.AddThought(&ItJustALotOfMemes)
	a.AddThought(&ButAtLeastWeShareTheSameTasteInMemes)
	a.AddThought(&AndThatsImportantIn2021Right)
	a.AddThought(&AndYoureAttractive)
	a.AddThought(&ButItsKindaFuckedUpNo)
	a.AddThought(&WhoIveBeenWith)
	a.AddThought(&WhoYouveBeenFriendsWith)
	a.AddThought(&Sigh)
	a.AddThought(&ItsAMess)
	a.AddThought(&ImAMess)
	a.AddThought(&MaybeItsJustNotWorthIt)
	a.AddThought(&ButItCouldBe)
	a.AddThought(&WhyNot)
	a.AddThought(&IGuessTheresALotOfReasonsLol)

	a.Connect(&ToBeHonestImKindaIntoYou, &LikeYoureReallyCool)
	a.Connect(&ToBeHonestImKindaIntoYou, &AndWeDontTalkMuch)
	a.Connect(&AndWeDontTalkMuch, &WellIGuessWeKindaDo)
	a.Connect(&AndWeDontTalkMuch, &ButItCanBeKindaDumb)
	a.Connect(&AndWeDontTalkMuch, &ButAtLeastWeShareTheSameTasteInMemes)
	a.Connect(&ButAtLeastWeShareTheSameTasteInMemes, &AndThatsImportantIn2021Right)
	a.Connect(&ToBeHonestImKindaIntoYou, &AndYoureAttractive)
	a.Connect(&ToBeHonestImKindaIntoYou, &ButItsKindaFuckedUpNo)
	a.Connect(&ButItsKindaFuckedUpNo, &WhoIveBeenWith)
	a.Connect(&WhoIveBeenWith, &Sigh)
	a.Connect(&WhoIveBeenWith, &ItsAMess)
	a.Connect(&WhoIveBeenWith, &ImAMess)
	a.Connect(&ButItsKindaFuckedUpNo, &MaybeItsJustNotWorthIt)
	a.Connect(&ButItsKindaFuckedUpNo, &ButItCouldBe)
	a.Connect(&ButItsKindaFuckedUpNo, &WhyNot)
	a.Connect(&WhyNot, &IGuessTheresALotOfReasonsLol)

	// Loop
	a.Connect(&IGuessTheresALotOfReasonsLol, &ButItsKindaFuckedUpNo)
}
