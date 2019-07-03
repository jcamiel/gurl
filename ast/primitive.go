package ast

import "fmt"

func (p *Parser) parseWhitespaces() *Whitespaces {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	whitespaces, err := p.readRunesWhile(isWhitespace)
	if err != nil || len(whitespaces) == 0 {
		p.err = newSyntaxError(p, "space, tab or newline is expected at whitespaces beginning")
		return nil
	}

	return &Whitespaces{
		Position{current, line},
		Position{p.current, p.line},
		string(whitespaces),
	}
}

func (p *Parser) parseSpaces() *Spaces {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	spaces, err := p.readRunesWhile(isSpace)
	if err != nil || len(spaces) == 0 {
		p.err = newSyntaxError(p, "space or tab is expected at spaces beginning")
		return nil
	}

	return &Spaces{
		Position{current, line},
		Position{p.current, p.line},
		string(spaces),
	}
}

func (p *Parser) parseComment() *Comment {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	r, err := p.nextRune()
	if p.err = err; err != nil {
		return nil
	}
	if r != hash {
		p.err = newSyntaxError(p, "# is expected at comment beginning")
		return nil
	}
	comment, _ := p.readRunesWhile(isNotNewLine)

	return &Comment{
		Position{current, line},
		Position{p.current, p.line},
		string(comment),
	}
}

func (p *Parser) parseCommentLine() *CommentLine {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	comment := p.parseComment()
	eol := p.parseEol()
	whitespaces := p.tryParseWhitespaces()

	if p.err != nil {
		return nil
	}
	return &CommentLine{
		Position{current, line},
		Position{p.current, p.line},
		comment,
		eol,
		whitespaces,
	}
}

func (p *Parser) parseComments() *Comments {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	var lines []*CommentLine
	for {
		c := p.tryParseCommentLine()
		if c == nil {
			break
		}
		lines = append(lines, c)
	}
	if len(lines) == 0 {
		p.err = newSyntaxError(p, "comments is expected")
		return nil
	}

	return &Comments{
		Position{current, line},
		Position{p.current, p.line},
		lines,
	}
}

func (p *Parser) parseJsonString() *JsonString {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	r, err := p.readRune()
	if p.err = err; err != nil {
		return nil
	}
	if r != '"' {
		p.err = newSyntaxError(p, "\" is expected at json-string beginning")
		return nil
	}
	value := make([]rune, 0)
	for {
		chars, err := p.readRunesWhile(func(r rune) bool {
			return r != '"' && r != '\\' && !isControlCharacter(r)
		})
		if p.err = err; err != nil {
			return nil
		}
		value = append(value, chars...)

		r, err = p.readRune()
		if p.err = err; err != nil {
			return nil
		}
		if isControlCharacter(r) {
			p.err = newSyntaxError(p, "control character not allowed in json-string")
			return nil
		}
		if r == '"' {
			break
		}
		// Parsing of escaped char
		if r == '\\' {
			r, err = p.readRune()
			if p.err = err; err != nil {
				return nil
			}
			if r == 'u' {
				p.err = newSyntaxError(p, "unicode literal not supported in json-string")
				return nil
			}
			controls := map[rune]rune{
				'"':  '"',
				'\\': '\\',
				'/':  '/',
				'b':  '\b',
				'f':  '\f',
				'n':  '\n',
				'r':  '\r',
				't':  '\t',
			}
			c, ok := controls[r]
			if !ok {
				p.err = newSyntaxError(p, "control characted is expected")
				return nil
			}
			value = append(value, c)
		}
	}

	return &JsonString{
		Position{current, line},
		Position{p.current, p.line},
		string(p.buffer[current:p.current]),
		string(value),
	}
}

func (p *Parser) parseKeyString() *KeyString {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	key, err := p.readRunesWhile(func(r rune) bool {
		return !isWhitespace(r) && r != ':' && r != '"' && r != '#'
	})
	if err != nil || len(key) == 0 {
		p.err = newSyntaxError(p, "char is expected at key-string beginning")
		return nil
	}

	return &KeyString{
		Position{current, line},
		Position{p.current, p.line},
		string(key),
	}
}

