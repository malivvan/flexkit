package kit

import (
	"sync"
	"github.com/satnamram/flexkit/dom"
)

type Icon struct {
	mutex sync.Mutex
	ref   *iconRef

	t IconType
	onClick   func()
}

type iconRef struct {
	icon *dom.Element
}

func NewIcon() *Icon {
	return &Icon{
		t: IconNone,
	}
}

func (i *Icon) OnClick(f func()) *Icon {
	i.mutex.Lock()
	i.onClick = f
	i.renderOnClick()
	i.mutex.Unlock()
	return i
}

func (i *Icon) renderOnClick() {
	if i.ref != nil {
		i.ref.icon.OnClick(i.onClick)
		if i.onClick == nil {
			i.ref.icon.RemoveClass("uk-icon-button")
		} else {
			i.ref.icon.AddClass("uk-icon-button")
		}
	}
}

func (i *Icon) Type(t IconType) *Icon {
	i.mutex.Lock()
	i.t = t
	i.renderType()
	i.mutex.Unlock()
	return i
}

func (i *Icon) renderType() {
	if i.ref != nil {
		i.ref.icon.SetAttribute("uk-icon", "icon: "+string(i.t))
	}
}

func (i *Icon) Render() *dom.Element {
	i.mutex.Lock()

	i.ref = &iconRef{
		icon: dom.NewElement("span"),
	}

	i.renderType()
	i.renderOnClick()

	i.mutex.Unlock()
	return i.ref.icon
}
