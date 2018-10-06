package flex

import "github.com/satnamram/flexkit/dom"

func init() {
	dom.HEAD.Append(dom.NewElement("style").Set("innerHTML", scrollbarCSS).SetAttribute("name", "scrollbar.css"))
	dom.HEAD.Append(dom.NewElement("style").Set("innerHTML", CSS).SetAttribute("name", "flex.css"))
}

const CSS = `
body > :first-child {
	height: 100vh;
	width: 100vw;
}
body > :first-child > * {
	flex-grow: 0;
	flex-shrink: 0;
	flex-basis: 0;
}
.expand-height {
	height: 100%;
	overflow-y: auto;
}
.expand-width {
	width: 100%;
	overflow-x: auto;
}
.hidden {
	display:none;
}`

const scrollbarCSS = `
::-webkit-scrollbar-track{
	-webkit-box-shadow: inset 0 0 6px rgba(0,0,0,0.3);
	border-radius: 10px;
	background-color: #F5F5F5;
}
::-webkit-scrollbar {
	width: 12px;
}
::-webkit-scrollbar-thumb{
	border-radius: 10px;
	-webkit-box-shadow: inset 0 0 6px rgba(0,0,0,.3);
	background-color: #D62929;
}`
