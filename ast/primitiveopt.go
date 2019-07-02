package ast

func (p *Parser) tryParseString(s string) bool {
	current, line := p.current, p.line
	runes := []rune(s)
	next, err := p.readRunes(len(runes))
	if err != nil || !equal(runes, next) {
		p.current, p.line = current, line
		return false
	}
	return true
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

func (p *Parser) tryParseSpaces() (*Spaces, error) {
	current, line := p.current, p.line

	node, err := p.parseSpaces()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

func (p *Parser) tryParseCommentLine() (*CommentLine, error) {
	current, line := p.current, p.line

	node, err := p.parseCommentLine()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

func (p *Parser) tryParseKeyString() (*KeyString, error) {
	current, line := p.current, p.line

	node, err := p.parseKeyString()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
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

func (p *Parser) tryParseValueString() (*ValueString, error) {
	current, line := p.current, p.line

	node, err := p.parseValueString()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}
