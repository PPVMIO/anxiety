package journal

import (
	"github.com/ppelayo/anxiety/graph"
)

var IGoOutALot = graph.Thought{Value: "IGoOutALot", First: true}
var ILikeGettingFuckedUp = graph.Thought{Value: "ILikeGettingFuckedUp"}
var WhyDoIDoItSoMuch = graph.Thought{Value: "WhyDoIDoItSoMuch"}
var IOnlyReallyDoPartyDrugsNow = graph.Thought{Value: "IOnlyReallyDoPartyDrugsNow"}
var ICantDoAnyOfThatPsychedelicShit = graph.Thought{Value: "ICantDoAnyOfThatPsychedelicShit"}
var ImTooHighStrug = graph.Thought{Value: "ImTooHighStrug"}
var IGuessTheresYourAnswer = graph.Thought{Value: "IGuessTheresYourAnswer"}
var ThatsWhyIGoOut = graph.Thought{Value: "ThatsWhyIGoOut"}
var ISpendTheEntireWorkWeek = graph.Thought{Value: "ISpendTheEntireWorkWeek"}
var FunctioningAtThisSuperHighCapacity = graph.Thought{Value: "FunctioningAtThisSuperHighCapacity"}
var DissectingProblems = graph.Thought{Value: "DissectingProblems"}
var WorkingOut = graph.Thought{Value: "WorkingOut"}
var WorkingOnProjects = graph.Thought{Value: "WorkingOnProjects"}
var TryingToPersonallyImprove = graph.Thought{Value: "TryingToPersonallyImprove"}
var TheAmountOfMentalPressureIPutOnMyselfIsALot = graph.Thought{Value: "TheAmountOfMentalPressureIPutOnMyselfIsALot"}
var IhavePerfectionistTendencies = graph.Thought{Value: "IhavePerfectionistTendencies"}
var AndINeedABreakFromAllOfThat = graph.Thought{Value: "AndINeedABreakFromAllOfThat"}
var SoIDrinkAndDoDrugs = graph.Thought{Value: "SoIDrinkAndDoDrugs"}
var ButHonestly = graph.Thought{Value: "ButHonestly"}
var ItsNotQuiteEnoughAnymore = graph.Thought{Value: "ItsNotQuiteEnoughAnymore"}
var IDontKnowIfItsThePandemic = graph.Thought{Value: "IDontKnowIfItsThePandemic"}
var CuzGoingOutNowIsNotWhatItWasPrepandemic = graph.Thought{Value: "CuzGoingOutNowIsNotWhatItWasPrepandemic"}
var ButEmotionallyImStartingToThinkItDrainsMeJustAsMuchAsWork = graph.Thought{Value: "ButEmotionallyImStartingToThinkItDrainsMeJustAsMuchAsWork"}
var ItsNotThatIDontLovePartying = graph.Thought{Value: "ItsNotThatIDontLovePartying"}
var IStillThinkThereIsLegitamiteValueToGoingOut = graph.Thought{Value: "IStillThinkThereIsLegitamiteValueToGoingOut"}
var ButIReallyNeedToReexamineMyRelationshipWithTheWholeConcept = graph.Thought{Value: "ButIReallyNeedToReexamineMyRelationshipWithTheWholeConcept"}
var TheresNothingWorse = graph.Thought{Value: "TheresNothingWorse"}
var ThanGoingToBedFuckedUpAndAloneLol = graph.Thought{Value: "ThanGoingToBedFuckedUpAndAloneLol"}
var AndMentallyInABadPlace = graph.Thought{Value: "AndMentallyInABadPlace"}
var AndIFeelLikeIEndUpThereTooOften = graph.Thought{Value: "AndIFeelLikeIEndUpThereTooOften"}
var ItGoesAwayWhenIWakeup = graph.Thought{Value: "ItGoesAwayWhenIWakeup"}
var ButIllProbablyGoOutTonight = graph.Thought{Value: "ButIllProbablyGoOutTonight"}
var BeerIsMyFavorite = graph.Thought{Value: "BeerIsMyFavorite"}

