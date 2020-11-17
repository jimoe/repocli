package arguments

import (
	"errors"
	"fmt"
)

type SubPath struct {
	str
}

func NewSubPath(s string) *SubPath {
	return &SubPath{str(s)}
}

func (sp *SubPath) Validate() error {
	if !sp.onlyValidChars(ValidPathChars) {
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
