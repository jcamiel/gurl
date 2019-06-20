package parser

func (p *Parser) parseWhitespaces() (Node, error) {

	begin, beginLine := p.Current, p.Line

	whitespaces, err := p.readRunesWhile(func(r rune) bool {
		return r == space || r == tab || r == newLine
	})

	if err != nil || len(whitespaces) == 0 {
		return nil, newSyntaxError(p, "space, tab or newline expected")
	}

	end, endLine := p.Current, p.Line

	beginPos := Position{begin, beginLine}
	endPos := Position{end, endLine}
	return &Whitespaces{beginPos, endPos, string(whitespaces)}, nil
}

func (p *Parser) parseSpaces() (Node, error) {

	begin, beginLine := p.Current, p.Line

	spaces, err := p.readRunesWhile(func(r rune) bool {
		return r == space || r == tab
	})

	if err != nil || len(spaces) == 0 {
		return nil, newSyntaxError(p, "space or tab expected")
	}

	end, endLine := p.Current, p.Line

	beginPos := Position{begin, beginLine}
	endPos := Position{end, endLine}
	return &Spaces{beginPos, endPos, string(spaces)}, nil
}
