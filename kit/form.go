package kit

import (
	"sync"
	"github.com/gopherjs/gopherjs/js"
	"github.com/satnamram/flexkit/dom"
)

type Form struct {
	mutex sync.Mutex
	ref   *formRef

	margins []*formMargin

	formNode *js.Object
}

type formRef struct {
	form *dom.Element
}

type formMargin struct {
	label   string
	renderable dom.Renderable
}

func NewForm() *Form {
	return &Form{
		margins: []*formMargin{},
	}
}

func (f *Form) AddTextbox(label string, t *Textbox) *Form {
	f.mutex.Lock()

	f.margins = append(f.margins, &formMargin{
		label:   label,
		renderable: t,
	})

	f.mutex.Unlock()
	return f
}

func (f *Form) Render() *dom.Element {
	f.mutex.Lock()

	f.ref = &formRef{
		form: dom.NewElement("form"),
	}

	for _, margin := range f.margins {

		node := dom.NewElement("div").AddClass("uk-margin")
		if margin.label != "" {
			node.Append(dom.NewElement("div").Set("innerHTML", margin.label).AddClass("uk-form-label"))
		}
		node.Append(margin.renderable.Render())
		f.ref.form.Append(node)
	}

	f.mutex.Unlock()
	return f.ref.form
}
