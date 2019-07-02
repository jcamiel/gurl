package ast

func (p *Parser) tryParseHeaders() (*Headers, error) {
	current, line := p.current, p.line

	node, err := p.parseHeaders()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}

func (p *Parser) tryParseCookies() (*Cookies, error) {
	current, line := p.current, p.line

	node, err := p.parseCookies()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}

func (p *Parser) tryParseCookie() (*Cookie, error) {
	current, line := p.current, p.line

	node, err := p.parseCookie()
	if err != nil {
		p.current, p.line = current, line
		return nil, err
	}

	return node, nil
}
