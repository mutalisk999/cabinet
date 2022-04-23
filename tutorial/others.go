package tutorial

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func otherScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax()
}
