package kit

import (
	"sync"
	"github.com/satnamram/flexkit/dom"
)

type HTML struct {
	mutex sync.Mutex
	ref   *htmlRef

	content string
}

type htmlRef struct {
	html *dom.Element
}

func NewHTML() *HTML {
	return &HTML{}
}

func (html *HTML) Set(content string) *HTML {
	html.mutex.Lock()
	html.content = content
	html.renderHTML()
	html.mutex.Unlock()
	return html
}

func (html *HTML) renderHTML() {
	if html.ref != nil {
		html.ref.html.Set("innerHTML", html.content)
	}
}

func (html *HTML) Render() *dom.Element {
	html.mutex.Lock()
	html.ref = &htmlRef{
		html: dom.NewElement("div"),
	}
	html.renderHTML()
	html.mutex.Unlock()
	return html.ref.html
}
