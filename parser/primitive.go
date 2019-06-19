package parser

func (p *Parser) parseWhitespaces() (Node, error) {

	begin, beginLine := p.Current, p.Line
	end, endLine := begin, beginLine

	for {
		r, err := p.nextRune()
		if err != nil {
			break
		}
		if isWhiteSpace(r) || isNewLine(r) {
			_, _ = p.readRune()
			end, endLine = p.Current, p.Line
		} else {
			break
		}
	}

	if begin == end {
		return nil, nil
	}
	beginPos := Position{begin, beginLine}
	endPos := Position{end, endLine}
	whitespaces := string(p.Buffer[begin:end])
	return &Whitespaces{beginPos, endPos, whitespaces}, nil
}

func (p *Parser) parseSpaces() (Node, error) {

	begin, beginLine := p.Current, p.Line
	end, endLine := begin, beginLine

	for {
		r, err := p.nextRune()
		if err != nil {
			break
		}
		if isWhiteSpace(r) {
			_, _ = p.readRune()
			end, endLine = p.Current, p.Line
		} else {
			break
		}
	}

	if begin == end {
		return nil, nil
	}
	beginPos := Position{begin, beginLine}
	endPos := Position{end, endLine}
	spaces := string(p.Buffer[begin:end])
	return &Spaces{beginPos, endPos, spaces}, nil
}