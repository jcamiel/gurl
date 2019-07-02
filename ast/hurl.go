package ast

import (
	"fmt"
	"io"
)

func (p *Parser) parseHurlFile() (*HurlFile, error) {
	current, line := p.current, p.line

	whitespaces, _ := p.tryParseWhitespaces()
	var entries []*Entry
	for {
		e, err := p.parseEntry()
		if err != nil {
			if p.hasRuneToRead() {
				p.skipToNextEol()
				continue
			}
			break
		}
		entries = append(entries, e)
	}

	return &HurlFile{
		Position{current, line},
		Position{p.current, p.line},
		whitespaces,
		entries,
	}, nil
}

func (p *Parser) parseEntry() (*Entry, error) {
	current, line := p.current, p.line

	request, err := p.parseRequest()
	if err != nil {
		return nil, err
	}

	return &Entry{
		Position{current, line},
		Position{p.current, p.line},
		request,
	}, nil
}

func (p *Parser) parseRequest() (*Request, error) {
	current, line := p.current, p.line

	comments, _ := p.tryParseComments()
	method, err := p.parseMethod()
	if err != nil {
		return nil, err
	}
	spaces0, err := p.parseSpaces()
	if err != nil {
		return nil, err
	}
	url, err := p.parseUrl()
	if err != nil {
		return nil, err
	}
	spaces1, _ := p.tryParseSpaces()
	comment, _ := p.tryParseComment()
	eol, err := p.parseEol()
	if err != nil {
		return nil, err
	}
	headers, _ := p.tryParseHeaders()
	cookies, _ := p.tryParseCookies()

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
	}, nil
}

func (p *Parser) parseMethod() (*Method, error) {
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
			}, nil
		}
	}
	return nil, newSyntaxError(p, fmt.Sprintf("method %v is expected", methods))
}

func (p *Parser) parseUrl() (*Url, error) {
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
		return nil, newSyntaxError(p, "url is expected")
	}

	return &Url{
		Position{current, line},
		Position{p.current, p.line},
		string(url),
	}, nil
}

func (p *Parser) parseEol() (*Eol, error) {
	current, line := p.current, p.line

	eol, err := p.readRunesWhile(func(r rune) bool {
		return isNewLine(r)
	})
	if err != nil && err != io.EOF {
		return nil, newSyntaxError(p, "newline is expected")
	}

	if err != io.EOF {
		if len(eol) == 0 {
			return nil, newSyntaxError(p, "newline is expected")
		}
		_, _ = p.readRunesWhile(func(r rune) bool {
			return isWhitespace(r)
		})
	}

	return &Eol{
		Position{current, line},
		Position{p.current, p.line},
		string(p.buffer[current:p.current]),
	}, nil
}

func (p *Parser) parseHeaders() (*Headers, error) {
	current, line := p.current, p.line

	headers := make([]*KeyValue, 0)
	for {
		h, err := p.tryParseKeyValue()
		if err != nil {
			break
		}
		headers = append(headers, h)
	}
	if len(headers) == 0 {
		return nil, newSyntaxError(p, "headers are expected")
	}

	return &Headers{
		Position{current, line},
		Position{p.current, p.line},
		headers,
	}, nil
}

func (p *Parser) parseCookieValue() (*CookieValue, error) {
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
		return nil, newSyntaxError(p, "[A-Za-z0-9:/%] char is expected for cookie-value")
	}

	return &CookieValue{
		Position{current, line},
		Position{p.current, p.line},
		string(cookie),
	}, nil
}

func (p *Parser) parseCookie() (*Cookie, error) {
	current, line := p.current, p.line

	comments, _ := p.tryParseComments()
	key, err := p.parseKey()
	if err != nil {
		return nil, err
	}
	spaces0, _ := p.tryParseSpaces()
	colon, err := p.parseColon()
	if err != nil {
		return nil, err
	}
	spaces1, _ := p.tryParseSpaces()
	cookieValue, err := p.parseCookieValue()
	if err != nil {
		return nil, err
	}
	spaces2, _ := p.tryParseSpaces()
	comment, _ := p.tryParseComment()
	eol, err := p.parseEol()
	if err != nil {
		return nil, err
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
	}, nil
}

func (p *Parser) parseCookies() (*Cookies, error) {
	current, line := p.current, p.line

	comments, _ := p.tryParseComments()
	section, err := p.parseSectionHeader("Cookies")
	if err != nil {
		return nil, err
	}
	spaces, _ := p.tryParseSpaces()
	eol, err := p.parseEol()
	if err != nil {
		return nil, err
	}
	cookies := make([]*Cookie, 0)
	for {
		c, err := p.tryParseCookie()
		if err != nil {
			break
		}
		cookies = append(cookies, c)
	}

	return &Cookies{
		Position{current, line},
		Position{p.current, p.line},
		comments,
		section,
		spaces,
		eol,
		cookies,
	}, nil
}

// Specific debug
func (p *Parser) skipToNextEol() {
	_, _ = p.readRunesWhile(func(r rune) bool {
		return !isWhitespace(r)
	})

	_, _ = p.parseWhitespaces()
}
