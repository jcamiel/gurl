package ast

import "fmt"

func (p *Parser) parseWhitespaces() (*Whitespaces, error) {
	current, line := p.current, p.line

	whitespaces, err := p.readRunesWhile(func(r rune) bool {
		return isWhitespace(r)
	})
	if err != nil || len(whitespaces) == 0 {
		return nil, newSyntaxError(p, "space, tab or newline is expected at whitespaces beginning")
	}

	return &Whitespaces{
		Position{current, line},
		Position{p.current, p.line},
		string(whitespaces),
	}, nil
}

func (p *Parser) parseSpaces() (*Spaces, error) {
	current, line := p.current, p.line

	spaces, err := p.readRunesWhile(func(r rune) bool {
		return isSpace(r)
	})
	if err != nil || len(spaces) == 0 {
		return nil, newSyntaxError(p, "space or tab is expected at spaces beginning")
	}

	return &Spaces{
		Position{current, line},
		Position{p.current, p.line},
		string(spaces),
	}, nil
}

func (p *Parser) parseComment() (*Comment, error) {
	current, line := p.current, p.line

	r, err := p.nextRune()
	if err != nil {
		return nil, err
	}
	if r != hash {
		return nil, newSyntaxError(p, "# is expected at comment beginning")
	}
	comment, err := p.readRunesWhile(func(r rune) bool {
		return !isNewLine(r)
	})

	return &Comment{
		Position{current, line},
		Position{p.current, p.line},
		string(comment),
	}, nil
}

func (p *Parser) parseCommentLine() (*CommentLine, error) {
	current, line := p.current, p.line

	comment, err := p.parseComment()
	if err != nil {
		return nil, err
	}
	eol, err := p.parseEol()
	if err != nil {
		return nil, err
	}
	whitespaces, _ := p.tryParseWhitespaces()

	return &CommentLine{
		Position{current, line},
		Position{p.current, p.line},
		comment,
		eol,
		whitespaces,
	}, nil

}

func (p *Parser) parseComments() (*Comments, error) {
	current, line := p.current, p.line

	var comments []*CommentLine
	for {
		c, err := p.tryParseCommentLine()
		if err != nil {
			break
		}
		comments = append(comments, c)
	}
	if len(comments) == 0 {
		return nil, newSyntaxError(p, "comments is expected")
	}

	return &Comments{
		Position{current, line},
		Position{p.current, p.line},
		comments,
	}, nil

}

func (p *Parser) parseJsonString() (*JsonString, error) {

	current, line := p.current, p.line

	r, err := p.readRune()
	if err != nil {
		return nil, err
	}
	if r != '"' {
		return nil, newSyntaxError(p, "\" is expected at json-string beginning")
	}
	value := make([]rune, 0)
	for {
		chars, err := p.readRunesWhile(func(r rune) bool {
			return r != '"' && r != '\\' && !isControlCharacter(r)
		})
		if err != nil {
			return nil, err
		}
		value = append(value, chars...)

		r, err = p.readRune()
		if err != nil {
			return nil, err
		}
		if isControlCharacter(r) {
			return nil, newSyntaxError(p, "control character not allowed in json-string")
		}
		if r == '"' {
			break
		}
		// Parsing of escaped char
		if r == '\\' {
			r, err = p.readRune()
			if err != nil {
				return nil, err
			}

			if r == 'u' {
				return nil, newSyntaxError(p, "unicode literal not supported")
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
				return nil, newSyntaxError(p, "control characted is expected")
			}
			value = append(value, c)
		}
	}

	return &JsonString{
		Position{current, line},
		Position{p.current, p.line},
		string(p.buffer[current:p.current]),
		string(value),
	}, nil
}

func (p *Parser) parseKeyString() (*KeyString, error) {
	current, line := p.current, p.line

	key, err := p.readRunesWhile(func(r rune) bool {
		return !isWhitespace(r) && r != ':' && r != '"' && r != '#'
	})
	if err != nil || len(key) == 0 {
		return nil, newSyntaxError(p, "char is expected at key-string beginning")
	}

	return &KeyString{
		Position{current, line},
		Position{p.current, p.line},
		string(key),
	}, nil
}

func (p *Parser) parseKey() (*Key, error) {
	current, line := p.current, p.line

	var keyString *KeyString
	var jsonString *JsonString

	keyString, err := p.tryParseKeyString()
	if err != nil {
		jsonString, err = p.parseJsonString()
		if err != nil {
			return nil, newSyntaxError(p, "key-string or json-string is expected at key beginning")
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
	}, nil
}

func (p *Parser) parseValue() (*Value, error) {
	current, line := p.current, p.line

	var valueString *ValueString
	var jsonString *JsonString

	valueString, err := p.tryParseValueString()
	if err != nil {
		jsonString, err = p.parseJsonString()
		if err != nil {
			return nil, newSyntaxError(p, "key-string or json-string is expected at key beginning")
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
	}, nil
}

func (p *Parser) parseColon() (*Colon, error) {
	current, line := p.current, p.line

	r, err := p.readRune()
	if err != nil || r != ':' {
		return nil, newSyntaxError(p, ": is expected")
	}

	return &Colon{
		Position{current, line},
		Position{p.current, p.line},
		":",
	}, nil
}

func (p *Parser) parseKeyValue() (*KeyValue, error) {
	current, line := p.current, p.line

	comments, _ := p.tryParseComments()
	key, err := p.parseKey()
	if err != nil {
		return nil, err
	}
	spaces0, _ := p.tryParseSpaces()
	colon, err := p.parseColon()
	if err != nil {
		return nil, err
	}
	spaces1, _ := p.tryParseSpaces()
	value, err := p.parseValue()
	if err != nil {
		return nil, err
	}
	space2, _ := p.tryParseSpaces()
	comment, _ := p.tryParseComment()
	eol, err := p.parseEol()
	if err != nil {
		return nil, err
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
	}, nil
}

func (p *Parser) parseValueString() (*ValueString, error) {
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

		// if we have trailing spaces, we must break
		s, err := p.readRunesWhile(func(r rune) bool {
			return isSpace(r)
		})
		if err != nil {
			break
		}
		n, err = p.nextRune()
		if err != nil || isNewLine(n) || n == '#' {
			break
		}
		value = append(value, s...)
	}

	if len(value) == 0 {
		return nil, newSyntaxError(p, "# or whitespaces is forbidden at value-string beginning")
	}
	return &ValueString{
		Position{current, line},
		Position{p.current, p.line},
		string(value),
	}, nil
}

func (p *Parser) parseSectionHeader(section string) (*SectionHeader, error) {
	current, line := p.current, p.line
	s := fmt.Sprintf("[%s]", section)
	if !p.tryParseString(s) {
		return nil, newSyntaxError(p, fmt.Sprintf("[%s] is expected", section))
	}
	return &SectionHeader{
		Position{current, line},
		Position{p.current, p.line},
		s,
	}, nil
}