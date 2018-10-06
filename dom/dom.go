package dom

import "github.com/gopherjs/gopherjs/js"

var (
	DOC  = js.Global.Get("document")
	HEAD = getElement("head")
	BODY = getElement("body")
)
