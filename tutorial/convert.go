package tutorial

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"math/big"
)

func convertScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter()
}

type binEntry struct {
	widget.Entry
}

func newBinEntry() *binEntry {
	e := &binEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[01]*$`, "Must contain a binary number")
	return e
}

type octEntry struct {
	widget.Entry
}

func newOctEntry() *octEntry {
	e := &octEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-7]*$`, "Must contain a octal number")
	return e
}

type decEntry struct {
	widget.Entry
}

func newDecEntry() *decEntry {
	e := &decEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9]*$`, "Must contain a decimal number")
	return e
}

type hexEntry struct {
	widget.Entry
}

func newHexEntry() *hexEntry {
	e := &hexEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9a-fA-F]*$`, "Must contain a hexadecimal number")
	return e
}

func baseConvertScreen(_ fyne.Window) fyne.CanvasObject {
	binEntryValidated := newBinEntry()
	binEntryValidated.SetPlaceHolder("Must contain a binary number")
	octEntryValidated := newOctEntry()
	octEntryValidated.SetPlaceHolder("Must contain a octal number")
	decEntryValidated := newDecEntry()
	decEntryValidated.SetPlaceHolder("Must contain a decimal number")
	hexEntryValidated := newHexEntry()
	hexEntryValidated.SetPlaceHolder("Must contain a hexadecimal number")

	binEntryValidated.OnChanged = func(valStr string) {
		val, ok := new(big.Int).SetString(valStr, 2)
		if ok {
			octEntryValidated.SetText(val.Text(8))
			decEntryValidated.SetText(val.Text(10))
			hexEntryValidated.SetText(val.Text(16))
		} else {
			octEntryValidated.SetText("")
			decEntryValidated.SetText("")
			hexEntryValidated.SetText("")
		}
	}
	octEntryValidated.OnChanged = func(valStr string) {
		val, ok := new(big.Int).SetString(valStr, 8)
		if ok {
			binEntryValidated.SetText(val.Text(2))
			decEntryValidated.SetText(val.Text(10))
			hexEntryValidated.SetText(val.Text(16))
		} else {
			binEntryValidated.SetText("")
			decEntryValidated.SetText("")
			hexEntryValidated.SetText("")
		}
	}
	decEntryValidated.OnChanged = func(valStr string) {
		val, ok := new(big.Int).SetString(valStr, 10)
		if ok {
			binEntryValidated.SetText(val.Text(2))
			octEntryValidated.SetText(val.Text(8))
			hexEntryValidated.SetText(val.Text(16))
		} else {
			binEntryValidated.SetText("")
			octEntryValidated.SetText("")
			hexEntryValidated.SetText("")
		}
	}
	hexEntryValidated.OnChanged = func(valStr string) {
		val, ok := new(big.Int).SetString(valStr, 16)
		if ok {
			binEntryValidated.SetText(val.Text(2))
			octEntryValidated.SetText(val.Text(8))
			decEntryValidated.SetText(val.Text(10))
		} else {
			binEntryValidated.SetText("")
			octEntryValidated.SetText("")
			decEntryValidated.SetText("")
		}
	}

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Please input a binary number:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(binEntryValidated.Text)
			})), binEntryValidated, nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Please input a octal number:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(octEntryValidated.Text)
			})), octEntryValidated, nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Please input a decimal number:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(decEntryValidated.Text)
			})), decEntryValidated, nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Please input a hexadecimal number:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(hexEntryValidated.Text)
			})), hexEntryValidated, nil, nil),
	)
}
