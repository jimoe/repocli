package arguments

import (
	"fmt"
	"regexp"
)

type str string

func (s *str) String() string {
	return string(*s)
}

func (s *str) subStr(first, last int) string {
	return string([]rune(s.String())[first:last])
}

func (s *str) onlyValidChars(legal string) bool {
	regStr := fmt.Sprintf(`^[%s]+$`, legal)
	return regexp.MustCompile(regStr).MatchString(s.String())
}
