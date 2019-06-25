package parser

func (p *Parser) parseWhitespaces() (*Whitespaces, error) {

	begin, beginLine := p.Current, p.Line

	whitespaces, err := p.readRunesWhile(func(r rune) bool {
		return IsSpace(r) || IsNewline(r)
	})

	if err != nil || len(whitespaces) == 0 {
		return nil, newSyntaxError(p, "space, tab or newline is expected")
	}

	end, endLine := p.Current, p.Line

	return &Whitespaces{
		Position{begin, beginLine},
		Position{end, endLine},
		string(whitespaces)}, nil
}

func (p *Parser) tryParseWhitespaces() (*Whitespaces, error) {

	begin, beginLine := p.Current, p.Line

	node, err := p.parseWhitespaces()

	if err != nil {
		p.Current, p.Line = begin, beginLine
		return nil, err
	}

	return node, nil
}

func (p *Parser) parseSpaces() (*Spaces, error) {

	begin, beginLine := p.Current, p.Line

	spaces, err := p.readRunesWhile(func(r rune) bool {
		return IsSpace(r)
	})

	if err != nil || len(spaces) == 0 {
		return nil, newSyntaxError(p, "space or tab is expected")
	}

	end, endLine := p.Current, p.Line

	return &Spaces{
		Position{begin, beginLine},
		Position{end, endLine},
		string(spaces)}, nil
}

func (p *Parser) tryParseSpaces() (*Spaces, error) {

	begin, beginLine := p.Current, p.Line

	node, err := p.parseSpaces()

	if err != nil {
		p.Current, p.Line = begin, beginLine
		return nil, err
	}

	return node, nil
}

func (p *Parser) parseComment() (*Comment, error) {

	begin, beginLine := p.Current, p.Line

	r, err := p.nextRune()
	if err != nil {
		return nil, err
	}

	if r != hash {
		return nil, newSyntaxError(p, "# is expected at the beginning of a comment")
	}

	comment, err := p.readRunesWhile(func(r rune) bool {
		return !IsNewline(r)
	})

	end, endLine := p.Current, p.Line

	return &Comment{
		Position{begin, beginLine},
		Position{end, endLine},
		string(comment)}, nil

}

func (p *Parser) tryParseComment() (*Comment, error) {
	begin, beginLine := p.Current, p.Line

	node, err := p.parseComment()

	if err != nil {
		p.Current, p.Line = begin, beginLine
		return nil, err
	}

	return node, nil
}