package dom

func init() {
	HEAD.Append(NewElement("script").Set("innerHTML", domJS).SetAttribute("name", "dom.js"))
}

const domJS = `
function makePropertyWriteable(obj, property) {Object.defineProperty(obj, property, {writable: true});}
`
