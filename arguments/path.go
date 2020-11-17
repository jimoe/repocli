package arguments

import (
	"errors"
	"fmt"
)

type Path string
type SubPath struct {
	Path
}

const ValidPathChars = "a-zA-Z0-9-_/."

func (p *Path) Validate() error {
	if !onlyValidChars(p.String(), ValidPathChars) {
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

func (sp *SubPath) Validate() error {
	if !onlyValidChars(sp.String(), ValidPathChars) {
		return fmt.Errorf("illegal character in subpath (%s)", ValidPathChars)
	}

	if sp.subStr(0, 1) == "/" {
		return errors.New("first char in a SubPath should not be '/'")
	}

	length := len(sp.String())
	if sp.subStr(length-1, length) == "/" {
		return errors.New("SubPath should not end with '/'")
	}

	return nil
}

func (p *Path) String() string {
	return string(*p)
}

func (p *Path) subStr(first, last int) string {
	return string([]rune(p.String())[first:last])
}
