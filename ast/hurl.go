package ast

import (
	"io"
)

func (p *Parser) parseHurlFile() *HurlFile {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	whitespaces := p.tryParseWhitespaces()
	entries := p.parseNEntry()

	if p.err != nil {
		return nil
	}
	return &HurlFile{Node{pos, p.pos}, whitespaces, entries}
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
			if p.more() {
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
	pos := p.pos

	request := p.parseRequest()
	whitespaces := p.tryParseWhitespaces()
	response := p.tryParseResponse()

	if p.err != nil {
		return nil
	}
	return &Entry{Node{pos, p.pos}, request, whitespaces, response}
}

func (p *Parser) parseRequest() *Request {
	if p.err != nil {
		return nil
	}
	pos := p.pos

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

	var formparams *FormParams
	var body *Body
	formparams = p.tryParseFormParams()
	if formparams == nil {
		body = p.tryParseBody()
	}

	if p.err != nil {
		return nil
	}
	return &Request{
		Node{pos, p.pos},
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
		formparams,
		body,
	}
}

func (p *Parser) parseResponse() *Response {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	version := p.parseVersion()
	spaces0 := p.parseSpaces()
	status := p.parseStatus()
	spaces1 := p.tryParseSpaces()
	comment := p.tryParseComment()
	eol := p.parseEol()
	headers := p.tryParseHeaders()
	captures := p.tryParseCaptures()

	if p.err != nil {
		return nil
	}
	return &Response{
		Node{pos, p.pos},
		comments,
		version,
		spaces0,
		status,
		spaces1,
		comment,
		eol,
		headers,
		captures,
	}
}

func (p *Parser) parseMethod() *Method {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	methods := []string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	}
	for _, m := range methods {
		if p.tryParseString(m) {
			return &Method{Node{pos, p.pos}, m}
		}
	}
	p.err = p.newSyntaxError("method is expected")
	return nil
}

func (p *Parser) parseVersion() *Version {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	methods := []string{"HTTP/1.1", "HTTP/2"}
	for _, v := range methods {
		if p.tryParseString(v) {
			return &Version{Node{pos, p.pos}, v}
		}
	}
	p.err = p.newSyntaxError("method is expected")
	return nil
}

func (p *Parser) parseUrl() *Url {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	isGenDelims := func(r rune) bool {
		return r == ':' || r == '/' || r == '?' || r == '#' || r == '[' || r == ']' || r == '@'
	}
	isSubDelims := func(r rune) bool {
		return r == '!' || r == '$' || r == '&' || r == '\'' || r == '(' || r == ')' ||
			r == '*' || r == '+' || r == ',' || r == ';' || r == '='
	}
	url, err := p.readRunesWhile(func(r rune) bool {
		isUnreserved := isAsciiLetter(r) || isDigit(r) || r == '-' || r == '.' || r == '_' || r == '~'
		isReserved := isGenDelims(r) || isSubDelims(r)
		isHurlSpecific := r == '{' || r == '}'
		return isReserved || isUnreserved || isHurlSpecific
	})
	if err != nil || len(url) == 0 {
		p.err = p.newSyntaxError("url is expected")
		return nil
	}

	return &Url{Node{pos, p.pos}, string(url)}
}

func (p *Parser) parseEol() *Eol {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	eol, err := p.readRunesWhile(isNewLine)
	if err != nil && err != io.EOF {
		p.err = p.newSyntaxError("newline is expected")
		return nil
	}

	if err != io.EOF {
		if len(eol) == 0 {
			p.err = p.newSyntaxError("newline is expected")
			return nil
		}
		_, _ = p.readRunesWhile(isWhitespace)
	}

	return &Eol{Node{pos, p.pos}, string(p.buffer[pos.Offset:p.pos.Offset])}
}

func (p *Parser) parseHeaders() *Headers {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	headers := p.parseNKeyValue()
	if p.err != nil {
		return nil
	}
	return &Headers{Node{pos, p.pos}, headers,}
}

func (p *Parser) parseCookieValue() *CookieValue {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	cookie, err := p.readRunesWhile(func(r rune) bool {
		return isAsciiLetter(r) ||
			isDigit(r) ||
			r == ':' ||
			r == '/' ||
			r == '%'
	})
	if err != nil {
		p.err = p.newSyntaxError("[A-Za-z0-9:/%] char is expected in cookie-value")
		return nil
	}

	return &CookieValue{Node{pos, p.pos}, string(cookie)}
}

