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
	// alias "." is allowed for open current directory
	if !a.onlyValidChars(ValidAliasChars) && a.String() != "." {
		return fmt.Errorf("illegal character in <alias> (%s)", ValidAliasChars)
	}

	return nil
}
