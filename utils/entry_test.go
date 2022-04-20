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
