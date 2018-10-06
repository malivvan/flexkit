package flex

import (
	"bytes"
)

type Container struct {
	id    string
	items []*Item

	attributes   *containerAttributes
	mediaQueries map[ScreenSize]*containerAttributes
}

func NewContainer() *Container {
	return &Container{
		id:    generateUniqueID(),
		items: []*Item{},
		attributes: newDefaultContainerAttributes(),
		mediaQueries: make(map[ScreenSize]*containerAttributes),
	}
}

func newDefaultContainerAttributes() *containerAttributes {
	return &containerAttributes{
		direction:      Row,
		wrap:           NoWrap,
		justifyContent: Start,
		alignItems:     Stretch,
		alignContent:   Stretch,
	}
}

type containerAttributes struct {
	background string
	border 			string
	padding        string
	margin         string
	wrap           FlexWrap
	direction      FlexDirection
	justifyContent FlexAlignment
	alignItems     FlexAlignment
	alignContent   FlexAlignment
}

func (ca *containerAttributes) getBackground() string {
	if ca.background == "" {
		return ""
	}
	return "background-color:" + ca.background + ";"
}

func (ca *containerAttributes) getBorder() string {
	if ca.border == "" {
		return ""
	}
	return "border:" + ca.border + ";"
}

func (ca *containerAttributes) getPadding() string {
	if ca.padding == "" {
		return ""
	}
	return "padding:" + ca.padding + ";"
}

func (ca *containerAttributes) getMargin() string {
	if ca.margin == "" {
		return ""
	}
	return "margin:" + ca.margin + ";"
}

func (ca *containerAttributes) getWrap(omitDefault bool) string {
	if omitDefault && ca.wrap == NoWrap {
		return ""
	}
	return string(ca.wrap)
}

func (ca *containerAttributes) getDirection(omitDefault bool) string {
	if omitDefault && ca.direction == Row {
		return ""
	}
	return string(ca.direction)
}

func (ca *containerAttributes) getJustifyContent() string {
	if ca.justifyContent == Start {
		return ""
	}
	return ca.justifyContent.justifyContent()
}

func (ca *containerAttributes) getAlignItems() string {
	if ca.alignItems == Stretch {
		return ""
	}
	return ca.alignItems.alignItems()
}

func (ca *containerAttributes) getAlignContent() string {
	if ca.alignContent == Stretch {
		return ""
	}
	return ca.alignContent.alignContent()
}

func (c *Container) CSS() string {
	var buf bytes.Buffer
	buf.WriteString("#" + c.id + " {")
	buf.WriteString("display:flex;")
	buf.WriteString(c.attributes.getPadding())
	buf.WriteString(c.attributes.getMargin())
	buf.WriteString(c.attributes.getBorder())
	buf.WriteString(c.attributes.getBackground())
	buf.WriteString(c.attributes.getDirection(true))
	buf.WriteString(c.attributes.getWrap(true))
	buf.WriteString(c.attributes.getJustifyContent())
	buf.WriteString(c.attributes.getAlignContent())
	buf.WriteString(c.attributes.getAlignItems())
	buf.WriteString("}\n")
	for _, screenSize := range ScreenSizes {
		if mediaQuery, exist := c.mediaQueries[screenSize]; exist {
			buf.WriteString(string(screenSize) + " {"+" #" + c.id+" {")
			buf.WriteString(mediaQuery.getPadding())
			buf.WriteString(mediaQuery.getMargin())
			buf.WriteString(mediaQuery.getBorder())
			buf.WriteString(mediaQuery.getBackground())
			buf.WriteString(mediaQuery.getDirection(false))
			buf.WriteString(mediaQuery.getWrap(false))
			buf.WriteString(mediaQuery.getJustifyContent())
			buf.WriteString(mediaQuery.getAlignContent())
			buf.WriteString(mediaQuery.getAlignItems())
			buf.WriteString("}}\n")
		}
	}
	for _, item := range c.items {
		buf.WriteString(item.generateCSS())
	}
	return buf.String()
}

func (c *Container) Background(v string, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.background = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.background = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].background = v
	}
	return c
}


func (c *Container) Border(v string, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.border = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.border = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].border = v
	}
	return c
}


func (c *Container) Padding(v string, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.padding = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.padding = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].padding = v
	}
	return c
}

func (c *Container) Margin(v string, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.margin = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.margin = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].margin = v
	}
	return c
}

func (c *Container) Wrap(v FlexWrap, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.wrap = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.wrap = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].wrap = v
	}
	return c
}

func (c *Container) Direction(v FlexDirection, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.direction = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.direction = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].direction = v
	}
	return c
}
func (c *Container) JustifyContent(v FlexAlignment, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.justifyContent = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.justifyContent = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].justifyContent = v
	}
	return c
}

func (c *Container) AlignItems(v FlexAlignment, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.alignItems = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.alignItems = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].alignItems = v
	}
	return c
}

func (c *Container) AlignContent(v FlexAlignment, sizes ...ScreenSize) *Container {
	if len(sizes) == 0 {
		c.attributes.alignContent = v
		return c
	}
	for _, size := range sizes {
		if size == ScreenSizeXSmall {
			c.attributes.alignContent = v
			continue
		}
		if c.mediaQueries[size] == nil {
			c.mediaQueries[size] = newDefaultContainerAttributes()
		}
		c.mediaQueries[size].alignContent = v
	}
	return c
}

func (c *Container) Append(i *Item) *Container {
	c.items = append(c.items, i)
	return c
}