func (p *Parser) parseKey() *Key {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	var keyString *KeyString
	var jsonString *JsonString

	keyString = p.tryParseKeyString()
	if keyString == nil {
		jsonString = p.parseJsonString()
		if p.err != nil {
			p.err = newSyntaxError(p, "key-string or json-string is expected at key beginning")
			return nil
		}
	}

	var value string
	if keyString != nil {
		value = keyString.Value
	}
	if jsonString != nil {
		value = jsonString.Value
	}
	return &Key{
		Position{current, line},
		Position{p.current, p.line},
		keyString,
		jsonString,
		value,
	}
}

func (p *Parser) parseValue() *Value {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	var valueString *ValueString
	var jsonString *JsonString

	valueString = p.tryParseValueString()
	if valueString == nil {
		jsonString = p.parseJsonString()
		if p.err != nil {
			p.err = newSyntaxError(p, "key-string or json-string is expected at key beginning")
			return nil
		}
	}

	var value string
	if valueString != nil {
		value = valueString.Value
	}
	if jsonString != nil {
		value = jsonString.Value
	}
	return &Value{
		Position{current, line},
		Position{p.current, p.line},
		valueString,
		jsonString,
		value,
	}
}

func (p *Parser) parseColon() *Colon {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	r, err := p.readRune()
	if err != nil || r != ':' {
		p.err = newSyntaxError(p, ": is expected")
		return nil
	}

	return &Colon{
		Position{current, line},
		Position{p.current, p.line},
		":",
	}
}

func (p *Parser) parseKeyValue() *KeyValue {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	comments := p.tryParseComments()
	key := p.parseKey()
	spaces0 := p.tryParseSpaces()
	colon := p.parseColon()
	spaces1 := p.tryParseSpaces()
	value := p.parseValue()
	space2 := p.tryParseSpaces()
	comment := p.tryParseComment()
	eol := p.parseEol()

	if p.err != nil {
		return nil
	}
	return &KeyValue{
		Position{current, line},
		Position{p.current, p.line},
		comments,
		key,
		spaces0,
		colon,
		spaces1,
		value,
		space2,
		comment,
		eol,
	}
}

func (p *Parser) parseValueString() *ValueString {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	value := make([]rune, 0)
	for {
		v, err := p.readRunesWhile(func(r rune) bool {
			return !isWhitespace(r) && r != '#'
		})
		if err != nil {
			break
		}
		value = append(value, v...)
		n, err := p.nextRune()
		if err != nil || isNewLine(n) || n == '#' {
			break
		}
		if !isSpace(n) {
			r, _ := p.readRune()
			value = append(value, r)
			continue
		}

		// Next spaces can be either part of the trailing spaces (with optional comment)
		// or part if the value-string. In the first case, we must not consume it, and
		// in the second we must continue and keep parsing value-string.
		if p.isTrailingSpaces() {
			break
		}
		s, _ := p.readRunesWhile(isSpace)
		value = append(value, s...)
	}

	if len(value) == 0 {
		p.err = newSyntaxError(p, "# or whitespaces is forbidden at value-string beginning")
		return nil
	}
	return &ValueString{
		Position{current, line},
		Position{p.current, p.line},
		string(value),
	}
}

// must start with spaces
func (p *Parser) isTrailingSpaces() bool {
	current := p.current
	_, err := p.readRunesWhile(isSpace)
	if err != nil {
		p.current = current
		return true
	}
	n, err := p.nextRune()
	if err != nil || isNewLine(n) || n == '#' {
		p.current = current
		return true
	}
	p.current = current
	return false
}

func (p *Parser) parseSectionHeader(section string) *SectionHeader {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line
	s := fmt.Sprintf("[%s]", section)
	if !p.tryParseString(s) {
		p.err = newSyntaxError(p, fmt.Sprintf("[%s] is expected", section))
		return nil
	}
	return &SectionHeader{
		Position{current, line},
		Position{p.current, p.line},
		s,
	}
}