package tutorial

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/mutalisk999/cabinet/utils"
	"math/rand"
	"strconv"
	"strings"
	"time"
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
					go func() {
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
					}()
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

func randomPasswordScreen(_ fyne.Window) fyne.CanvasObject {
	randomPassLabel := widget.NewLabelWithStyle("please set password generation rules:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})

	includeCheckGroup := widget.NewCheckGroup([]string{"0-9", "a-z", "A-Z", "!@#$%^&*"}, nil)
	includeCheckGroup.Horizontal = true
	includeCheckGroup.SetSelected([]string{"0-9", "a-z", "A-Z"})

	excludeSlider := widget.NewSlider(0, 1)
	excludeSlider.Value = excludeSlider.Min

	excludeEntry := utils.NewCharExcludeEntry()
	excludeEntry.Text = "iIl1o0O"

	passLengthEntry := utils.NewPassLengthEntry()
	passLengthEntry.Text = "8"

	passCountEntry := utils.NewPassCountEntry()
	passCountEntry.Text = "1"

	passGeneratedResultEntry := widget.NewMultiLineEntry()
	passGeneratedResultEntry.Wrapping = fyne.TextWrapBreak
	passGeneratedResultEntry.SetPlaceHolder("This entry is for generated result of random password")

	return container.NewVBox(
		widget.NewSeparator(),
		randomPassLabel,
		container.NewHBox(widget.NewLabel("include:"), includeCheckGroup),
		container.NewBorder(nil, nil,
			container.NewHBox(widget.NewLabel("exclude:"), excludeEntry),
			excludeSlider),
		container.NewHBox(widget.NewLabel("length:"), passLengthEntry, widget.NewLabel("[1-100]")),
		container.NewHBox(widget.NewLabel("count:"), passCountEntry, widget.NewLabel("[1-1000]")),
		container.NewBorder(nil, nil, nil,
			container.NewHBox(
				widget.NewButton("generate", func() {
					go func() {
						validChars := make([]byte, 0)
						for _, selected := range includeCheckGroup.Selected {
							if selected == "0-9" {
								validChars = append(validChars, "0123456789"...)
							} else if selected == "a-z" {
								validChars = append(validChars, "abcdefghijklmnopqrstuvwxyz"...)
							} else if selected == "A-Z" {
								validChars = append(validChars, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"...)
							} else if selected == "!@#$%^&*" {
								validChars = append(validChars, "!@#$%^&*"...)
							}
						}
						validCharsStr := string(validChars)
						if excludeSlider.Value == excludeSlider.Max {
							for _, chr := range excludeEntry.Text {
								validCharsStr = strings.ReplaceAll(validCharsStr, string(chr), "")
							}
						}

						if passLengthEntry.Validate() != nil {
							return
						}
						if passCountEntry.Validate() != nil {
							return
						}
						passLength, _ := strconv.Atoi(passLengthEntry.Text)
						passCount, _ := strconv.Atoi(passCountEntry.Text)

						rand.Seed(time.Now().UnixNano())
						randomPassAll := make([]string, 0)
						for i := 0; i < passCount; i++ {
							randomPass := make([]byte, passLength)
							for j := 0; j < passLength; j++ {
								randomPass[j] = validCharsStr[rand.Intn(len(validCharsStr))]
							}
							randomPassAll = append(randomPassAll, string(randomPass))
						}
						passGeneratedResultEntry.SetText(strings.Join(randomPassAll, "\n"))
					}()
				}),
				widget.NewButton("copy", func() {
					_ = clipboard.WriteAll(passGeneratedResultEntry.Text)
				}),
				widget.NewButton("clear", func() {
					passGeneratedResultEntry.SetText("")
				})),
		),
		passGeneratedResultEntry,
	)
}
