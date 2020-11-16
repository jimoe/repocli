package arguments

import (
	"fmt"
	"regexp"
)

type Path string

const ValidPathChars = "a-zA-Z0-9-_/."

func (p *Path) Validate() error {
	regStr := fmt.Sprintf(`^[%s]+$`, ValidPathChars)
	if regexp.MustCompile(regStr).MatchString(p.String()) {
		return nil
	}
	return fmt.Errorf("illegal character in <path> (%s)", ValidPathChars)
}

func (p *Path) String() string {
	return string(*p)
}
