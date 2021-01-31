package journal

import (
	"github.com/ppelayo/anxiety/graph"
)

var ItAintACakeWalkOverHere = graph.Thought{Value: "ItAintACakeWalkOverHere", First: true}
var IConsiderMyselfAHighFunctioningAnxiousPerson = graph.Thought{Value: "IConsiderMyselfAHighFunctioningAnxiousPerson"}
var InTheSenseThatIWouldConsiderMyselfHighlyCapable = graph.Thought{Value: "InTheSenseThatIWouldConsiderMyselfHighlyCapable"}
var AsFarAs = graph.Thought{Value: "AsFarAs"}
var Work = graph.Thought{Value: "Work"}
var PhysicalFitness = graph.Thought{Value: "PhysicalFitness"}
var Smart = graph.Thought{Value: "Smart"}
var PopularLol = graph.Thought{Value: "PopularLol"}
var GoodAtALotOFThings = graph.Thought{Value: "GoodAtALotOFThings"}
var ButActuallyImAnxiousALot = graph.Thought{Value: "ButActuallyImAnxiousALot"}
var ItsAMajorMotivatorForMe = graph.Thought{Value: "ItsAMajorMotivatorForMe"}
var Anxiety = graph.Thought{Value: "Anxiety"}
var TheFearOfFailing = graph.Thought{Value: "TheFearOfFailing"}
var TheFearOfMissingOut = graph.Thought{Value: "TheFearOfMissingOut"}
var TheFearOfNotBeingEnough = graph.Thought{Value: "TheFearOfNotBeingEnough"}
var ItsAllRelated = graph.Thought{Value: "ItsAllRelated"}
var IWantEverythingIDoToBeDeliberate = graph.Thought{Value: "IWantEverythingIDoToBeDeliberate"}
var AndThoughtThrough = graph.Thought{Value: "AndThoughtThrough"}
var AndPlanned = graph.Thought{Value: "AndPlanned"}
var ItMakesMeWorkHarder = graph.Thought{Value: "ItMakesMeWorkHarder"}
var ItMakesMePushMyself = graph.Thought{Value: "ItMakesMePushMyself"}
var ItMakesMeSuccessfulByLotsOfStandards = graph.Thought{Value: "ItMakesMeSuccessfulByLotsOfStandards"}
var ButItAlsoMakesItReallyDifficultToAcceptFailure = graph.Thought{Value: "ButItAlsoMakesItReallyDifficultToAcceptFailure"}
var IPutTimeIntoEverythingIDo = graph.Thought{Value: "IPutTimeIntoEverythingIDo"}
var AndWhenTheEffortIPutDoesNotMatchTheOutcome = graph.Thought{Value: "AndWhenTheEffortIPutDoesNotMatchTheOutcome"}
var WhetherThats = graph.Thought{Value: "WhetherThats"}
var Relationships = graph.Thought{Value: "Relationships"}
var MyCareer = graph.Thought{Value: "MyCareer"}
var MyFriendships = graph.Thought{Value: "MyFriendships"}
var EvenSomethingAsSimpleAsWeekendPlans = graph.Thought{Value: "EvenSomethingAsSimpleAsWeekendPlans"}
var OrASimpleText = graph.Thought{Value: "OrASimpleText"}
var ITakeItReallyHard = graph.Thought{Value: "ITakeItReallyHard"}
var IShutDown = graph.Thought{Value: "IShutDown"}
var PretendLikeIDontCare = graph.Thought{Value: "PretendLikeIDontCare"}
var OrImBusyWithOtherThings = graph.Thought{Value: "OrImBusyWithOtherThings"}
var CantFailAtSomethingThatNeverMattered = graph.Thought{Value: "CantFailAtSomethingThatNeverMattered"}
var LikeISaid = graph.Thought{Value: "LikeISaid"}
var AintACakeWalkOverHere = graph.Thought{Value: "AintACakeWalkOverHere"}

