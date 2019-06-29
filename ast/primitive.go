package ast

func (p *Parser) parseWhitespaces() (*Whitespaces, error) {
	current, line := p.current, p.line

	whitespaces, err := p.readRunesWhile(func(r rune) bool {
		return isSpace(r) || isNewLine(r)
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

func (p *Parser) tryParseWhitespaces() (*Whitespaces, error) {
	current, line := p.current, p.line

	node, err := p.parseWhitespaces()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

func (p *Parser) tryParseSpaces() (*Spaces, error) {
	current, line := p.current, p.line

	node, err := p.parseSpaces()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

func (p *Parser) tryParseComment() (*Comment, error) {
	current, line := p.current, p.line

	node, err := p.parseComment()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

func (p *Parser) tryParseCommentLine() (*CommentLine, error) {
	current, line := p.current, p.line

	node, err := p.parseCommentLine()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}

func (p *Parser) parseComments() (*Comments, error) {

	current, line := p.current, p.line

	var comments []*CommentLine
	for {
		c, err := p.tryParseCommentLine()
		if err != nil { break }
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

func (p *Parser) tryParseComments() (*Comments, error) {
	current, line := p.current, p.line

	node, err := p.parseComments()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

	value := make([]rune,0)
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
		string(p.buffer[current: p.current]),
		string(value),
	}, nil
}

func (p *Parser) parseKeyString() (*KeyString, error) {
	current, line := p.current, p.line

	key, err := p.readRunesWhile(func(r rune) bool {
		return !isNewLine(r) && r != ':' && r != '"' && r != '#'
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

func (p *Parser) tryParseKeyString() (*KeyString, error) {
	current, line := p.current, p.line

	node, err := p.parseKeyString()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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
		value = keyString.Text
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
		},nil
}

func (p *Parser) parseValue() (*Value, error) {
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
		value = keyString.Text
	}
	if jsonString != nil {
		value = jsonString.Value
	}

	return &Value{
		Position{current, line},
		Position{p.current, p.line},
		keyString,
		jsonString,
		value,
	},nil
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

func (p *Parser) tryParseKeyValue() (*KeyValue, error) {
	current, line := p.current, p.line

	node, err := p.parseKeyValue()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}