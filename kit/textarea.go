package kit

import (
	"sync"
	"github.com/satnamram/flexkit/dom"
)

type Textarea struct {
	mutex sync.Mutex
	ref *textareaRef

	state  TextareaState
	resize TextareaResize
	hidden bool
	icon   IconType
	onInput func()
}

type textareaRef struct {
	textarea *dom.Element
}

type TextareaState string

const (
	TextareaRegular  TextareaState = ""
	TextareaDanger   TextareaState = "uk-form-danger"
	TextareaSuccess  TextareaState = "uk-form-success"
	TextareaDisabled TextareaState = "disabled"
)

type TextareaResize string

const (
	TextAreaResizeVertical TextareaResize = "vertical"
	TextAreaResizeHorizontal TextareaResize = "horizontal"
	TextAreaResizeBoth TextareaResize = "both"
	TextAreaResizeNone TextareaResize = "none"
)

func NewTextArea() *Textarea {
	return &Textarea{
		state: TextareaRegular,
		resize:TextAreaResizeNone,
		icon:  IconNone,
	}
}

func (t *Textarea) String() string {
	if t.ref != nil {
		return t.ref.textarea.Get("value").String()
	}
	return ""
}

func (t *Textarea) Resize(v TextareaResize) *Textarea {
	t.mutex.Lock()
	t.resize = v
	t.renderResize()
	t.mutex.Unlock()
	return t
}

func (t *Textarea) renderResize() {
	if t.ref != nil {
		t.ref.textarea.Get("style").Set("resize", string(t.resize))
	}
}

func (t *Textarea) Hidden(v bool) *Textarea {
	t.mutex.Lock()
	t.hidden = v
	t.renderHidden()
	t.mutex.Unlock()
	return t
}

func (t *Textarea) renderHidden() {
	if t.ref != nil {
		if t.hidden {
			t.ref.textarea.Set("type", "password")
		} else {
			t.ref.textarea.Set("type", "text")
		}
	}
}

func (t *Textarea) State(state TextareaState) *Textarea {
	t.mutex.Lock()
	t.cleanupState()
	t.state = state
	t.renderState()
	t.mutex.Unlock()
	return t
}


func (t *Textarea) cleanupState() {
	if t.ref != nil {
		if t.state == TextareaDisabled {
			t.ref.textarea.Set("disabled", "false")
		} else if t.state != TextareaRegular {
			t.ref.textarea.RemoveClass(string(t.state))
		}
	}
}

func (t *Textarea) renderState() {
	if t.ref != nil {
		if t.state == TextareaDisabled {
			t.ref.textarea.Set("disabled", "true")
		} else if t.state != TextareaRegular {
			t.ref.textarea.AddClass(string(t.state))
		}
	}
}

func(t *Textarea) OnInput(f func()) *Textarea {
	t.mutex.Lock()
	t.onInput = f
	if t.ref != nil && t.onInput != nil {
		t.ref.textarea.OnInput(t.onInput)
	}
	t.mutex.Unlock()
	return t
}

func (t *Textarea) Render() *dom.Element {
	t.mutex.Lock()

	t.ref = &textareaRef{
		textarea: dom.NewElement("textarea"),
	}

	t.ref.textarea.AddClass("uk-textarea")

	t.renderState()
	t.renderHidden()
	t.renderResize()
	if t.ref != nil && t.onInput != nil {
		t.ref.textarea.OnInput(t.onInput)
	}


	t.mutex.Unlock()
	return t.ref.textarea
}