func itAintACakeWalkOverHere(a *graph.Anxiety) {
	a.AddThought(&ItAintACakeWalkOverHere)
	a.AddThought(&IConsiderMyselfAHighFunctioningAnxiousPerson)
	a.AddThought(&InTheSenseThatIWouldConsiderMyselfHighlyCapable)
	a.AddThought(&AsFarAs)
	a.AddThought(&Work)
	a.AddThought(&PhysicalFitness)
	a.AddThought(&Smart)
	a.AddThought(&PopularLol)
	a.AddThought(&GoodAtALotOFThings)
	a.AddThought(&ButActuallyImAnxiousALot)
	a.AddThought(&ItsAMajorMotivatorForMe)
	a.AddThought(&Anxiety)
	a.AddThought(&TheFearOfFailing)
	a.AddThought(&TheFearOfMissingOut)
	a.AddThought(&TheFearOfNotBeingEnough)
	a.AddThought(&ItsAllRelated)
	a.AddThought(&IWantEverythingIDoToBeDeliberate)
	a.AddThought(&AndThoughtThrough)
	a.AddThought(&AndPlanned)
	a.AddThought(&ItMakesMeWorkHarder)
	a.AddThought(&ItMakesMePushMyself)
	a.AddThought(&ItMakesMeSuccessfulByLotsOfStandards)
	a.AddThought(&ButItAlsoMakesItReallyDifficultToAcceptFailure)
	a.AddThought(&IPutTimeIntoEverythingIDo)
	a.AddThought(&AndWhenTheEffortIPutDoesNotMatchTheOutcome)
	a.AddThought(&WhetherThats)
	a.AddThought(&Relationships)
	a.AddThought(&MyCareer)
	a.AddThought(&MyFriendships)
	a.AddThought(&EvenSomethingAsSimpleAsWeekendPlans)
	a.AddThought(&OrASimpleText)
	a.AddThought(&ITakeItReallyHard)
	a.AddThought(&IShutDown)
	a.AddThought(&PretendLikeIDontCare)
	a.AddThought(&OrImBusyWithOtherThings)
	a.AddThought(&CantFailAtSomethingThatNeverMattered)
	a.AddThought(&LikeISaid)
	a.AddThought(&AintACakeWalkOverHere)

	a.Connect(&ItAintACakeWalkOverHere, &IConsiderMyselfAHighFunctioningAnxiousPerson)
	a.Connect(&InTheSenseThatIWouldConsiderMyselfHighlyCapable, &AsFarAs)
	a.Connect(&AsFarAs, &Work)
	a.Connect(&AsFarAs, &PhysicalFitness)
	a.Connect(&AsFarAs, &Smart)
	a.Connect(&AsFarAs, &PopularLol)
	a.Connect(&AsFarAs, &GoodAtALotOFThings)
	a.Connect(&IConsiderMyselfAHighFunctioningAnxiousPerson, &ButActuallyImAnxiousALot)
	a.Connect(&ButActuallyImAnxiousALot, &ItsAMajorMotivatorForMe)
	a.Connect(&ItsAMajorMotivatorForMe, &Anxiety)
	a.Connect(&Anxiety, &TheFearOfFailing)
	a.Connect(&Anxiety, &TheFearOfMissingOut)
	a.Connect(&Anxiety, &TheFearOfNotBeingEnough)
	a.Connect(&ItsAMajorMotivatorForMe, &ItsAllRelated)
	a.Connect(&ButActuallyImAnxiousALot, &IWantEverythingIDoToBeDeliberate)
	a.Connect(&IWantEverythingIDoToBeDeliberate, &AndThoughtThrough)
	a.Connect(&IWantEverythingIDoToBeDeliberate, &AndPlanned)
	a.Connect(&IWantEverythingIDoToBeDeliberate, &ItMakesMeWorkHarder)
	a.Connect(&IWantEverythingIDoToBeDeliberate, &ItMakesMePushMyself)
	a.Connect(&IWantEverythingIDoToBeDeliberate, &ItMakesMeSuccessfulByLotsOfStandards)
	a.Connect(&IConsiderMyselfAHighFunctioningAnxiousPerson, &ButItAlsoMakesItReallyDifficultToAcceptFailure)
	a.Connect(&ButItAlsoMakesItReallyDifficultToAcceptFailure, &IPutTimeIntoEverythingIDo)
	a.Connect(&ButItAlsoMakesItReallyDifficultToAcceptFailure, &AndWhenTheEffortIPutDoesNotMatchTheOutcome)
	a.Connect(&AndWhenTheEffortIPutDoesNotMatchTheOutcome, &WhetherThats)
	a.Connect(&WhetherThats, &Relationships)
	a.Connect(&WhetherThats, &MyCareer)
	a.Connect(&WhetherThats, &MyFriendships)
	a.Connect(&WhetherThats, &EvenSomethingAsSimpleAsWeekendPlans)
	a.Connect(&WhetherThats, &OrASimpleText)
	a.Connect(&AndWhenTheEffortIPutDoesNotMatchTheOutcome, &ITakeItReallyHard)
	a.Connect(&AndWhenTheEffortIPutDoesNotMatchTheOutcome, &IShutDown)
	a.Connect(&AndWhenTheEffortIPutDoesNotMatchTheOutcome, &PretendLikeIDontCare)
	a.Connect(&AndWhenTheEffortIPutDoesNotMatchTheOutcome, &OrImBusyWithOtherThings)
	a.Connect(&ButItAlsoMakesItReallyDifficultToAcceptFailure, &CantFailAtSomethingThatNeverMattered)
	a.Connect(&CantFailAtSomethingThatNeverMattered, &LikeISaid)
	a.Connect(&LikeISaid, &AintACakeWalkOverHere)

	a.Connect(&AintACakeWalkOverHere, &ItAintACakeWalkOverHere)
}
