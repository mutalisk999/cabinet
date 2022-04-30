package tutorial

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/mutalisk999/cabinet/utils"
	"strconv"
	"strings"
)

func otherScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax()
}

func uuidScreen(_ fyne.Window) fyne.CanvasObject {
	hyphenLabel := widget.NewLabelWithStyle("hyphenated",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})
	hyphenSlider := widget.NewSlider(0, 1)
	hyphenSlider.Value = hyphenSlider.Max
	hyphenSlider.OnChanged = func(valFloat float64) {
		if valFloat == hyphenSlider.Min {
			hyphenLabel.SetText("non-hyphenated")
		} else if valFloat == hyphenSlider.Max {
			hyphenLabel.SetText("hyphenated")
		}
	}

	upperLabel := widget.NewLabelWithStyle("lower",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})
	upperSlider := widget.NewSlider(0, 1)
	upperSlider.Value = upperSlider.Min
	upperSlider.OnChanged = func(valFloat float64) {
		if valFloat == upperSlider.Min {
			upperLabel.SetText("lower")
		} else if valFloat == upperSlider.Max {
			upperLabel.SetText("UPPER")
		}
	}

	uuidVersionLabel := widget.NewLabelWithStyle("uuid version:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})
	uuidVersionSelect := widget.NewSelect([]string{"uuid 1", "uuid 4"}, nil)
	uuidVersionSelect.SetSelected("uuid 1")

	generateCountEntry := utils.NewNumberEntry()
	generateCountEntry.SetText("1")
	generateCountEntry.OnChanged = func(valStr string) {
		count, err := strconv.Atoi(generateCountEntry.Text)
		if err != nil {
			generateCountEntry.SetText("1")
			return
		}
		if count < 0 {
			generateCountEntry.SetText("0")
			return
		}
		if count > 1000 {
			generateCountEntry.SetText("1000")
			return
		}
	}

	uuidGeneratedResultEntry := widget.NewMultiLineEntry()
	uuidGeneratedResultEntry.Wrapping = fyne.TextWrapWord
	uuidGeneratedResultEntry.SetPlaceHolder("This entry is for generated result of uuid")

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewBorder(nil, nil, hyphenLabel, hyphenSlider),
		container.NewBorder(nil, nil, upperLabel, upperSlider),
		container.NewBorder(nil, nil, uuidVersionLabel, uuidVersionSelect),
		widget.NewSeparator(),
		container.NewBorder(nil, nil,
			container.NewHBox(widget.NewLabelWithStyle("generate count:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}), generateCountEntry),
			container.NewHBox(
				widget.NewButton("generate", func() {
					if generateCountEntry.Validate() != nil {
						return
					}
					count, err := strconv.Atoi(generateCountEntry.Text)
					if err != nil {
						uuidGeneratedResultEntry.SetText(err.Error())
						return
					}

					uuidsTotal := make([]string, 0)
					for i := 0; i < count; i++ {
						u := ""
						if uuidVersionSelect.Selected == "uuid 1" {
							u1, err := uuid.NewUUID()
							if err != nil {
								uuidGeneratedResultEntry.SetText(err.Error())
								return
							}
							u = u1.String()
						} else if uuidVersionSelect.Selected == "uuid 4" {
							u4 := uuid.New()
							u = u4.String()
						}
						if upperSlider.Value == upperSlider.Max {
							u = strings.ToUpper(u)
						}
						if hyphenSlider.Value == hyphenSlider.Min {
							u = strings.ReplaceAll(u, "-", "")
						}
						uuidsTotal = append(uuidsTotal, u)
					}
					uuidGeneratedResultEntry.SetText(strings.Join(uuidsTotal, "\n"))
				}),
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(uuidGeneratedResultEntry.Text)
				}),
				widget.NewButton("clear", func() {
					uuidGeneratedResultEntry.SetText("")
				}),
			),
		),
		uuidGeneratedResultEntry,
	)
}
