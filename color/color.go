package color

import (
	"fmt"
	"strings"
)

const format = "\x1b[%sm%s\x1b[0m"

type Color string

const (
	Blue Color = "1;34"
	Red Color = "1;31"
)

func (c Color) Printf(str string, subStrings ...interface{}) {
	words := strings.Split(str, " ")
	for i, word := range words {
		if word[:1] == "%" || word[:2] == "'%" {
			continue
		}
		words[i] = fmt.Sprintf(format, c, word)
	}
	str = strings.Join(words, " ")

	for i, subStr := range subStrings {
		if str, ok := subStr.(string); ok {
			if len(str) > 0 && str[:1] != "\x1b" {
				subStrings[i] = fmt.Sprintf(format, c, subStr)
			}
		}
	}

	fmt.Printf(str, subStrings...)
}

func (c Color) Println(parts ...interface{}) {
	for i, part := range parts {
		if str, ok := part.(string); ok {
			if len(str) > 0 && str[:1] != "\x1b" {
				parts[i] = fmt.Sprintf(format, c, part)
			}
		}
	}
	fmt.Println(parts...)
}

func Printf(str string, subStrings ...interface{}) {
	Blue.Printf(str, subStrings...)
}

func Println(parts ...interface{}) {
	Blue.Println(parts...)
}

func Highlight(str string) string {
	return fmt.Sprintf("\x1b[0;33m%s\x1b[0m", str)
}

func Error(str string) string {
	return fmt.Sprintf("\x1b[1;31m%s\x1b[0m", str)
}
