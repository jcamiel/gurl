package ast

func (p *Parser) tryParseHeaders() *Headers {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseHeaders()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseCookies() *Cookies {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseCookies()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseCookie() *Cookie {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseCookie()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseNCookie() []*Cookie {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseNCookie()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseQsParams() *QsParams {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseQsParams()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseFormParams() *FormParams {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseFormParams()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseBody() *Body {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseBody()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

func (p *Parser) tryParseResponse() *Response {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseResponse()
	if p.err != nil {
		p.pos, p.err = pos, nil
		return nil
	}
	return node
}

