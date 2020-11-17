package arguments

import (
	"fmt"
)

type Alias struct {
	str
}

const ValidAliasChars = "a-z-"

func NewAlias(s string) *Alias {
	return &Alias{str(s)}
}

func (a *Alias) Validate() error {
	if !a.onlyValidChars(ValidAliasChars) {
		return fmt.Errorf("illegal character in <aliases> (%s)", ValidAliasChars)
	}

	return nil
}
