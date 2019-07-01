package ast

const (
	tab            = '\u0009'
	lineFeed       = '\u000a'
	carriageReturn = '\u000d'
	space          = '\u0020'
	quotationMark  = '\u0022'
	hash           = '\u0023'
	reverseSolidus = '\u005c'
)

func equal(a, b []rune) bool {
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

func isNewLine(r rune) bool {
	return r == '\n' || r == '\r'
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isWhitespace(r rune) bool {
	return isNewLine(r) || isSpace(r)
}

func isControlCharacter(r rune) bool {
	return r == '\b' || r == '\f' || r == '\n' || r == '\r' || r == '\t'
}
