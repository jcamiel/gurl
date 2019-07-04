package ast

func (p *Parser) tryParseString(s string) bool {
	if p.err != nil {
		return false
	}
	current, line := p.current, p.line
	runes := []rune(s)
	next, err := p.readRunes(len(runes))
	if err != nil || !equal(runes, next) {
		p.current, p.line = current, line
		p.err = nil
		return false
	}
	return true
}

func (p *Parser) tryParseWhitespaces() *Whitespaces {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseWhitespaces()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseSpaces() *Spaces {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseSpaces()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseComment() *Comment {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseComment()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseCommentLine() *CommentLine {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseCommentLine()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseComments() *Comments {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseComments()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseKeyString() *KeyString {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseKeyString()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseKeyValue() *KeyValue {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseKeyValue()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseValueString() *ValueString {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseValueString()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}

func (p *Parser) tryParseNKeyValue() []*KeyValue {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	node := p.parseNKeyValue()
	if p.err != nil {
		p.current, p.line = current, line
		p.err = nil
		return nil
	}
	return node
}
