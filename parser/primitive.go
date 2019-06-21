package parser

func (p *Parser) parseWhitespaces() (*Whitespaces, error) {

	begin, beginLine := p.Current, p.Line

	whitespaces, err := p.readRunesWhile(func(r rune) bool {
		return r == space || r == tab || r == newLine
	})

	if err != nil || len(whitespaces) == 0 {
		return nil, newSyntaxError(p, "space, tab or newline expected")
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
		return r == space || r == tab
	})

	if err != nil || len(spaces) == 0 {
		return nil, newSyntaxError(p, "space or tab expected")
	}

	end, endLine := p.Current, p.Line

	return &Spaces{
		Position{begin, beginLine},
		Position{end, endLine},
		string(spaces)}, nil
}