func iGoOutALot(a *graph.Anxiety) {

	a.AddThought(&IGoOutALot)
	a.AddThought(&ILikeGettingFuckedUp)
	a.AddThought(&WhyDoIDoItSoMuch)
	a.AddThought(&IOnlyReallyDoPartyDrugsNow)
	a.AddThought(&ICantDoAnyOfThatPsychedelicShit)
	a.AddThought(&ImTooHighStrug)
	a.AddThought(&IGuessTheresYourAnswer)
	a.AddThought(&ThatsWhyIGoOut)
	a.AddThought(&ISpendTheEntireWorkWeek)
	a.AddThought(&FunctioningAtThisSuperHighCapacity)
	a.AddThought(&DissectingProblems)
	a.AddThought(&WorkingOut)
	a.AddThought(&WorkingOnProjects)
	a.AddThought(&TryingToPersonallyImprove)
	a.AddThought(&TheAmountOfMentalPressureIPutOnMyselfIsALot)
	a.AddThought(&IhavePerfectionistTendencies)
	a.AddThought(&AndINeedABreakFromAllOfThat)
	a.AddThought(&SoIDrinkAndDoDrugs)
	a.AddThought(&ButHonestly)
	a.AddThought(&ItsNotQuiteEnoughAnymore)
	a.AddThought(&IDontKnowIfItsThePandemic)
	a.AddThought(&CuzGoingOutNowIsNotWhatItWasPrepandemic)
	a.AddThought(&ButEmotionallyImStartingToThinkItDrainsMeJustAsMuchAsWork)
	a.AddThought(&ItsNotThatIDontLovePartying)
	a.AddThought(&IStillThinkThereIsLegitamiteValueToGoingOut)
	a.AddThought(&ButIReallyNeedToReexamineMyRelationshipWithTheWholeConcept)
	a.AddThought(&ButIllProbablyGoOutTonight)
	a.AddThought(&BeerIsMyFavorite)

	a.Connect(&IGoOutALot, &ILikeGettingFuckedUp)
	a.Connect(&ILikeGettingFuckedUp, &WhyDoIDoItSoMuch)
	a.Connect(&WhyDoIDoItSoMuch, &ImTooHighStrug)
	a.Connect(&ImTooHighStrug, &IGuessTheresYourAnswer)
	a.Connect(&IGuessTheresYourAnswer, &ThatsWhyIGoOut)
	a.Connect(&ThatsWhyIGoOut, &ISpendTheEntireWorkWeek)
	a.Connect(&ISpendTheEntireWorkWeek, &FunctioningAtThisSuperHighCapacity)
	a.Connect(&ISpendTheEntireWorkWeek, &DissectingProblems)
	a.Connect(&ISpendTheEntireWorkWeek, &WorkingOut)
	a.Connect(&ISpendTheEntireWorkWeek, &WorkingOnProjects)
	a.Connect(&ISpendTheEntireWorkWeek, &TryingToPersonallyImprove)
	a.Connect(&ISpendTheEntireWorkWeek, &TheAmountOfMentalPressureIPutOnMyselfIsALot)
	a.Connect(&ISpendTheEntireWorkWeek, &IhavePerfectionistTendencies)
	a.Connect(&ISpendTheEntireWorkWeek, &AndINeedABreakFromAllOfThat)
	a.Connect(&AndINeedABreakFromAllOfThat, &SoIDrinkAndDoDrugs)
	a.Connect(&SoIDrinkAndDoDrugs, &IOnlyReallyDoPartyDrugsNow)
	a.Connect(&SoIDrinkAndDoDrugs, &ICantDoAnyOfThatPsychedelicShit)
	a.Connect(&SoIDrinkAndDoDrugs, &ICantDoAnyOfThatPsychedelicShit)
	a.Connect(&SoIDrinkAndDoDrugs, &ButHonestly)
	a.Connect(&ButHonestly, &ItsNotQuiteEnoughAnymore)
	a.Connect(&ItsNotQuiteEnoughAnymore, &IDontKnowIfItsThePandemic)
	a.Connect(&CuzGoingOutNowIsNotWhatItWasPrepandemic, &ButEmotionallyImStartingToThinkItDrainsMeJustAsMuchAsWork)
	a.Connect(&ItsNotThatIDontLovePartying, &IStillThinkThereIsLegitamiteValueToGoingOut)
	a.Connect(&ItsNotThatIDontLovePartying, &ButIReallyNeedToReexamineMyRelationshipWithTheWholeConcept)
	a.Connect(&ItsNotThatIDontLovePartying, &ButIReallyNeedToReexamineMyRelationshipWithTheWholeConcept)
	a.Connect(&ButIReallyNeedToReexamineMyRelationshipWithTheWholeConcept, &ButIllProbablyGoOutTonight)

	a.Connect(&ButIllProbablyGoOutTonight, &IGoOutALot)

}
