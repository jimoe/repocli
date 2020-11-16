package arguments

import (
	"errors"
	"fmt"
	"regexp"
)

type Path string

const ValidPathChars = "a-zA-Z0-9-_/."

func (p *Path) Validate() error {
	regStr := fmt.Sprintf(`^[%s]+$`, ValidPathChars)
	if !regexp.MustCompile(regStr).MatchString(p.String()) {
		return fmt.Errorf("illegal character in <path> (%s)", ValidPathChars)
	}

	if p.subStr(0, 1) != "/" {
		return errors.New("first char in <path> must be '/'")
	}

	length := len(p.String())
	if p.subStr(length-1, length) == "/" {
		return errors.New("the <path> should not end with '/'")
	}

	return nil
}

func (p *Path) String() string {
	return string(*p)
}

func (p *Path) subStr(first, last int) string {
	return string([]rune(p.String())[first:last])
}
