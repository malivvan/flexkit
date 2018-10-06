package flex

import (
	"sync"
	"math/rand"
)



var (
	idRunes       = []rune("abcdefghijklmnopqrstuvwxyz")
	idLength      = 8
	idRetries     = 100
	generatedIDs  = []string{}
	generatedIDs_ sync.Mutex
)

func generateUniqueID() string {
	generatedIDs_.Lock()
	for i := 0; i < idRetries; i++ {
		id := generateID()
		idIsUnique := true
		for _, existingID := range generatedIDs {
			if existingID == id {
				idIsUnique = false
				break
			}
		}
		if idIsUnique {
			generatedIDs_.Unlock()
			return id
		}
	}
	generatedIDs_.Unlock()
	panic("error generating unique id")
}

func freeID(id string) {
	generatedIDs_.Lock()
	for i, existingID := range generatedIDs {
		if existingID == id {
			generatedIDs = append(generatedIDs[:i], generatedIDs[i+1:]...)
			break
		}
	}
	generatedIDs_.Unlock()
}

func generateID() string {
	b := make([]rune, idLength)
	for i := range b {
		b[i] = idRunes[rand.Intn(len(idRunes))]
	}
	return string(b)
}

type ScreenSize string

var ScreenSizes = []ScreenSize{
	ScreenSizeXSmall,
	ScreenSizeSmall,
	ScreenSizeMedium,
	ScreenSizeLarge,
	ScreenSizeXLarge,
}

const (
	ScreenSizeXSmall ScreenSize = ""
	ScreenSizeSmall  ScreenSize = "@media all and (min-width: 600px)"
	ScreenSizeMedium ScreenSize = "@media all and (min-width: 768px)"
	ScreenSizeLarge  ScreenSize = "@media all and (min-width: 992px)"
	ScreenSizeXLarge ScreenSize = "@media all and (min-width: 1200px)"
)

type FlexWrap string

const (
	Wrap        FlexWrap = "flex-wrap:wrap;"
	NoWrap      FlexWrap = "flex-wrap:nowrap;"
	WrapReverse FlexWrap = "flex-wrap:wrap-reverse;"
)

type FlexDirection string

const (
	Row           FlexDirection = "flex-direction:row;"
	RowReverse    FlexDirection = "flex-direction:row-reverse;"
	Column        FlexDirection = "flex-direction:column;"
	ColumnReverse FlexDirection = "flex-direction:column-reverse;"
)

type FlexAlignment string

const (
	Auto         FlexAlignment = "auto;"
	Start        FlexAlignment = "flex-start;"
	End          FlexAlignment = "flex-end;"
	Center       FlexAlignment = "center;"
	SpaceBetween FlexAlignment = "space-between;"
	SpaceAround  FlexAlignment = "space-around;"
	SpaceEvenly  FlexAlignment = "space-evenly;"
	Baseline     FlexAlignment = "baseline;"
	Stretch      FlexAlignment = "stretch;"
)

var flexAlignmentJustifyContent = "justify-content:"
var flexAlignmentJustifyContentInputs = []FlexAlignment{Start, End, Center, SpaceBetween, SpaceAround, SpaceEvenly}

func (fa FlexAlignment) justifyContent() string {
	for _, alignment := range flexAlignmentJustifyContentInputs {
		if fa == alignment {
			return flexAlignmentJustifyContent + string(alignment)
		}
	}
	return ""
}

var flexAlignmentAlignItems = "align-items:"
var flexAlignmentAlignItemsInputs = []FlexAlignment{Stretch, Start, End, Center, Baseline}

func (fa FlexAlignment) alignItems() string {
	for _, alignment := range flexAlignmentAlignItemsInputs {
		if fa == alignment {
			return flexAlignmentAlignItems + string(alignment)
		}
	}
	return ""
}

var flexAlignmentAlignContent = "align-content:"
var flexAlignmentAlignContentInputs = []FlexAlignment{Stretch, Start, End, Center, SpaceBetween, SpaceAround}

func (fa FlexAlignment) alignContent() string {
	for _, alignment := range flexAlignmentAlignContentInputs {
		if fa == alignment {
			return flexAlignmentAlignContent + string(alignment)
		}
	}
	return ""
}

var flexAlignmentAlignSelf = "align-self:"
var flexAlignmentAlignSelfInputs = []FlexAlignment{Auto, Start, End, Center, Baseline, Stretch}

func (fa FlexAlignment) alignSelf() string {
	for _, alignment := range flexAlignmentAlignSelfInputs {
		return flexAlignmentAlignSelf + string(alignment)
	}
	return ""
}


