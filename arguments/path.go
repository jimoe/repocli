package arguments

import (
	"errors"
	"fmt"
)

type Path struct {
	commonStr
}

const ValidPathChars = "a-zA-Z0-9-_/."

func NewPath(s string) *Path {
	return &Path{commonStr(s)}
}

func (p *Path) Validate() error {
	if !p.onlyValidChars(ValidPathChars) {
		return fmt.Errorf("illegal character in path (%s)", ValidPathChars)
	}

	if p.subStr(0, 1) != "/" {
		return errors.New("first char in path must be '/'")
	}

	length := len(p.String())
	if p.subStr(length-1, length) == "/" {
		return errors.New("the path should not end with '/'")
	}

	return nil
}
