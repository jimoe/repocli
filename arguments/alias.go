package arguments

import (
	"fmt"
	"regexp"
)

type Alias string

const ValidAliasChars = "a-z-"

func (a *Alias) Validate() error {
	if !onlyValidChars(a.String(), ValidAliasChars) {
		return fmt.Errorf("illegal character in <aliases> (%s)", ValidAliasChars)
	}

	return nil
}

func (a *Alias) String() string {
	return string(*a)
}

func onlyValidChars(s, legal string) bool {
	regStr := fmt.Sprintf(`^[%s]+$`, legal)
	return regexp.MustCompile(regStr).MatchString(s)
}
