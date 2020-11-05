package arguments

type Arglist []string

func (a Arglist) Contains(argument string) bool {
	for _, s := range a {
		if s == argument {
			return true
		}
	}
	return false
}
