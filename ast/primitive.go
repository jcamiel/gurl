package ast

func (p *Parser) parseWhitespaces() (*Whitespaces, error) {

	current, line := p.current, p.line

	whitespaces, err := p.readRunesWhile(func(r rune) bool {
		return IsSpace(r) || IsNewline(r)
	})

	if err != nil || len(whitespaces) == 0 {
		return nil, newSyntaxError(p, "space, tab or newline is expected")
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
		return IsSpace(r)
	})

	if err != nil || len(spaces) == 0 {
		return nil, newSyntaxError(p, "space or tab is expected")
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
		return nil, newSyntaxError(p, "# is expected at the beginning of a comment")
	}

	comment, err := p.readRunesWhile(func(r rune) bool {
		return !IsNewline(r)
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

func (p *Parser) tryParseComments() (*Comments, error) {

	current, line := p.current, p.line

	node, err := p.parseComments()

	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}
