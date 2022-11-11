package arguments

import (
	"fmt"
)

type Alias struct {
	commonStr
}

const ValidAliasChars = "a-z0-9-"

func NewAlias(s string) *Alias {
	return &Alias{commonStr(s)}
}

func (a *Alias) Validate() error {
	if !a.onlyValidChars(ValidAliasChars) {
		return fmt.Errorf("illegal character in <alias> (%s)", ValidAliasChars)
	}

	return nil
}
