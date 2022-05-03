package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

type binEntry struct {
	widget.Entry
}

func NewBinEntry() *binEntry {
	e := &binEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapBreak
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[01]*$`, "Must be a binary number")
	return e
}

type octEntry struct {
	widget.Entry
}

func NewOctEntry() *octEntry {
	e := &octEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapBreak
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-7]*$`, "Must be a octal number")
	return e
}

type decEntry struct {
	widget.Entry
}

func NewDecEntry() *decEntry {
	e := &decEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapBreak
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9]*$`, "Must be a decimal number")
	return e
}

type hexEntry struct {
	widget.Entry
}

func NewHexEntry() *hexEntry {
	e := &hexEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapBreak
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9a-fA-F]*$`, "Must be a hexadecimal number")
	return e
}

type timestampEntry struct {
	widget.Entry
}

func NewTimestampEntry() *timestampEntry {
	e := &timestampEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9]{1,11}$`, "Must be a timestamp number")
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
	e.Validator = validation.NewRegexp(dateTimeRegexPattern, "Must be a date time string, such as 1970/01/01 00:00:00")
	return e
}

type caseOriginEntry struct {
	widget.Entry
}

func NewCaseOriginEntry() *caseOriginEntry {
	e := &caseOriginEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapBreak
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^[0-9a-zA-Z_ ]*$`, "Must be a string contains numbers/alphabets/underscore character/space character")
	return e
}

type caseResultEntry struct {
	widget.Entry
}

func NewCaseResultEntry() *caseResultEntry {
	e := &caseResultEntry{}
	e.MultiLine = true
	e.Wrapping = fyne.TextWrapBreak
	e.ExtendBaseWidget(e)
	return e
}

type portEntry struct {
	widget.Entry
}

func NewPortEntry() *portEntry {
	e := &portEntry{}
	e.ExtendBaseWidget(e)

	// port allowed section: [30000, 59999]
	//e.Validator = validation.NewRegexp(`^([0-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`, "Choose port not in use from 1 to 65535")
	e.Validator = validation.NewRegexp(`^[345][0-9]{4}$`, "Choose port not in use from 30000 to 59999")
	return e
}

func NewIPv4Entry() *caseResultEntry {
	e := &caseResultEntry{}
	e.MultiLine = false
	e.Wrapping = fyne.TextWrapOff
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`^((([0-9])|([1-9]\d)|(1\d{2})|(2[0-4]\d)|(25[0-5]))\.){3}(([0-9])|([1-9]\d)|(1\d{2})|(2[0-4]\d)|(25[0-5]))$`, "Must be a IP V4 address")
	return e
}

func NewMaskBitEntry() *caseResultEntry {
	e := &caseResultEntry{}
	e.MultiLine = false
	e.Wrapping = fyne.TextWrapOff
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp("^([0-9]|[12][0-9]|30)$", "Must be a MaskBit number from 0 to 30")
	return e
}

func NewNumberEntry() *caseResultEntry {
	e := &caseResultEntry{}
	e.MultiLine = false
	e.Wrapping = fyne.TextWrapOff
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp("^([0-9]|[1-9][0-9]{1,})$", "Must be a number value")
	return e
}

type charExcludeEntry struct {
	widget.Entry
}

func NewCharExcludeEntry() *charExcludeEntry {
	e := &charExcludeEntry{}
	e.MultiLine = false
	e.Wrapping = fyne.TextWrapOff
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp("^([0-9a-zA-Z]|[!@#$%^&*])*$", "Characters must be in 0-9a-zA-Z or in !@#$%^&*")
	return e
}

type passLengthEntry struct {
	widget.Entry
}

func NewPassLengthEntry() *passLengthEntry {
	e := &passLengthEntry{}
	e.MultiLine = false
	e.Wrapping = fyne.TextWrapOff
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp("^([1-9]|[1-9][0-9]|100)$", "Password length must be 1-100")
	return e
}

type passCountEntry struct {
	widget.Entry
}

func NewPassCountEntry() *passCountEntry {
	e := &passCountEntry{}
	e.MultiLine = false
	e.Wrapping = fyne.TextWrapOff
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp("^([1-9]|[1-9][0-9]{1,2}|1000)$", "Password count must be 1-1000")
	return e
}
