package ast

func (p *Parser) tryParseString(s string) bool {
	if p.err != nil {
		return false
	}
	pos := p.pos
	runes := []rune(s)
	next, err := p.readRunes(len(runes))
	if err != nil || !equal(runes, next) {
		p.rewindTo(pos)
		return false
	}
	return true
}

func (p *Parser) tryParseWhitespaces() *Whitespaces {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseWhitespaces()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseSpaces() *Spaces {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseSpaces()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseComment() *Comment {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseComment()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseCommentLine() *CommentLine {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseCommentLine()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseComments() *Comments {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseComments()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseKeyString() *KeyString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseKeyString()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseKeyValue() *KeyValue {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseKeyValue()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseValueString() *ValueString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseValueString()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseJson() (value Json, text string) {
	if p.err != nil {
		return nil, ""
	}
	pos := p.pos

	value, text = p.parseJson()
	if p.err != nil {
		p.rewindTo(pos)
		return nil, ""
	}
	return
}

func (p *Parser) tryParseXml() string {
	if p.err != nil {
		return ""
	}
	pos := p.pos

	text := p.parseXml()
	if p.err != nil {
		p.rewindTo(pos)
		return ""
	}
	return text
}

func (p *Parser) tryParseBase64() (value []byte, text string) {
	if p.err != nil {
		return nil, ""
	}
	pos := p.pos

	value, text = p.parseBase64()
	if p.err != nil {
		p.rewindTo(pos)
		return nil, ""
	}
	return
}

func (p *Parser) tryParseQueryString() *QueryString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseQueryString()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseInteger() *Integer {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseInteger()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseFloat() *Float {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseFloat()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseBool() *Bool {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseBool()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}