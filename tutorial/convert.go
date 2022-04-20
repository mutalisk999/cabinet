package tutorial

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"github.com/mutalisk999/cabinet/utils"
	"math/big"
	"strconv"
	"time"
)

func convertScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewCenter()
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

	utcYearEntryValidated := utils.NewYearEntry()
	utcYearEntryValidated.SetPlaceHolder("Year")
	utcMonthEntryValidated := utils.NewMonthEntry()
	utcMonthEntryValidated.SetPlaceHolder("Month")
	utcDayEntryValidated := utils.NewDayEntry()
	utcDayEntryValidated.SetPlaceHolder("Day")
	utcHourEntryValidated := utils.NewHourEntry()
	utcHourEntryValidated.SetPlaceHolder("Hour")
	utcMinuteEntryValidated := utils.NewMinuteEntry()
	utcMinuteEntryValidated.SetPlaceHolder("Minute")
	utcSecondEntryValidated := utils.NewSecondEntry()
	utcSecondEntryValidated.SetPlaceHolder("Second")

	localYearEntryValidated := utils.NewYearEntry()
	localYearEntryValidated.SetPlaceHolder("Year")
	localMonthEntryValidated := utils.NewMonthEntry()
	localMonthEntryValidated.SetPlaceHolder("Month")
	localDayEntryValidated := utils.NewDayEntry()
	localDayEntryValidated.SetPlaceHolder("Day")
	localHourEntryValidated := utils.NewHourEntry()
	localHourEntryValidated.SetPlaceHolder("Hour")
	localMinuteEntryValidated := utils.NewMinuteEntry()
	localMinuteEntryValidated.SetPlaceHolder("Minute")
	localSecondEntryValidated := utils.NewSecondEntry()
	localSecondEntryValidated.SetPlaceHolder("Second")

	timestampEntryValidated.OnChanged = func(valStr string) {
		timestamp, _ := strconv.Atoi(valStr)
		timeUtc := time.Unix(int64(timestamp), int64(0)).UTC()
		yearUtc, monthUtc, dayUtc := timeUtc.Date()
		hourUtc, minuteUtc, secondUtc := timeUtc.Clock()

		timeLocal := time.Unix(int64(timestamp), int64(0)).Local()
		yearLocal, monthLocal, dayLocal := timeLocal.Date()
		hourLocal, minuteLocal, secondLocal := timeLocal.Clock()

		utcYearEntryValidated.SetText(strconv.Itoa(yearUtc))
		utcMonthEntryValidated.SetText(strconv.Itoa(int(monthUtc)))
		utcDayEntryValidated.SetText(strconv.Itoa(dayUtc))
		utcHourEntryValidated.SetText(strconv.Itoa(hourUtc))
		utcMinuteEntryValidated.SetText(strconv.Itoa(minuteUtc))
		utcSecondEntryValidated.SetText(strconv.Itoa(secondUtc))

		localYearEntryValidated.SetText(strconv.Itoa(yearLocal))
		localMonthEntryValidated.SetText(strconv.Itoa(int(monthLocal)))
		localDayEntryValidated.SetText(strconv.Itoa(dayLocal))
		localHourEntryValidated.SetText(strconv.Itoa(hourLocal))
		localMinuteEntryValidated.SetText(strconv.Itoa(minuteLocal))
		localSecondEntryValidated.SetText(strconv.Itoa(secondLocal))
	}

	utcTimeToTimestamp := func(_ string) {
		if utcYearEntryValidated.Validate() == nil &&
			utcMonthEntryValidated.Validate() == nil &&
			utcDayEntryValidated.Validate() == nil &&
			utcHourEntryValidated.Validate() == nil &&
			utcMinuteEntryValidated.Validate() == nil &&
			utcSecondEntryValidated.Validate() == nil {
			year, _ := strconv.Atoi(utcYearEntryValidated.Text)
			month, _ := strconv.Atoi(utcMonthEntryValidated.Text)
			day, _ := strconv.Atoi(utcDayEntryValidated.Text)
			hour, _ := strconv.Atoi(utcHourEntryValidated.Text)
			minute, _ := strconv.Atoi(utcMinuteEntryValidated.Text)
			second, _ := strconv.Atoi(utcSecondEntryValidated.Text)
			utcTime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
			timestampEntryValidated.SetText(strconv.Itoa(int(utcTime.Unix())))
		}
	}
	utcYearEntryValidated.OnChanged = utcTimeToTimestamp
	utcMonthEntryValidated.OnChanged = utcTimeToTimestamp
	utcDayEntryValidated.OnChanged = utcTimeToTimestamp
	utcHourEntryValidated.OnChanged = utcTimeToTimestamp
	utcMinuteEntryValidated.OnChanged = utcTimeToTimestamp
	utcSecondEntryValidated.OnChanged = utcTimeToTimestamp

	localTimeToTimestamp := func(_ string) {
		if localYearEntryValidated.Validate() == nil &&
			localMonthEntryValidated.Validate() == nil &&
			localDayEntryValidated.Validate() == nil &&
			localHourEntryValidated.Validate() == nil &&
			localMinuteEntryValidated.Validate() == nil &&
			localSecondEntryValidated.Validate() == nil {
			year, _ := strconv.Atoi(localYearEntryValidated.Text)
			month, _ := strconv.Atoi(localMonthEntryValidated.Text)
			day, _ := strconv.Atoi(localDayEntryValidated.Text)
			hour, _ := strconv.Atoi(localHourEntryValidated.Text)
			minute, _ := strconv.Atoi(localMinuteEntryValidated.Text)
			second, _ := strconv.Atoi(localSecondEntryValidated.Text)
			localTime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
			timestampEntryValidated.SetText(strconv.Itoa(int(localTime.Unix())))
		}
	}
	localYearEntryValidated.OnChanged = localTimeToTimestamp
	localMonthEntryValidated.OnChanged = localTimeToTimestamp
	localDayEntryValidated.OnChanged = localTimeToTimestamp
	localHourEntryValidated.OnChanged = localTimeToTimestamp
	localMinuteEntryValidated.OnChanged = localTimeToTimestamp
	localSecondEntryValidated.OnChanged = localTimeToTimestamp

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
				text := fmt.Sprintf("%04s/%02s/%02s %02s:%02s:%02s", utcYearEntryValidated.Text,
					utcMonthEntryValidated.Text, utcDayEntryValidated.Text, utcHourEntryValidated.Text,
					utcMinuteEntryValidated.Text, utcSecondEntryValidated.Text)
				_ = clipboard.WriteAll(text)
			})),
			container.NewHBox(utcYearEntryValidated, widget.NewLabel("/"),
				utcMonthEntryValidated, widget.NewLabel("/"),
				utcDayEntryValidated, widget.NewLabel(" "),
				utcHourEntryValidated, widget.NewLabel(":"),
				utcMinuteEntryValidated, widget.NewLabel(":"),
				utcSecondEntryValidated), nil, nil),
		widget.NewSeparator(),
		container.NewBorder(container.NewBorder(nil, nil,
			widget.NewLabelWithStyle("Date & Time (local):",
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true, Italic: true}),
			widget.NewButton("copy to clipboard", func() {
				text := fmt.Sprintf("%04s/%02s/%02s %02s:%02s:%02s", localYearEntryValidated.Text,
					localMonthEntryValidated.Text, localDayEntryValidated.Text, localHourEntryValidated.Text,
					localMinuteEntryValidated.Text, localSecondEntryValidated.Text)
				_ = clipboard.WriteAll(text)
			})),
			container.NewHBox(localYearEntryValidated, widget.NewLabel("/"),
				localMonthEntryValidated, widget.NewLabel("/"),
				localDayEntryValidated, widget.NewLabel(" "),
				localHourEntryValidated, widget.NewLabel(":"),
				localMinuteEntryValidated, widget.NewLabel(":"),
				localSecondEntryValidated), nil, nil),
	)
}
