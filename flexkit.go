// Package flexkit implements a programmatic user interface for single page applications.
package flexkit

import (
	"github.com/gopherjs/gopherjs/js"
	"time"
	"sync"
	"github.com/satnamram/flexkit/flex"
	"github.com/satnamram/flexkit/dom"
)

var application *App

func Goto(view string) {
	if application == nil {
		panic("flexkit not initialized")
	}
	application.mutex.Lock()
	for _, v := range application.views {
		if v.name == view {
			application.vstack = append(application.vstack, v.name)
			v.container.RenderToBody()
			break
		}
	}
	application.mutex.Unlock()
}

func Back() {
	if application == nil {
		panic("flexkit not initialized")
	}
	application.mutex.Lock()

	// no parent - nothing to do
	if len(application.vstack) <= 1 {
		application.mutex.Unlock()
		return
	}
	application.vstack = application.vstack[:len(application.vstack)-1]

	// goto parent
	for _, v := range application.views {
		if v.name == application.vstack[len(application.vstack)-1] {
			v.container.RenderToBody()
			break
		}
	}
	application.mutex.Unlock()
}

type App struct {
	mutex sync.Mutex

	views  []*appView
	vstack []string
}

type appView struct {
	name      string
	container *flex.Container
}

func Init() *App {
	a := &App{
		views: []*appView{},
	}
	a.mutex.Lock() // lock until start
	return a
}

func (a *App) View(name string, container *flex.Container) *App {
	for _, view := range a.views {
		if view.name == name {
			panic("view '" + name + "' does already exist")
		}
	}
	a.views = append(a.views, &appView{
		name:      name,
		container: container,
	})
	return a
}

func (a *App) Start(initalView string) {
	for _, view := range a.views {
		if view.name == initalView {

			// startup, unlock and block
			a.vstack = append(a.vstack, initalView)
			view.container.RenderToBody()
			application = a
			application.mutex.Unlock()
			for {
				time.Sleep(60 * time.Minute)
			}
		}
	}
	panic("view '" + initalView + "' does not exist!")
}

func (a *App) Title(t string) *App {
	dom.DOC.Set("title", t)
	return a
}

var currentFavicon *js.Object

// Favicon as base64 href entry ("data:image/png;base64,iVBORw0KGgoAAAA...").
func (a *App) Favicon(href string) *App {
	head := dom.DOC.Get("head")
	if currentFavicon != nil {
		head.Call("removeChild", *currentFavicon)
	}
	faviconLink := dom.DOC.Call("createElement", "link")
	currentFavicon = faviconLink
	faviconLink.Set("rel", "shortcut icon")
	faviconLink.Set("href", href)
	dom.DOC.Get("head").Call("prepend", faviconLink)
	return a
}

func (a *App) DarkTheme(backgroundColor string) *App {
	if backgroundColor == "default" || backgroundColor == "" {
		backgroundColor = "#222"
	}
	dom.DOC.Get("documentElement").Get("style").Set("background-color", backgroundColor)
	dom.BODY.AddClass("uk-light")
	return a
}

func (a *App) LightTheme(backgroundColor string) *App {
	if backgroundColor == "default" || backgroundColor == "" {
		backgroundColor = "#f8f8f8"
	}
	dom.DOC.Get("documentElement").Get("style").Set("background-color", backgroundColor)
	dom.BODY.RemoveClass("uk-light")
	return a
}
