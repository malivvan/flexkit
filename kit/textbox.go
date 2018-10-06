package kit

import (
	"sync"
	"github.com/satnamram/flexkit/dom"
)

type Textbox struct {
	mutex sync.Mutex
	ref *textboxRef

	state  TextboxState
	hidden bool
	icon   IconType
}

type textboxRef struct {
	textarea *dom.Element
	wrapper  *dom.Element
}

type TextboxState string

const (
	TextboxRegular  TextboxState = ""
	TextboxDanger   TextboxState = "uk-form-danger"
	TextboxSuccess  TextboxState = "uk-form-success"
	TextboxDisabled TextboxState = "disabled"
)

func NewTextBox() *Textbox {
	return &Textbox{
		state: TextboxRegular,
		icon:  IconNone,
	}
}

func (t *Textbox) Hidden(v bool) *Textbox {
	t.mutex.Lock()
	t.hidden = v
	t.renderHidden()
	t.mutex.Unlock()
	return t
}

func (t *Textbox) renderHidden() {
	if t.ref != nil {
		if t.hidden {
			t.ref.textarea.Set("type", "password")
		} else {
			t.ref.textarea.Set("type", "text")
		}
	}
}

func (t *Textbox) Icon(icon IconType) *Textbox {
	t.mutex.Lock()
	t.cleanupIcon()
	t.icon = icon
	t.renderIcon(true)
	t.mutex.Unlock()
	return t
}

func (t *Textbox) cleanupIcon() {
	if t.icon != IconNone && t.ref != nil && t.ref.wrapper != nil {
		t.ref.textarea.Unwrap(t.ref.wrapper)
		t.ref.wrapper = nil
	}
}

func (t *Textbox) renderIcon(isInitialized bool) {
	if t.ref != nil && t.icon != IconNone {
		icon := dom.NewElement("span").
			AddClass("uk-form-icon").
			SetAttribute("uk-icon", "icon: "+string(t.icon))
		t.ref.wrapper = dom.NewElement("div").
			AddClass("uk-inline").
			Append(icon)

		// Note: Plain inputs expand to full width. This does not happen
		// when using the icon wrapper div. Therefore we force full width.
		t.ref.wrapper.Set("style", "width:100%;")

		t.ref.textarea.Wrap(t.ref.wrapper, isInitialized)
	}
}

func (t *Textbox) State(state TextboxState) *Textbox {
	t.mutex.Lock()
	t.cleanupState()
	t.state = state
	t.renderState()
	t.mutex.Unlock()
	return t
}

func (t *Textbox) cleanupState() {
	if t.ref != nil {
		if t.state == TextboxDisabled {
			t.ref.textarea.Set("disabled", "false")
		} else if t.state != TextboxRegular {
			t.ref.textarea.RemoveClass(string(t.state))
		}
	}
}

func (t *Textbox) renderState() {
	if t.ref != nil {
		if t.state == TextboxDisabled {
			t.ref.textarea.Set("disabled", "true")
		} else if t.state != TextboxRegular {
			t.ref.textarea.AddClass(string(t.state))
		}
	}
}

func (t *Textbox) Render() *dom.Element {
	t.mutex.Lock()

	t.ref = &textboxRef{
		textarea: dom.NewElement("input"),
		// wrapper is only added if icon is present
	}

	t.ref.textarea.Set("type", "text")
	t.ref.textarea.AddClass("uk-input")

	t.renderState()
	t.renderHidden()
	t.renderIcon(false)

	t.mutex.Unlock()
	if t.ref.wrapper != nil {
		return t.ref.wrapper
	}
	return t.ref.textarea
}
