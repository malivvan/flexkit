package flex

import (
	"github.com/satnamram/flexkit/dom"
)

func (c *Container) Render() *dom.Element {
	containerDiv := dom.NewElement("div").Set("id", c.id)

	styleTag := dom.NewElement("style")
	styleTag.Set("innerHTML", c.CSS())
	containerDiv.Append(styleTag)

	for _, item := range c.items {

		itemRoot := item.renderable.Render()
		if item.expandWidth {
			itemRoot.AddClass("expand-width")
		}
		if item.expandHeight {
			itemRoot.AddClass("expand-height")
		}

		itemWrapper := dom.NewElement("div").
			Set("id", item.id).
			Append(itemRoot)


		// hold item reference and enforce show/hide
		item.mutex.Lock()
		item.ref = itemWrapper
		if item.hidden {
			item.ref.AddClass("hidden")
		}
		item.mutex.Unlock()

		containerDiv.Append(itemWrapper)
	}
	return containerDiv
}

func (c *Container) RenderToBody() {
	dom.BODY.Set("innerHTML", "")
	dom.BODY.Append(c.Padding("0px").Margin("0px").Render())
}
