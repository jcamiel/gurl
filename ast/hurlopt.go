package ast

func (p *Parser) tryParseHeaders() *Headers {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseHeaders()
	if p.err != nil {
		p.rewindTo(pos)
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
		p.rewindTo(pos)
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
		p.rewindTo(pos)
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
		p.rewindTo(pos)
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
		p.rewindTo(pos)
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
		p.rewindTo(pos)
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
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseCapture() *Capture {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseCapture()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseCaptures() *Captures {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseCaptures()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseAssert() *Assert {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseAssert()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node

}

func (p *Parser) tryParseAsserts() *Asserts {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseAsserts()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseNEntry() []*Entry {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseNEntry()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}

func (p *Parser) tryParseEntry() *Entry {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	node := p.parseEntry()
	if p.err != nil {
		p.rewindTo(pos)
		return nil
	}
	return node
}