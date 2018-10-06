package dom

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
)

type Element struct {
	Value   *js.Object
	classes *js.Object

	onClickCallback func()
	onInputCallback func()
}

func NewElement(t string) *Element {
	value := DOC.Call("createElement", t)
	classes := value.Get("classList")
	return &Element{
		Value:   value,
		classes: classes,
	}
}

func getElement(t string) *Element {
	documentHead := DOC.Get(t)
	documentHeadClasses := documentHead.Get("classList")
	return &Element{
		Value:   documentHead,
		classes: documentHeadClasses,
	}
}

func (element *Element) Set(p string, x interface{}) *Element {
	if p == "type" {
		js.Global.Call("makePropertyWriteable", *element.Value, "type")
	}
	element.Value.Set(p, x)
	return element
}

func (element *Element) Get(p string) *js.Object {
	return element.Value.Get(p)
}

func (element *Element) Prepend(e *Element) *Element {
	element.Value.Call("prepend", e.Value)
	return element
}

func (element *Element) Append(e *Element) *Element {
	element.Value.Call("append", *e.Value)
	return element
}

func (element *Element) AddClass(class string) *Element {
	element.classes.Call("add", class)
	return element
}

func (element *Element) RemoveClass(class string) *Element {
	element.classes.Call("remove", class)
	return element
}

func (element *Element) ToggleClass(class string) *Element {
	element.classes.Call("toggle", class)
	return element
}

func (element *Element) SetAttribute(attr string, v string) *Element {
	element.Value.Call("setAttribute", attr, v)
	return element
}

func (element *Element) Wrap(wrapper *Element, isInitialized bool) *Element {
	if !isInitialized {
		wrapper.Append(element)
		return element
	}
	parentNode := element.Value.Get("parentNode")
	parentNode.Call("insertBefore", *wrapper.Value, *element.Value)
	wrapper.Value.Call("appendChild", *element.Value)
	return element
}

func (element *Element) Unwrap(wrapper *Element) *Element {
	parentNode := element.Value.Get("parentNode")
	parentNode.Call("insertBefore", *element.Value, *wrapper.Value)
	parentNode.Call("removeChild", *wrapper)
	return element
}

func (element *Element) SetContent(v interface{}) *Element {
	switch v := v.(type) {
	case nil:
		element.Set("innerHTML", "nil")
	case string:
		element.Set("innerHTML", v)
	case int:
		element.Set("innerHTML", strconv.Itoa(v))
	default:
		if e, ok := v.(Renderable); ok {
			element.Append(e.Render())
			break
		}
		panic("type unknown")
	}
	return element
}

// TODO: callback return Value
func (element *Element) OnInput(f func()) {
	if element.onInputCallback != nil {
		element.Value.Call("removeEventListener", "input", element.onInputCallback)
		element.onInputCallback = nil
	}
	if f != nil {
		element.onInputCallback = f
		element.Value.Call("addEventListener", "input", element.onInputCallback)
	}
}

// TODO: callback return Value
func (element *Element) OnClick(f func()) {
	if element.onClickCallback != nil {
		element.Value.Call("removeEventListener", "click", element.onClickCallback)
		element.onClickCallback = nil
	}
	if f != nil {
		element.onClickCallback = f
		element.Value.Call("addEventListener", "click", element.onClickCallback)
	}
}
