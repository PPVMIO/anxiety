package journal

import (
	"github.com/ppelayo/anxiety/graph"
)

var WhatILikeAboutThisProject = graph.Thought{Value: "WhatILikeAboutThisProject", First: true}
var IsThatItsActuallySuperPersonalLol = graph.Thought{Value: "IsThatItsActuallySuperPersonalLol"}
var LikeReallyPersonal = graph.Thought{Value: "LikeReallyPersonal"}
var LikeIReferenceALotOfPersonalIssues = graph.Thought{Value: "LikeIReferenceALotOfPersonalIssues"}
var ThatINeverTalkAboutInPerson = graph.Thought{Value: "ThatINeverTalkAboutInPerson"}
var IMakeSomeReferencesToIndividualPeople = graph.Thought{Value: "IMakeSomeReferencesToIndividualPeople"}
var AndIFeelLikeTheyWouldntCare = graph.Thought{Value: "AndIFeelLikeTheyWouldntCare"}
var ActuallyTheyWouldCare = graph.Thought{Value: "ActuallyTheyWouldCare"}
var AndIShouldBringItUp = graph.Thought{Value: "AndIShouldBringItUp"}
var OrPrayTheyNeverSeeIt = graph.Thought{Value: "OrPrayTheyNeverSeeIt"}
var MostPeopleInMyCircleDontKnowHowToRunThis = graph.Thought{Value: "MostPeopleInMyCircleDontKnowHowToRunThis"}
var ButEvenIfYoureSomeoneThatDidKnowHowToRunThis = graph.Thought{Value: "ButEvenIfYoureSomeoneThatDidKnowHowToRunThis"}
var ThenYoureProbablyANerd = graph.Thought{Value: "ThenYoureProbablyANerd"}
var MoreImportantlyThatMeansYouWentThroughALotOfEffort = graph.Thought{Value: "MoreImportantlyThatMeansYouWentThroughALotOfEffort"}
var ToLearnAboutMe = graph.Thought{Value: "ToLearnAboutMe"}
var AndThatMeansALot = graph.Thought{Value: "AndThatMeansALot"}
var IWouldSayThisIsActuallyAJournal = graph.Thought{Value: "IWouldSayThisIsActuallyAJournal"}
var DisguisedAsAnArtProject = graph.Thought{Value: "DisguisedAsAnArtProject"}
var DisguisedAsAProgram = graph.Thought{Value: "DisguisedAsAProgram"}
var InceptionFuckingShit = graph.Thought{Value: "InceptionFuckingShit"}
var AnywaysIfYoureSeeingThis = graph.Thought{Value: "AnywaysIfYoureSeeingThis"}
var KnowThatIAppreciateYou = graph.Thought{Value: "KnowThatIAppreciateYou"}

func whatILikeAboutThisProject(a *graph.Anxiety) {

	a.AddThought(&WhatILikeAboutThisProject)
	a.AddThought(&IsThatItsActuallySuperPersonalLol)
	a.AddThought(&LikeReallyPersonal)
	a.AddThought(&LikeIReferenceALotOfPersonalIssues)
	a.AddThought(&ThatINeverTalkAboutInPerson)
	a.AddThought(&IMakeSomeReferencesToIndividualPeople)
	a.AddThought(&AndIFeelLikeTheyWouldntCare)
	a.AddThought(&ActuallyTheyWouldCare)
	a.AddThought(&AndIShouldBringItUp)
	a.AddThought(&OrPrayTheyNeverSeeIt)
	a.AddThought(&MostPeopleInMyCircleDontKnowHowToRunThis)
	a.AddThought(&ButEvenIfYoureSomeoneThatDidKnowHowToRunThis)
	a.AddThought(&ThenYoureProbablyANerd)
	a.AddThought(&MoreImportantlyThatMeansYouWentThroughALotOfEffort)
	a.AddThought(&ToLearnAboutMe)
	a.AddThought(&AndThatMeansALot)
	a.AddThought(&IWouldSayThisIsActuallyAJournal)
	a.AddThought(&DisguisedAsAnArtProject)
	a.AddThought(&DisguisedAsAProgram)
	a.AddThought(&InceptionFuckingShit)
	a.AddThought(&AnywaysIfYoureSeeingThis)
	a.AddThought(&KnowThatIAppreciateYou)

	a.Connect(&WhatILikeAboutThisProject, &IsThatItsActuallySuperPersonalLol)
	a.Connect(&IsThatItsActuallySuperPersonalLol, &LikeReallyPersonal)
	a.Connect(&IsThatItsActuallySuperPersonalLol, &LikeIReferenceALotOfPersonalIssues)
	a.Connect(&LikeIReferenceALotOfPersonalIssues, &ThatINeverTalkAboutInPerson)
	a.Connect(&WhatILikeAboutThisProject, &IMakeSomeReferencesToIndividualPeople)
	a.Connect(&IMakeSomeReferencesToIndividualPeople, &AndIFeelLikeTheyWouldntCare)
	a.Connect(&AndIFeelLikeTheyWouldntCare, &ActuallyTheyWouldCare)
	a.Connect(&AndIFeelLikeTheyWouldntCare, &AndIShouldBringItUp)
	a.Connect(&AndIFeelLikeTheyWouldntCare, &OrPrayTheyNeverSeeIt)
	a.Connect(&WhatILikeAboutThisProject, &MostPeopleInMyCircleDontKnowHowToRunThis)
	a.Connect(&MostPeopleInMyCircleDontKnowHowToRunThis, &ButEvenIfYoureSomeoneThatDidKnowHowToRunThis)
	a.Connect(&ButEvenIfYoureSomeoneThatDidKnowHowToRunThis, &ThenYoureProbablyANerd)
	a.Connect(&ButEvenIfYoureSomeoneThatDidKnowHowToRunThis, &MoreImportantlyThatMeansYouWentThroughALotOfEffort)
	a.Connect(&MoreImportantlyThatMeansYouWentThroughALotOfEffort, &ToLearnAboutMe)
	a.Connect(&MoreImportantlyThatMeansYouWentThroughALotOfEffort, &AndThatMeansALot)
	a.Connect(&AndThatMeansALot, &AnywaysIfYoureSeeingThis)
	a.Connect(&AnywaysIfYoureSeeingThis, &KnowThatIAppreciateYou)
	a.Connect(&WhatILikeAboutThisProject, &IWouldSayThisIsActuallyAJournal)
	a.Connect(&IWouldSayThisIsActuallyAJournal, &DisguisedAsAnArtProject)
	a.Connect(&DisguisedAsAnArtProject, &DisguisedAsAProgram)
	a.Connect(&DisguisedAsAProgram, &InceptionFuckingShit)
	a.Connect(&DisguisedAsAProgram, &InceptionFuckingShit)

}
