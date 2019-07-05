package ast

const (
	hash           = '\u0023'
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

func isNotNewLine(r rune) bool {
	return !isNewLine(r)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isWhitespace(r rune) bool {
	return isNewLine(r) || isSpace(r)
}

func isCombining(r rune) bool {
	return (r >= '\u0300' && r <= '\u036f') ||
		(r >= '\u1ab0' && r <= '\u1aff') ||
		(r >= '\u1dc0' && r <= '\u1dff') ||
		(r >= '\ufe20' && r <= '\ufe2f')
}