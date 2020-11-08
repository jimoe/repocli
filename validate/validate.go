package validate

import (
	"fmt"
	"regexp"
)

const ValidAliasChars = "a-z-"

func Alias(s string) error {
	if regexp.MustCompile(fmt.Sprintf(`^[%s]+$`, ValidAliasChars)).MatchString(s) {
		return nil
	}
	return fmt.Errorf("illegal character in <alias> (%s)", ValidAliasChars)
}
