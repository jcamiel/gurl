package parser

const (
	space   = '\u0020'
	tab     = '\u0009'
	newLine = '\u000a'
	quote   = '\u0022'
	hash    = '\u0023'
)

func isWhiteSpace(r rune) bool {
	return r == space || r == tab
}

func isNewLine(r rune) bool {
	return r == newLine
}

func isHash(r rune) bool {
	return r == hash
}

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
