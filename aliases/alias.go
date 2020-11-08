package aliases

import (
	"fmt"
	"regexp"
)

type Alias string

const ValidAliasChars = "a-z-"

func (a *Alias) Validate() error {
	regStr := fmt.Sprintf(`^[%s]+$`, ValidAliasChars)
	if regexp.MustCompile(regStr).MatchString(a.String()) {
		return nil
	}
	return fmt.Errorf("illegal character in <aliases> (%s)", ValidAliasChars)
}

func (a *Alias) String() string {
	return string(*a)
}
