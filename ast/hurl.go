package ast

import (
	"io"
)

func (p *Parser) parseHurlFile() *HurlFile {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	whitespaces := p.tryParseWhitespaces()
	entries := p.parseNEntry()

	if p.err != nil {
		return nil
	}
	return &HurlFile{
		Position{current, line},
		Position{p.current, p.line},
		whitespaces,
		entries,
	}
}

func (p *Parser) parseNEntry() []*Entry {
	if p.err != nil {
		return nil
	}

	entries := make([]*Entry, 0)
	for {
		e := p.parseEntry()
		if p.err != nil {
			// FIXME: for the moment, silent fail on error.
			p.err = nil
			if p.hasRuneToRead() {
				p.skipToNextEol()
				continue
			}
			break
		}
		entries = append(entries, e)
	}
	return entries
}

func (p *Parser) parseEntry() *Entry {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	request := p.parseRequest()

	if p.err != nil {
		return nil
	}
	return &Entry{
		Position{current, line},
		Position{p.current, p.line},
		request,
	}
}

func (p *Parser) parseRequest() *Request {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	comments := p.tryParseComments()
	method := p.parseMethod()
	spaces0 := p.parseSpaces()
	url := p.parseUrl()
	spaces1 := p.tryParseSpaces()
	comment := p.tryParseComment()
	eol := p.parseEol()
	headers := p.tryParseHeaders()
	cookies := p.tryParseCookies()
	qsparams := p.tryParseQsParams()

	if p.err != nil {
		return nil
	}
	return &Request{
		Position{current, line},
		Position{p.current, p.line},
		comments,
		method,
		spaces0,
		url,
		spaces1,
		comment,
		eol,
		headers,
		cookies,
		qsparams,
	}
}

func (p *Parser) parseMethod() *Method {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	methods := []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
		"PATCH",
	}
	for _, method := range methods {
		if p.tryParseString(method) {
			return &Method{
				Position{current, line},
				Position{p.current, p.line},
				method,
			}
		}
	}
	p.err = newSyntaxError(p, "method is expected")
	return nil
}

func (p *Parser) parseUrl() *Url {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	isGenDelims := func(r rune) bool {
		return r == ':' || r == '/' || r == '?' || r == '#' || r == '[' || r == ']' || r == '@'
	}
	isSubDelims := func(r rune) bool {
		return r == '!' || r == '$' || r == '&' || r == '\'' || r == '(' || r == ')' ||
			r == '*' || r == '+' || r == ',' || r == ';' || r == '='
	}
	url, err := p.readRunesWhile(func(r rune) bool {
		isAlpha := (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
		isDigit := r >= '0' && r <= '9'
		isUnreserved := isAlpha || isDigit || r == '-' || r == '.' || r == '_' || r == '~'
		isReserved := isGenDelims(r) || isSubDelims(r)
		isHurlSpecific := r == '{' || r == '}'
		return isReserved || isUnreserved || isHurlSpecific
	})
	if err != nil || len(url) == 0 {
		p.err = newSyntaxError(p, "url is expected")
		return nil
	}

	return &Url{
		Position{current, line},
		Position{p.current, p.line},
		string(url),
	}
}

func (p *Parser) parseEol() *Eol {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	eol, err := p.readRunesWhile(isNewLine)
	if err != nil && err != io.EOF {
		p.err = newSyntaxError(p, "newline is expected")
		return nil
	}

	if err != io.EOF {
		if len(eol) == 0 {
			p.err = newSyntaxError(p, "newline is expected")
			return nil
		}
		_, _ = p.readRunesWhile(isWhitespace)
	}

	return &Eol{
		Position{current, line},
		Position{p.current, p.line},
		string(p.buffer[current:p.current]),
	}
}

func (p *Parser) parseHeaders() *Headers {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	headers := p.parseNKeyValue()
	if p.err != nil {
		return nil
	}
	return &Headers{
		Position{current, line},
		Position{p.current, p.line},
		headers,
	}
}

func (p *Parser) parseCookieValue() *CookieValue {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	cookie, err := p.readRunesWhile(func(r rune) bool {
		return (r >= 'A' && r <= 'Z') ||
			(r >= 'a' && r <= 'z') ||
			(r >= '0' && r <= '9') ||
			r == ':' ||
			r == '/' ||
			r == '%'
	})
	if err != nil {
		p.err = newSyntaxError(p, "[A-Za-z0-9:/%] char is expected for cookie-value")
		return nil
	}

	return &CookieValue{
		Position{current, line},
		Position{p.current, p.line},
		string(cookie),
	}
}

func (p *Parser) parseCookie() *Cookie {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	comments := p.tryParseComments()
	key := p.parseKey()
	spaces0 := p.tryParseSpaces()
	colon := p.parseColon()
	spaces1 := p.tryParseSpaces()
	cookieValue := p.parseCookieValue()
	spaces2 := p.tryParseSpaces()
	comment := p.tryParseComment()
	eol := p.parseEol()

	if p.err != nil {
		return nil
	}
	return &Cookie{
		Position{current, line},
		Position{p.current, p.line},
		comments,
		key,
		spaces0,
		colon,
		spaces1,
		cookieValue,
		spaces2,
		comment,
		eol,
	}
}

func (p *Parser) parseNCookie() []*Cookie {
	if p.err != nil {
		return nil
	}
	cookies := make([]*Cookie, 0)
	for {
		c := p.tryParseCookie()
		if c == nil {
			break
		}
		cookies = append(cookies, c)
	}
	if len(cookies) == 0 {
		p.err = newSyntaxError(p, "At least one comment-line is expected")
		return nil
	}
	return cookies
}

func (p *Parser) parseCookies() *Cookies {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	comments := p.tryParseComments()
	section := p.parseSectionHeader("Cookies")
	spaces := p.tryParseSpaces()
	eol := p.parseEol()
	cookies := p.tryParseNCookie()

	if p.err != nil {
		return nil
	}
	return &Cookies{
		Position{current, line},
		Position{p.current, p.line},
		comments,
		section,
		spaces,
		eol,
		cookies,
	}
}

func (p *Parser) parseQsParams() *QsParams {
	if p.err != nil {
		return nil
	}
	current, line := p.current, p.line

	comments := p.tryParseComments()
	section := p.parseSectionHeader("QueryParams")
	spaces := p.tryParseSpaces()
	eol := p.parseEol()
	params := p.tryParseNKeyValue()

	if p.err != nil {
		return nil
	}
	return &QsParams{
		Position{current, line},
		Position{p.current, p.line},
		comments,
		section,
		spaces,
		eol,
		params,
	}
}

// Specific debug
func (p *Parser) skipToNextEol() {
	_, _ = p.readRunesWhile(isNotNewLine)
	_ = p.parseWhitespaces()
}
