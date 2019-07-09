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

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAsciiLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isCombining(r rune) bool {
	return (r >= '\u0300' && r <= '\u036f') ||
		(r >= '\u1ab0' && r <= '\u1aff') ||
		(r >= '\u1dc0' && r <= '\u1dff') ||
		(r >= '\ufe20' && r <= '\ufe2f')
}

func isControlCharacter(r rune) bool {
	return r == '\b' || r == '\f' || r == '\n' || r == '\r' || r == '\t'
}

// getu4 decodes \uXXXX from s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	var r rune
	for _, c := range s[:4] {
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}
		r = r*16 + rune(c)
	}
	return r
}