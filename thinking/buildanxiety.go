package thinking

import "github.com/ppelayo/anxiety/graph"

func BuildAnxiety(a *graph.Anxiety) {

	// Is my work good enough?
	a.AddThought(&IsMyWorkGoodEnough)
	a.AddThought(&WillAnyoneCare)
	a.AddThought(&DoesItMatter)
	a.AddThought(&WhatIfTheyDoCare)
	a.AddThought(&WhatIfTheyDont)

	a.Connect(&IsMyWorkGoodEnough, &WillAnyoneCare)
	a.Connect(&WillAnyoneCare, &DoesItMatter)
	a.Connect(&DoesItMatter, &WhatIfTheyDoCare)
	a.Connect(&DoesItMatter, &WhatIfTheyDont)

	// to be honest I'm kinda into you
	a.AddThought(&ToBeHonestImKindaIntoYou)
	a.AddThought(&LikeYoureReallCool)
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

	a.Connect(&ToBeHonestImKindaIntoYou, &LikeYoureReallCool)
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

	// loops
	a.Connect(&IGuessTheresALotOfReasonsLol, &ButItsKindaFuckedUpNo)
	a.Connect(&WhatIfTheyDont, &IsMyWorkGoodEnough)

}

// func Loop(a *graph.Anxiety) {

// 	// a.SetLoop()
// 	a.RandomConnect()
// }
