package tutorial

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/cleanmachine1/capitalise"
	"github.com/iancoleman/strcase"
	"github.com/mutalisk999/cabinet/utils"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func convertScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewMax()
}

func baseConvertScreen(_ fyne.Window) fyne.CanvasObject {
	binEntryValidated := utils.NewBinEntry()
	binEntryValidated.SetPlaceHolder("Must contain a binary number")
	octEntryValidated := utils.NewOctEntry()
	octEntryValidated.SetPlaceHolder("Must contain a octal number")
	decEntryValidated := utils.NewDecEntry()
	decEntryValidated.SetPlaceHolder("Must contain a decimal number")
	hexEntryValidated := utils.NewHexEntry()
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

func timeConvertScreen(_ fyne.Window) fyne.CanvasObject {
	timestampEntryValidated := utils.NewTimestampEntry()
	timestampEntryValidated.SetPlaceHolder("Must contain a timestamp number")

	utcDateTimeEntryValidated := utils.NewDateTimeEntry()
	utcDateTimeEntryValidated.SetPlaceHolder("Must contain a date time string, such as 1970/01/01 00:00:00")

	localDateTimeEntryValidated := utils.NewDateTimeEntry()
	localDateTimeEntryValidated.SetPlaceHolder("Must contain a date time string, such as 1970/01/01 00:00:00")

	timestampEntryValidated.OnChanged = func(valStr string) {
		if timestampEntryValidated.Validate() != nil {
			return
		}

		timestamp, _ := strconv.Atoi(valStr)
		timeUtc := time.Unix(int64(timestamp), int64(0)).UTC()
		yearUtc, monthUtc, dayUtc := timeUtc.Date()
		hourUtc, minuteUtc, secondUtc := timeUtc.Clock()

		timeLocal := time.Unix(int64(timestamp), int64(0)).Local()
		yearLocal, monthLocal, dayLocal := timeLocal.Date()
		hourLocal, minuteLocal, secondLocal := timeLocal.Clock()

		utcDateTimeEntryValidated.SetText(fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", yearUtc,
			monthUtc, dayUtc, hourUtc, minuteUtc, secondUtc))

		localDateTimeEntryValidated.SetText(fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", yearLocal,
			monthLocal, dayLocal, hourLocal, minuteLocal, secondLocal))
	}

	utcDateTimeEntryValidated.OnChanged = func(valStr string) {
		if utcDateTimeEntryValidated.Validate() != nil {
			return
		}

		yearUtc, monthUtc, dayUtc := 0, 0, 0
		hourUtc, minuteUtc, secondUtc := 0, 0, 0
		n, err := fmt.Sscanf(valStr, "%04d/%02d/%02d %02d:%02d:%02d",
			&yearUtc, &monthUtc, &dayUtc, &hourUtc, &minuteUtc, &secondUtc)
		if n != 6 || err != nil {
			return
		}
		utcTime := time.Date(yearUtc, time.Month(monthUtc), dayUtc,
			hourUtc, minuteUtc, secondUtc, 0, time.UTC)
		timestampEntryValidated.SetText(strconv.Itoa(int(utcTime.Unix())))
	}

	localDateTimeEntryValidated.OnChanged = func(valStr string) {
		if localDateTimeEntryValidated.Validate() != nil {
			return
		}

		yearLocal, monthLocal, dayLocal := 0, 0, 0
		hourLocal, minuteLocal, secondLocal := 0, 0, 0
		n, err := fmt.Sscanf(valStr, "%04d/%02d/%02d %02d:%02d:%02d",
			&yearLocal, &monthLocal, &dayLocal, &hourLocal, &minuteLocal, &secondLocal)
		if n != 6 || err != nil {
			return
		}
		localTime := time.Date(yearLocal, time.Month(monthLocal), dayLocal,
			hourLocal, minuteLocal, secondLocal, 0, time.Local)
		timestampEntryValidated.SetText(strconv.Itoa(int(localTime.Unix())))
	}

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Please input a timestamp number:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			container.NewBorder(nil, nil,
				widget.NewButton("now", func() {
					timestampEntryValidated.SetText(strconv.Itoa(int(time.Now().Unix())))
				}),
				widget.NewButton("copy to clipboard", func() {
					_ = clipboard.WriteAll(timestampEntryValidated.Text)
				})),
		), timestampEntryValidated, nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Date & Time (UTC):",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(utcDateTimeEntryValidated.Text)
			})),
			utcDateTimeEntryValidated, nil, nil),

		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Date & Time (local):",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(localDateTimeEntryValidated.Text)
			})),
			localDateTimeEntryValidated, nil, nil),
	)
}

func caseConvertScreen(_ fyne.Window) fyne.CanvasObject {
	caseOriginEntryValidated := utils.NewCaseOriginEntry()
	caseOriginEntryValidated.SetPlaceHolder("Must contain a string contains numbers/alphabets/underscore character/space character")

	radioGroup := widget.NewRadioGroup([]string{
		"UPPER CASE", "lower case", "Capital case", "Title Case",
		"camelCase", "PascalCase", "snake_case", "CONSTANT_CASE"}, nil)
	radioGroup.Horizontal = true
	caseResultEntryValidated := utils.NewCaseResultEntry()

	caseConvert := func(valStr string) {
		if caseOriginEntryValidated.Validate() != nil {
			return
		}
		if radioGroup.Selected == "" {
			return
		}
		caseOrigin := caseOriginEntryValidated.Text
		caseResult := ""
		radioSelected := radioGroup.Selected
		if radioSelected == "UPPER CASE" {
			caseResult = strings.ToUpper(caseOrigin)
		} else if radioSelected == "lower case" {
			caseResult = strings.ToLower(caseOrigin)
		} else if radioSelected == "Capital case" {
			caseResult = capitalise.First(strings.ToLower(caseOrigin))
		} else if radioSelected == "Title Case" {
			caseResult = strings.Title(strings.ToLower(caseOrigin))
		} else if radioSelected == "camelCase" {
			caseResult = strcase.ToLowerCamel(caseOrigin)
		} else if radioSelected == "PascalCase" {
			caseResult = strcase.ToCamel(caseOrigin)
		} else if radioSelected == "snake_case" {
			caseResult = strcase.ToSnake(caseOrigin)
		} else if radioSelected == "CONSTANT_CASE" {
			caseResult = strcase.ToScreamingSnake(caseOrigin)
		}
		caseResultEntryValidated.SetText(caseResult)
	}
	caseOriginEntryValidated.OnChanged = caseConvert
	radioGroup.OnChanged = caseConvert

	return container.NewVBox(
		widget.NewSeparator(),
		container.NewBorder(
			widget.NewLabelWithStyle("Please input a string contains numbers/alphabets/underscore character/space character:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			caseOriginEntryValidated, nil, nil), radioGroup,
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Conversion Result:",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				_ = clipboard.WriteAll(caseResultEntryValidated.Text)
			})),
			caseResultEntryValidated, nil, nil),
	)
}
