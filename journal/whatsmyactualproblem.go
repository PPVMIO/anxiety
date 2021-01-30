package journal

import (
	"github.com/ppelayo/anxiety/graph"
)

var WhatsMyActualProblem = graph.Thought{Value: "WhatsMyActualProblem", First: true}
var IveBeenCalledAnOverthinker = graph.Thought{Value: "IveBeenCalledAnOverthinker"}
var AndWhileThatsTrue = graph.Thought{Value: "AndWhileThatsTrue"}
var ItsAnOversimplificationWithNegativeConnotation = graph.Thought{Value: "ItsAnOversimplificationWithNegativeConnotation"}
var HonestlyThinkingIsMyStrength = graph.Thought{Value: "HonestlyThinkingIsMyStrength"}
var TheReasonIGotToCollege = graph.Thought{Value: "TheReasonIGotToCollege"}
var TheReasonIGotThroughBeingAMathMajor = graph.Thought{Value: "TheReasonIGotThroughBeingAMathMajor"}
var TheReasonIHaveAJobAsAnEngineer = graph.Thought{Value: "TheReasonIHaveAJobAsAnEngineer"}
var TheReasonWhyImAnAmazingEngineer = graph.Thought{Value: "TheReasonWhyImAnAmazingEngineer"}
var IsBecauseIThink = graph.Thought{Value: "IsBecauseIThink"}
var ICantHelpIt = graph.Thought{Value: "ICantHelpIt"}
var ItsLiterallyMyJobToOverthinkAndOverAnalyze = graph.Thought{Value: "ItsLiterallyMyJobToOverthinkAndOverAnalyze"}
var ThatsWhatIFuckingDo = graph.Thought{Value: "ThatsWhatIFuckingDo"}
var AndItsNotJustWorkEither = graph.Thought{Value: "AndItsNotJustWorkEither"}
var TheresAReasonIShowUpOnTime = graph.Thought{Value: "TheresAReasonIShowUpOnTime"}
var TheresAReasonImAlwaysAtEvents = graph.Thought{Value: "TheresAReasonImAlwaysAtEvents"}
var TheresAReasonImAGoodFriend = graph.Thought{Value: "TheresAReasonImAGoodFriend"}
var AndIDontBail = graph.Thought{Value: "AndIDontBail"}
var AndImThereForPeople = graph.Thought{Value: "AndImThereForPeople"}
var BecauseIThinkAhead = graph.Thought{Value: "BecauseIThinkAhead"}
var IPlan = graph.Thought{Value: "IPlan"}
var SoMuchOfWhatIConsiderToBeThePositivesInMyLife = graph.Thought{Value: "SoMuchOfWhatIConsiderToBeThePositivesInMyLife"}
var ComeFromOverthinkingThings = graph.Thought{Value: "ComeFromOverthinkingThings"}
var TheProblemIsApplyThisTypeOfThinkingAllTheTime = graph.Thought{Value: "TheProblemIsApplyThisTypeOfThinkingAllTheTime"}
var EspeciallyToSituationsThatIKnowKnowInMyHeart = graph.Thought{Value: "EspeciallyToSituationsThatIKnowKnowInMyHeart"}
var ICantControl = graph.Thought{Value: "ICantControl"}
var ThatDoesntStopMeFromTrying = graph.Thought{Value: "ThatDoesntStopMeFromTrying"}
var ToGuessPeoplesMotives = graph.Thought{Value: "ToGuessPeoplesMotives"}
var ToTryToReadThem = graph.Thought{Value: "ToTryToReadThem"}
var AndPeopleDoNotAlwaysActLogically = graph.Thought{Value: "AndPeopleDoNotAlwaysActLogically"}
var WhichIsGood = graph.Thought{Value: "WhichIsGood"}
var IWishIWasMoreLikePeople = graph.Thought{Value: "IWishIWasMoreLikePeople"}
var AndMyMindStartsToSpiral = graph.Thought{Value: "AndMyMindStartsToSpiral"}
var InAnAttemptToUnderstandAProblem = graph.Thought{Value: "InAnAttemptToUnderstandAProblem"}
var ThatHasNoSolution = graph.Thought{Value: "ThatHasNoSolution"}
var MyProblemIsIhaveADifficultTimeShuttingMyBrainOff = graph.Thought{Value: "MyProblemIsIhaveADifficultTimeShuttingMyBrainOff"}
var IHaveToKeepItBusy = graph.Thought{Value: "IHaveToKeepItBusy"}
var OtherwiseIFindMyselfMentally = graph.Thought{Value: "OtherwiseIFindMyselfMentally"}
var InMyOwnHead = graph.Thought{Value: "InMyOwnHead"}
var InThisInfiniteLoopsICantControl = graph.Thought{Value: "InThisInfiniteLoopsICantControl"}

func whatsMyActualProblem(a *graph.Anxiety) {
	a.AddThought(&WhatsMyActualProblem)
	a.AddThought(&IveBeenCalledAnOverthinker)
	a.AddThought(&AndWhileThatsTrue)
	a.AddThought(&ItsAnOversimplificationWithNegativeConnotation)
	a.AddThought(&HonestlyThinkingIsMyStrength)
	a.AddThought(&TheReasonIGotToCollege)
	a.AddThought(&TheReasonIGotThroughBeingAMathMajor)
	a.AddThought(&TheReasonIHaveAJobAsAnEngineer)
	a.AddThought(&TheReasonWhyImAnAmazingEngineer)
	a.AddThought(&IsBecauseIThink)
	a.AddThought(&ICantHelpIt)
	a.AddThought(&ItsLiterallyMyJobToOverthinkAndOverAnalyze)
	a.AddThought(&ThatsWhatIFuckingDo)
	a.AddThought(&AndItsNotJustWorkEither)
	a.AddThought(&TheresAReasonIShowUpOnTime)
	a.AddThought(&TheresAReasonImAlwaysAtEvents)
	a.AddThought(&TheresAReasonImAGoodFriend)
	a.AddThought(&AndIDontBail)
	a.AddThought(&AndImThereForPeople)
	a.AddThought(&BecauseIThinkAhead)
	a.AddThought(&IPlan)
	a.AddThought(&SoMuchOfWhatIConsiderToBeThePositivesInMyLife)
	a.AddThought(&ComeFromOverthinkingThings)
	a.AddThought(&TheProblemIsApplyThisTypeOfThinkingAllTheTime)
	a.AddThought(&EspeciallyToSituationsThatIKnowKnowInMyHeart)
	a.AddThought(&ICantControl)
	a.AddThought(&ThatDoesntStopMeFromTrying)
	a.AddThought(&ToGuessPeoplesMotives)
	a.AddThought(&ToTryToReadThem)
	a.AddThought(&AndPeopleDoNotAlwaysActLogically)
	a.AddThought(&WhichIsGood)
	a.AddThought(&IWishIWasMoreLikePeople)
	a.AddThought(&AndMyMindStartsToSpiral)
	a.AddThought(&InAnAttemptToUnderstandAProblem)
	a.AddThought(&ThatHasNoSolution)
	a.AddThought(&MyProblemIsIhaveADifficultTimeShuttingMyBrainOff)
	a.AddThought(&IHaveToKeepItBusy)
	a.AddThought(&OtherwiseIFindMyselfMentally)
	a.AddThought(&InMyOwnHead)
	a.AddThought(&InThisInfiniteLoopsICantControl)
}
