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
	}, nil
}

func (p *Parser) parseMethod() (*Method, error) {
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
		if p.isNext(method) {
			count := len(method)
			begin := Position{p.current, p.line}
			end := Position{p.current + count, p.line}
			_, _ = p.readRunes(count)
			return &Method{begin, end, method}, nil
		}
	}
	return nil, newSyntaxError(p, fmt.Sprintf("method %v is expected", methods))
}

func (p *Parser) parseUrl() (*Url, error) {

	current, line := p.current, p.line

	genDelims := []rune{':', '/', '?', '#', '[', ']', '@'}
	subDelims := []rune{'!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '='}

	url, err := p.readRunesWhile(func(r rune) bool {
		isAlpha := (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
		isDigit := r >= '0' && r <= '9'
		isUnreserved := isAlpha || isDigit || r == '-' || r == '.' || r == '_' || r == '~'
		isReserved := RuneInSlice(r, genDelims) || RuneInSlice(r, subDelims)
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
		return IsNewline(r)
	})

	if err != nil && err != io.EOF {
		return nil, newSyntaxError(p, "newline is expected")
	}

	return &Eol{
		Position{current, line},
		Position{p.current, p.line},
		string(eol),
	}, nil

}


// Specific debug
func (p *Parser) skipToNextEol()  {
	_, _ = p.readRunesWhile(func(r rune) bool {
		return !IsNewline(r)
	})

	_, _ = p.parseWhitespaces()
}