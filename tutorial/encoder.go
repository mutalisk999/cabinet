package tutorial

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/jamesruan/go-rfc1924/base85"
	"github.com/mr-tron/base58"
	"github.com/mutalisk999/cabinet/utils"
	"github.com/whyrusleeping/base32"
	"unicode/utf8"
)

func encodeScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax()
}

func base64EncodeScreen(win fyne.Window) fyne.CanvasObject {
	transformLabel := widget.NewLabelWithStyle("Encode:",
		fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true})

	transformSlider := widget.NewSlider(0, 1)
	transformSlider.Value = transformSlider.Min

	radioGroup := widget.NewRadioGroup([]string{
		"base16", "base32", "base64", "base58", "base85"}, nil)
	radioGroup.Horizontal = true
	radioGroup.SetSelected("base64")

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.Wrapping = fyne.TextWrapBreak
	inputEntry.SetPlaceHolder("This entry is for origin data needed to be encode/decode")

	outputEntry := widget.NewMultiLineEntry()
	outputEntry.Wrapping = fyne.TextWrapBreak
	outputEntry.SetPlaceHolder("This entry is result encoded/decoded data")

	transformSlider.OnChanged = func(valFloat float64) {
		if valFloat == transformSlider.Min {
			transformLabel.SetText("Encode")
			inputEntry.SetText("")

		} else if valFloat == transformSlider.Max {
			transformLabel.SetText("Decode")
			inputEntry.SetText("")
		}
	}

	calcFunc := func(inputStr string) {
		go func(inputStr string) {
			if transformSlider.Value == transformSlider.Min {
				if radioGroup.Selected == "base16" {
					outputEntry.SetText(hex.EncodeToString([]byte(inputEntry.Text)))
				} else if radioGroup.Selected == "base32" {
					outputEntry.SetText(base32.StdEncoding.EncodeToString([]byte(inputEntry.Text)))
				} else if radioGroup.Selected == "base64" {
					outputEntry.SetText(base64.StdEncoding.EncodeToString([]byte(inputEntry.Text)))
				} else if radioGroup.Selected == "base58" {
					outputEntry.SetText(base58.Encode([]byte(inputEntry.Text)))
				} else if radioGroup.Selected == "base85" {
					outputEntry.SetText(base85.EncodeToString([]byte(inputEntry.Text)))
				} else {
					return
				}
			} else if transformSlider.Value == transformSlider.Max {
				if radioGroup.Selected == "base16" {
					res, err := hex.DecodeString(inputEntry.Text)
					if err != nil {
						outputEntry.SetText(err.Error())
					} else {
						outputEntry.SetText(string(utils.BytesToRunes(res)))
					}
				} else if radioGroup.Selected == "base32" {
					res, err := base32.StdEncoding.DecodeString(inputEntry.Text)
					if err != nil {
						outputEntry.SetText(err.Error())
					} else {
						outputEntry.SetText(string(utils.BytesToRunes(res)))
					}
				} else if radioGroup.Selected == "base64" {
					res, err := base64.StdEncoding.DecodeString(inputEntry.Text)
					if err != nil {
						outputEntry.SetText(err.Error())
					} else {
						outputEntry.SetText(string(utils.BytesToRunes(res)))
					}
				} else if radioGroup.Selected == "base58" {
					res, err := base58.Decode(inputEntry.Text)
					if err != nil {
						outputEntry.SetText(err.Error())
					} else {
						outputEntry.SetText(string(utils.BytesToRunes(res)))
					}
				} else if radioGroup.Selected == "base85" {
					res, err := base85.DecodeString(inputEntry.Text)
					if err != nil {
						outputEntry.SetText(err.Error())
					} else {
						outputEntry.SetText(string(utils.BytesToRunes(res)))
					}
				} else {
					return
				}
			}
		}(inputStr)
	}
	radioGroup.OnChanged = calcFunc
	inputEntry.OnChanged = calcFunc
	outputEntry.OnChanged = calcFunc

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewBorder(nil, nil, transformLabel, transformSlider),
		container.NewBorder(nil, nil, widget.NewLabelWithStyle("Type:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}), radioGroup),
		widget.NewSeparator(),
		container.NewBorder(nil, nil, widget.NewLabelWithStyle("Input Data:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			container.NewHBox(
				widget.NewButton("open file", func() {
					fileOpen := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
						if err != nil {
							dialog.ShowError(err, win)
							return
						}
						if reader == nil {
							return
						}
						defer reader.Close()

						readBuf100k := make([]byte, 100*1024)
						n, err := reader.Read(readBuf100k)
						if err != nil {
							dialog.ShowError(err, win)
							return
						}
						if n >= 100*1024 {
							dialog.ShowError(errors.New("not support: file size is greater than 100k"), win)
							return
						}

						if !utf8.Valid(readBuf100k[0:n]) {
							dialog.ShowError(errors.New("not a valid utf8 encoding file"), win)
							return
						}

						inputEntry.SetText(string(utils.BytesToRunes(readBuf100k[0:n])))
					}, win)
					fileOpen.Resize(fyne.Size{Width: 800, Height: 600})
					fileOpen.Show()
				}),
				widget.NewButton("clear", func() {
					inputEntry.SetText("")
				}),
			)),
		inputEntry,
		widget.NewSeparator(),
		container.NewBorder(nil, nil, widget.NewLabelWithStyle("Output Data:",
			fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(outputEntry.Text)
			}),
		),
		outputEntry,
	)
}