func (p *Parser) parseCookie() *Cookie {
	if p.err != nil {
		return nil
	}
	pos := p.pos

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
	return &Cookie{Node{pos, p.pos}, comments, key, spaces0, colon, spaces1, cookieValue, spaces2, comment, eol}
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
		p.err = p.newSyntaxError("At least one cookie is expected")
		return nil
	}
	return cookies
}

func (p *Parser) parseCookies() *Cookies {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	section := p.parseSectionHeader("Cookies")
	spaces := p.tryParseSpaces()
	eol := p.parseEol()
	cookies := p.tryParseNCookie()

	if p.err != nil {
		return nil
	}
	return &Cookies{Node{pos, p.pos}, comments, section, spaces, eol, cookies}
}

func (p *Parser) parseQsParams() *QsParams {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	section := p.parseSectionHeader("QueryParams")
	spaces := p.tryParseSpaces()
	eol := p.parseEol()
	params := p.tryParseNKeyValue()

	if p.err != nil {
		return nil
	}
	return &QsParams{Node{pos, p.pos}, comments, section, spaces, eol, params}
}

func (p *Parser) parseFormParams() *FormParams {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	section := p.parseSectionHeader("FormParams")
	spaces := p.tryParseSpaces()
	eol := p.parseEol()
	params := p.tryParseNKeyValue()

	if p.err != nil {
		return nil
	}
	return &FormParams{Node{pos, p.pos}, comments, section, spaces, eol, params}
}

func (p *Parser) parseBody() *Body {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	obj, text := p.tryParseJson()
	if obj != nil {
		return &Body{Node{pos, p.pos}, text, []byte(text)}
	}
	// TODO: `parseXml` should return an obj plus a text
	text = p.tryParseXml()
	if len(text) > 0 {
		return &Body{Node{pos, p.pos}, text, []byte(text)}
	}
	bs, text := p.tryParseBase64()
	if bs != nil {
		return &Body{Node{pos, p.pos}, text, bs}
	}

	p.err = p.newSyntaxError("body json, xml, base64 or file is expected")
	return nil
}

func (p *Parser) parseStatus() *Status {
	if p.err != nil {
		return nil
	}
	pos := p.pos
	value := p.parseNatural()
	if p.err != nil {
		return nil
	}
	return &Status{Node{pos, p.pos}, value}
}

func (p *Parser) parseQueryExpr() *QueryExpr {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var queryString *QueryString
	var jsonString *JsonString

	queryString = p.tryParseQueryString()
	if queryString == nil {
		jsonString = p.parseJsonString()
		if p.err != nil {
			p.err = p.newSyntaxError("query-string or json-string is expected in key")
			return nil
		}
	}

	var value string
	if queryString != nil {
		value = queryString.Value
	}
	if jsonString != nil {
		value = jsonString.Value
	}
	return &QueryExpr{Node{pos, p.pos}, queryString, jsonString, value}
}

func (p *Parser) parseQuery() *Query {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	spaces0 := p.tryParseSpaces()
	qt := p.parseQueryType()
	var spaces1 *Spaces
	var expr *QueryExpr
	switch qt.Value {
	case "header", "xpath", "jsonpath", "regex":
		spaces1 = p.parseSpaces()
		expr = p.parseQueryExpr()
	}
	spaces2 := p.tryParseSpaces()

	if p.err != nil {
		return nil
	}
	return &Query{Node{pos, p.pos}, spaces0, qt, spaces1, expr, spaces2}
}

func (p *Parser) parseCapture() *Capture {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	key := p.parseKey()
	spaces0 := p.tryParseSpaces()
	colon := p.parseColon()
	spaces1 := p.tryParseSpaces()
	query := p.parseQuery()
	spaces2 := p.tryParseSpaces()
	comment := p.tryParseComment()
	eol := p.parseEol()

	if p.err != nil {
		return nil
	}
	return &Capture{Node{pos, p.pos}, comments, key, spaces0, colon, spaces1, query, spaces2, comment, eol}
}

func (p *Parser) parseCaptures() *Captures {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	section := p.parseSectionHeader("Captures")
	spaces := p.tryParseSpaces()
	eol := p.parseEol()
	captures := p.tryParseNCapture()

	if p.err != nil {
		return nil
	}
	return &Captures{Node{pos, p.pos}, comments, section, spaces, eol, captures}
}

func (p *Parser) parseNCapture() []*Capture {
	if p.err != nil {
		return nil
	}
	captures := make([]*Capture, 0)
	for {
		c := p.tryParseCapture()
		if c == nil {
			break
		}
		captures = append(captures, c)
	}
	if len(captures) == 0 {
		p.err = p.newSyntaxError("At least one capture is expected")
		return nil
	}
	return captures
}