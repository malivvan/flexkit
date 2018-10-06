package kit

import (
	"sync"
	"github.com/satnamram/flexkit/dom"
)

type Button struct {
	mutex sync.Mutex
	ref   *buttonRef

	label     interface{}
	style     ButtonStyle
	size      ButtonSize
	fillWidth bool
	onClick   func()
}

type buttonRef struct {
	button *dom.Element
}

func NewButton() *Button {
	return &Button{
		style: ButtonDefault,
		size:  ButtonMedium,
	}
}

func (b *Button) Label(v interface{}) *Button {
	b.label = v
	b.renderLabel()
	return b
}

func (b *Button) renderLabel() {
	if b.ref != nil {
		b.ref.button.SetContent(b.label)
	}
}

type ButtonStyle string

const (
	ButtonDefault   ButtonStyle = "uk-button-default"
	ButtonPrimary   ButtonStyle = "uk-button-primary"
	ButtonSecondary ButtonStyle = "uk-button-secondary"
	ButtonDanger    ButtonStyle = "uk-button-danger"
	ButtonText      ButtonStyle = "uk-button-text"
	ButtonLink      ButtonStyle = "uk-button-link"
)

type ButtonSize string

const (
	ButtonSmall  ButtonSize = "uk-button-small"
	ButtonMedium ButtonSize = ""
	ButtonLarge  ButtonSize = "uk-button-large"
)

func (b *Button) Style(style ButtonStyle) *Button {
	b.mutex.Lock()
	b.cleanupStyle()
	b.style = style
	b.renderStyle()
	b.mutex.Unlock()
	return b
}

func (b *Button) cleanupStyle() {
	if b.ref != nil {
		b.ref.button.RemoveClass(string(b.style))
	}
}

func (b *Button) renderStyle() {
	if b.ref != nil {
		b.ref.button.AddClass(string(b.style))
	}
}

func (b *Button) Size(size ButtonSize) *Button {
	b.mutex.Lock()
	b.cleanupSize()
	b.size = size
	b.renderSize()
	b.mutex.Unlock()
	return b
}

func (b *Button) cleanupSize() {
	if b.ref != nil && b.size != ButtonMedium {
		b.ref.button.RemoveClass(string(b.size))
	}
}

func (b *Button) renderSize() {
	if b.ref != nil && b.size != ButtonMedium {
		b.ref.button.AddClass(string(b.size))
	}
}

const buttonFillWidth = "uk-width-1-1"

func (b *Button) FillWidth(v bool) *Button {
	b.mutex.Lock()
	b.fillWidth = v
	b.renderFillWidth()
	b.mutex.Unlock()
	return b
}

func (b *Button) renderFillWidth() {
	if b.ref != nil {
		if b.fillWidth {
			b.ref.button.AddClass(buttonFillWidth)
		} else {
			b.ref.button.RemoveClass(buttonFillWidth)
		}
	}
}

func (b *Button) OnClick(f func()) *Button {
	b.mutex.Lock()
	b.onClick = f
	if b.ref != nil {
		b.ref.button.OnClick(b.onClick)
	}
	b.mutex.Unlock()
	return b
}

func (b *Button) Render() *dom.Element {
	b.mutex.Lock()

	b.ref = &buttonRef{
		button: dom.NewElement("button"),
	}

	b.ref.button.AddClass("uk-button")

	b.renderLabel()
	b.renderStyle()
	b.renderSize()
	b.renderFillWidth()
	if b.ref != nil {
		b.ref.button.OnClick(b.onClick)
	}

	b.mutex.Unlock()
	return b.ref.button
}
