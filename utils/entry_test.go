package utils

import (
	"fmt"
	"regexp"
	"testing"
)

func TestNewRegexMatch(t *testing.T) {
	regexTpl := `^[0-9]{1,10}$`
	text := `12345678901`
	expression, _ := regexp.Compile(regexTpl)
	ok := expression.MatchString(text)
	fmt.Println("match:", ok)
}

func TestNewRegexMatch2(t *testing.T) {
	regexTpl := `^[1-9][0-9]{3}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01]) ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]$`
	text := "1970/01/02 03:41:55"
	expression, _ := regexp.Compile(regexTpl)
	ok := expression.MatchString(text)
	fmt.Println("match:", ok)
}
