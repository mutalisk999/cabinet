package utils

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

type binEntry struct {
	widget.Entry
}

func NewBinEntry() *binEntry {
	e := &binEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[01]*$`, "Must contain a binary number")
	return e
}

type octEntry struct {
	widget.Entry
}

func NewOctEntry() *octEntry {
	e := &octEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-7]*$`, "Must contain a octal number")
	return e
}

type decEntry struct {
	widget.Entry
}

func NewDecEntry() *decEntry {
	e := &decEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9]*$`, "Must contain a decimal number")
	return e
}

type hexEntry struct {
	widget.Entry
}

func NewHexEntry() *hexEntry {
	e := &hexEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9a-fA-F]*$`, "Must contain a hexadecimal number")
	return e
}

type timestampEntry struct {
	widget.Entry
}

func NewTimestampEntry() *timestampEntry {
	e := &timestampEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9]{1,11}$`, "Must contain a timestamp number")
	return e
}

type yearEntry struct {
	widget.Entry
}

func NewYearEntry() *yearEntry {
	e := &yearEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[1-9][0-9]{3}$`, "Year")
	return e
}

type monthEntry struct {
	widget.Entry
}

func NewMonthEntry() *monthEntry {
	e := &monthEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^(0?[1-9]|1[0-2])$`, "Month")
	return e
}

type dayEntry struct {
	widget.Entry
}

func NewDayEntry() *dayEntry {
	e := &dayEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^(0?[1-9]|[12][0-9]|3[01])$`, "Day")
	return e
}

type hourEntry struct {
	widget.Entry
}

func NewHourEntry() *hourEntry {
	e := &hourEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^([01]?[0-9]|2[0-3])$`, "Hour")
	return e
}

type minuteEntry struct {
	widget.Entry
}

func NewMinuteEntry() *minuteEntry {
	e := &minuteEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-5]?[0-9]$`, "Minute")
	return e
}

type secondEntry struct {
	widget.Entry
}

func NewSecondEntry() *secondEntry {
	e := &secondEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-5]?[0-9]$`, "Second")
	return e
}

type dateTimeEntry struct {
	widget.Entry
}

func NewDateTimeEntry() *dateTimeEntry {
	e := &dateTimeEntry{}
	e.ExtendBaseWidget(e)

	// use strict date time format
	//dateTimeRegexPattern := `^[1-9][0-9]{3}/(0?[1-9]|1[0-2])/(0?[1-9]|[12][0-9]|3[01]) ([01]?[0-9]|2[0-3]):[0-5]?[0-9]:[0-5]?[0-9]$`
	dateTimeRegexPattern := `^[1-9][0-9]{3}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01]) ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`
	e.Validator = validation.NewRegexp(dateTimeRegexPattern, "Must contain a date time string, such as 1970/01/01 00:00:00")
	return e
}
