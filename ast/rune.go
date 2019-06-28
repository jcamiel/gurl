package ast

const (
	tab            = '\u0009'
	lineFeed       = '\u000a'
	carriageReturn = '\u000d'
	space          = '\u0020'
	quote          = '\u0022'
	hash           = '\u0023'
	reverseSolidus = '\u005c'
)

func Equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func RuneInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsNewline(r rune) bool {
	return r == lineFeed || r == carriageReturn
}

func IsSpace(r rune) bool {
	return r == space || r == tab
}
