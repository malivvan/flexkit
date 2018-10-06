package flex

import (
	"strconv"
	"sync"
	"bytes"
	"github.com/satnamram/flexkit/dom"
)

type Item struct {
	id        string
	container *Container

	attributes   *itemAttributes
	mediaQueries map[ScreenSize]*itemAttributes

	// inner content
	expandWidth  bool
	expandHeight bool
	renderable   dom.Renderable

	// hide/show behavior (usable after intialization)
	mutex  sync.Mutex
	hidden bool
	ref    *dom.Element
}

func NewItem(renderable dom.Renderable) *Item {
	return &Item{
		id:           generateUniqueID(),
		attributes:   newDefaultItemAttributes(),
		mediaQueries: make(map[ScreenSize]*itemAttributes),
		renderable:   renderable,
	}
}

type itemAttributes struct {
	border     string
	padding    string
	margin     string
	background string
	order      int
	grow       int
	shrink     int
	basis      string
	align      FlexAlignment
}

func newDefaultItemAttributes() *itemAttributes {
	return &itemAttributes{
		grow:   0,
		shrink: 0,
		basis:  "auto",
		align:  Auto,
	}
}

func (ia *itemAttributes) getBackground() string {
	if ia.background == "" {
		return ""
	}
	return "background-color:" + ia.background + ";"
}

func (ia *itemAttributes) getBorder() string {
	if ia.border == "" {
		return ""
	}
	return "border:" + ia.border + ";"
}

func (ia *itemAttributes) getPadding() string {
	if ia.padding == "" {
		return ""
	}
	return "padding:" + ia.padding + ";"
}

func (ia *itemAttributes) getMargin() string {
	if ia.margin == "" {
		return ""
	}
	return "margin:" + ia.margin + ";"
}

func (ia *itemAttributes) getOrder() string {
	if ia.order == 0 {
		return ""
	}
	return "flex-order:" + strconv.Itoa(ia.order) + ";"
}

func (ia *itemAttributes) getGrow() string {
	if ia.grow == 0 {
		return ""
	}
	return "flex-grow:" + strconv.Itoa(ia.grow) + ";"
}

func (ia *itemAttributes) getShrink() string {
	if ia.shrink == 0 {
		return ""
	}
	return "flex-shrink:" + strconv.Itoa(ia.shrink) + ";"
}

func (ia *itemAttributes) getBasis() string {
	if ia.basis == "auto" {
		return ""
	}
	return "flex-basis:" + ia.basis + ";"
}

func (ia *itemAttributes) getAlignSelf() string {
	if ia.align == Auto {
		return ""
	}
	return ia.align.alignSelf()
}

func (i *Item) Background(v string, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.background = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].background = v
	}
	return i
}

func (i *Item) Border(v string, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.border = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].border = v
	}
	return i
}

func (i *Item) Padding(v string, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.padding = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].padding = v
	}
	return i
}

func (i *Item) Margin(v string, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.margin = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].margin = v
	}
	return i
}

func (i *Item) Order(v int, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.order = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].order = v
	}
	return i
}

func (i *Item) Grow(v int, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.grow = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].grow = v
	}
	return i
}

func (i *Item) Shrink(v int, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.shrink = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].shrink = v
	}
	return i
}

func (i *Item) Basis(v string, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.basis = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].basis = v
	}
	return i
}

func (i *Item) Align(v FlexAlignment, sizes ...ScreenSize) *Item {
	if len(sizes) == 0 {
		i.attributes.align = v
		return i
	}
	for _, size := range sizes {
		if i.mediaQueries[size] == nil {
			i.mediaQueries[size] = newDefaultItemAttributes()
		}
		i.mediaQueries[size].align = v
	}
	return i
}

// Expand item width and height.
func (i *Item) Expand() *Item {
	i.expandHeight = true
	i.expandWidth = true
	return i
}

// Expand item width.
func (i *Item) ExpandWidth() *Item {
	i.expandWidth = true
	return i
}

// Expand item height.
func (i *Item) ExpandHeight() *Item {
	i.expandHeight = true
	return i
}

// Toggle the visibility of the flex item (usable after initialization).
func (i *Item) Toggle() {
	i.mutex.Lock()
	if i.hidden {
		i.hidden = false
	} else {
		i.hidden = true
	}
	i.renderHidden()
	i.mutex.Unlock()
}

// Show the flex item (usable after initialization).
func (i *Item) Show() {
	i.mutex.Lock()
	i.hidden = false
	i.renderHidden()
	i.mutex.Unlock()
}

// Hide the flex item (usable after initilization).
func (i *Item) Hide() *Item {
	i.mutex.Lock()
	i.hidden = true
	i.renderHidden()
	i.mutex.Unlock()
	return i
}

func (i *Item) renderHidden() {
	if i.ref != nil {
		if i.hidden {
			i.ref.RemoveClass("hidden")
		} else {
			i.ref.AddClass("hidden")
		}
	}
}

func (i *Item) generateCSS() string {
	var buf bytes.Buffer
	buf.WriteString("#" + i.id + " {")
	buf.WriteString(i.attributes.getPadding())
	buf.WriteString(i.attributes.getMargin())
	buf.WriteString(i.attributes.getBorder())
	buf.WriteString(i.attributes.getBackground())
	buf.WriteString(i.attributes.getOrder())
	buf.WriteString(i.attributes.getGrow())
	buf.WriteString(i.attributes.getShrink())
	buf.WriteString(i.attributes.getBasis())
	buf.WriteString(i.attributes.getAlignSelf())
	buf.WriteString("}\n")
	for _, screenSize := range ScreenSizes {
		if mediaQuery, exist := i.mediaQueries[screenSize]; exist {
			buf.WriteString(string(screenSize) + " {" + " #" + i.id + " {")
			buf.WriteString(mediaQuery.getPadding())
			buf.WriteString(mediaQuery.getMargin())
			buf.WriteString(mediaQuery.getBorder())
			buf.WriteString(mediaQuery.getBackground())
			buf.WriteString(mediaQuery.getOrder())
			buf.WriteString(mediaQuery.getGrow())
			buf.WriteString(mediaQuery.getShrink())
			buf.WriteString(mediaQuery.getBasis())
			buf.WriteString(mediaQuery.getAlignSelf())
			buf.WriteString("}}\n")
		}
	}
	return buf.String()
}
