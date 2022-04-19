package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(171, 125))
	} else {
		logo.SetMinSize(fyne.NewSize(228, 167))
	}

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			widget.NewHyperlink("fyne.io", parseURL("https://fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://developer.fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("sponsor", parseURL("https://fyne.io/sponsor/")),
		),
	))
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{
		"welcome": {"Welcome", "", welcomeScreen},
		"canvas": {"Canvas",
			"See the canvas capabilities.",
			welcomeScreen,
		},
		"animations": {"Animations",
			"See how to animate components.",
			welcomeScreen,
		},
		"icons": {"Theme Icons",
			"Browse the embedded icons.",
			welcomeScreen,
		},
		"containers": {"Containers",
			"Containers group other widgets and canvas objects, organising according to their layout.\n" +
				"Standard containers are illustrated in this section, but developers can also provide custom " +
				"layouts using the fyne.NewContainerWithLayout() constructor.",
			welcomeScreen,
		},
		"apptabs": {"AppTabs",
			"A container to help divide up an application into functional areas.",
			welcomeScreen,
		},
		"border": {"Border",
			"A container that positions items around a central content.",
			welcomeScreen,
		},
		"box": {"Box",
			"A container arranges items in horizontal or vertical list.",
			welcomeScreen,
		},
		"center": {"Center",
			"A container to that centers child elements.",
			welcomeScreen,
		},
		"doctabs": {"DocTabs",
			"A container to display a single document from a set of many.",
			welcomeScreen,
		},
		"grid": {"Grid",
			"A container that arranges all items in a grid.",
			welcomeScreen,
		},
		"split": {"Split",
			"A split container divides the container in two pieces that the user can resize.",
			welcomeScreen,
		},
		"scroll": {"Scroll",
			"A container that provides scrolling for it's content.",
			welcomeScreen,
		},
		"widgets": {"Widgets",
			"In this section you can see the features available in the toolkit widget set.\n" +
				"Expand the tree on the left to browse the individual tutorial elements.",
			welcomeScreen,
		},
		"accordion": {"Accordion",
			"Expand or collapse content panels.",
			welcomeScreen,
		},
		"button": {"Button",
			"Simple widget for user tap handling.",
			welcomeScreen,
		},
		"card": {"Card",
			"Group content and widgets.",
			welcomeScreen,
		},
		"entry": {"Entry",
			"Different ways to use the entry widget.",
			welcomeScreen,
		},
		"form": {"Form",
			"Gathering input widgets for data submission.",
			welcomeScreen,
		},
		"input": {"Input",
			"A collection of widgets for user input.",
			welcomeScreen,
		},
		"text": {"Text",
			"Text handling widgets.",
			welcomeScreen,
		},
		"toolbar": {"Toolbar",
			"A row of shortcut icons for common tasks.",
			welcomeScreen,
		},
		"progress": {"Progress",
			"Show duration or the need to wait for a task.",
			welcomeScreen,
		},
		"collections": {"Collections",
			"Collection widgets provide an efficient way to present lots of content.\n" +
				"The List, Table, and Tree provide a cache and re-use mechanism that make it possible to scroll through thousands of elements.\n" +
				"Use this for large data sets or for collections that can expand as users scroll.",
			welcomeScreen,
		},
		"list": {"List",
			"A vertical arrangement of cached elements with the same styling.",
			welcomeScreen,
		},
		"table": {"Table",
			"A two dimensional cached collection of cells.",
			welcomeScreen,
		},
		"tree": {"Tree",
			"A tree based arrangement of cached elements with the same styling.",
			welcomeScreen,
		},
		"dialogs": {"Dialogs",
			"Work with dialogs.",
			welcomeScreen,
		},
		"windows": {"Windows",
			"Window function demo.",
			welcomeScreen,
		},
		"binding": {"Data Binding",
			"Connecting widgets to a data source.",
			welcomeScreen},
		"advanced": {"Advanced",
			"Debug and advanced information.",
			welcomeScreen,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"":            {"welcome", "canvas", "animations", "icons", "widgets", "collections", "containers", "dialogs", "windows", "binding", "advanced"},
		"collections": {"list", "table", "tree"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}
)

func main() {
	cabApp := app.NewWithID("cabinet")
	logLifecycle(cabApp)

	cabWin := cabApp.NewWindow("cabinet")
	topWindow = cabWin

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	setTutorial := func(t Tutorial) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)
		content.Objects = []fyne.CanvasObject{t.View(cabWin)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)

	split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
	split.Offset = 0.1
	cabWin.SetContent(split)

	cabWin.Resize(fyne.NewSize(800, 400))
	cabWin.ShowAndRun()
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func makeNav(setTutorial func(tutorial Tutorial), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return TutorialIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := TutorialIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := Tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := Tutorials[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
