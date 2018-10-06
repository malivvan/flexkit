package kit

import (
	"github.com/gopherjs/gopherjs/js"
	"sync"
	"github.com/satnamram/flexkit/dom"
)

type TableStyle string


var (
	TableStylePlain TableStyle= ""
	TableStyleDivider TableStyle= "uk-table-divider"
	TableStyleStriped TableStyle= "uk-table-striped"
)


type Table struct {
	mutex sync.Mutex
	ref   *tableRef

	style TableStyle
	caption   interface{}
	header    []interface{}
	widths    []string
	footer    []string
	sorttable bool

	items [][]interface{}
}

type tableRef struct {
	wrapper *dom.Element
	table   *dom.Element
	caption *dom.Element
	header  *dom.Element
	body    *dom.Element
	footer  *dom.Element
}

func NewTable() *Table {
	return &Table{
		caption: "",
		items: [][]interface{}{},
	}
}

func (t *Table) Append(v ...interface{}) *Table {
	t.mutex.Lock()
	t.items = append(t.items, v)
	t.renderBody()
	t.mutex.Unlock()
	return t
}

func (t *Table) Caption(v interface{}) *Table {
	t.mutex.Lock()
	t.caption = v
	t.renderCaption()
	t.mutex.Unlock()
	return t
}


func (t *Table) renderCaption() *Table {
	if t.ref != nil {
		t.ref.caption.SetContent(t.caption)
	}
	return t
}

func (t *Table) Width(widths ...string) *Table {
	t.mutex.Lock()
	t.widths = widths
	t.renderHeader()
	t.mutex.Unlock()
	return t
}

func (t *Table) Header(fields ...interface{}) *Table {
	t.mutex.Lock()
	t.header = fields
	t.renderHeader()
	t.mutex.Unlock()
	return t
}


func (t *Table) renderHeader() {
	if t.ref != nil {
		t.ref.header.Set("innerHTML", "")
		tr := dom.NewElement("tr")
		for i, field := range t.header {
			th := dom.NewElement("th").SetContent(field)

			// if there is a static width set it
			if i < len(t.widths) {
				if len(t.widths[i]) > 0 {
					th.AddClass("uk-width-" + t.widths[i])
				}

			}

			tr.Append(th)
		}
		t.ref.header.Set("innerHTML", "")
		t.ref.header.Append(tr)
	}
}

func (t *Table) Style(style TableStyle) *Table {
	t.mutex.Lock()
	t.cleanupStyle()
	t.style = style
	t.renderStyle()
	t.mutex.Unlock()
	return t
}


func (t *Table) cleanupStyle() {
	if t.ref != nil && t.style != TableStylePlain {
		t.ref.table.RemoveClass(string(t.style))
	}
}

func (t *Table) renderStyle() {
	if t.ref != nil && t.style != TableStylePlain {
		t.ref.table.AddClass(string(t.style))
	}
}

func (t *Table) renderBody() {
	if t.ref != nil {
		t.ref.body.Set("innerHTML", "")
		for _, item := range t.items {
			tr := dom.NewElement("tr")
			for _, field := range item {
				tr.Append(dom.NewElement("td").SetContent(field))
			}
			t.ref.body.Append(tr)
		}
	}
}

func (t *Table) Render() *dom.Element {
	t.mutex.Lock()

	t.ref = &tableRef{
		wrapper: dom.NewElement("div"),
		table:   dom.NewElement("table"),
		caption: dom.NewElement("caption"),
		header:  dom.NewElement("thead"),
		body:    dom.NewElement("tbody"),
		footer:  dom.NewElement("tfoot"),
	}

	t.ref.wrapper.Append(t.ref.table)
	t.ref.table.Append(t.ref.caption)
	t.ref.table.Append(t.ref.header)
	t.ref.table.Append(t.ref.body)

	t.ref.wrapper.AddClass("uk-overflow-auto")
	t.ref.table.AddClass("uk-table")

	t.renderStyle()
	t.renderCaption()
	t.renderHeader()
	t.renderBody()

	// make it sortable using sorttable.js
	js.Global.Get("sorttable").Call("makeSortable", *t.ref.table.Value)

	t.mutex.Unlock()
	return t.ref.wrapper
}
