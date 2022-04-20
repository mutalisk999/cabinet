package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/mutalisk999/cabinet/tutorial"
	"github.com/mutalisk999/cabinet/utils"
	"log"
)

const (
	preferenceCurrentTutorial = "currentTutorial"
)

var (
	topWindow fyne.Window
)

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

func makeToolBar(a fyne.App, w fyne.Window) fyne.CanvasObject {
	return container.New(layout.NewHBoxLayout(),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(utils.ThemeDark())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(utils.ThemeLight())
		}),
		widget.NewButton("x1", func() {
			w.Resize(utils.WinSizeX1())
			w.CenterOnScreen()
		}),
		widget.NewButton("x1.5", func() {
			w.Resize(utils.WinSizeX1_5())
			w.CenterOnScreen()
		}),
		widget.NewButton("x2", func() {
			w.Resize(utils.WinSizeX2())
			w.CenterOnScreen()
		}),
	)
}

func makeNav(setTutorial func(tutorial tutorial.Tutorial), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return tutorial.TutorialIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := tutorial.TutorialIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := tutorial.Tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := tutorial.Tutorials[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "convert")
		tree.Select(currentPref)
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}

func main() {
	cabApp := app.NewWithID("cabinet")
	logLifecycle(cabApp)

	cabWin := cabApp.NewWindow("cabinet")
	topWindow = cabWin

	// widgets (title, introduction, view)
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go here")
	intro.Wrapping = fyne.TextWrapWord
	content := container.NewMax()

	setTutorial := func(t tutorial.Tutorial) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)
		content.Objects = []fyne.CanvasObject{t.View(cabWin)}
		content.Refresh()
	}

	treeBorder := makeNav(setTutorial, true)
	viewBorder := container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro),
		nil, nil, nil, content)
	hSplit := container.NewHSplit(treeBorder, viewBorder)
	hSplit.Offset = 0.1

	buttons := makeToolBar(cabApp, cabWin)
	vSplit := container.NewVSplit(buttons, hSplit)
	vSplit.Offset = 0.05

	cabWin.SetContent(vSplit)
	cabWin.Resize(utils.WinSizeX2())
	cabWin.ShowAndRun()
}
