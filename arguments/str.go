package arguments

import (
	"fmt"
	"regexp"
)

type commonStr string

func (s *commonStr) String() string {
	return string(*s)
}

func (s *commonStr) subStr(first, last int) string {
	return string([]rune(s.String())[first:last])
}

func (s *commonStr) onlyValidChars(legal string) bool {
	regStr := fmt.Sprintf(`^[%s]+$`, legal)
	return regexp.MustCompile(regStr).MatchString(s.String())
}
