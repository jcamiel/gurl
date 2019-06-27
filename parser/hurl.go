package parser

import (
	"fmt"
	"io"
)

func (p *Parser) parseHurlFile() (*HurlFile, error) {

	begin, beginLine := p.Current, p.Line

	whitespaces, _ := p.tryParseWhitespaces()

	var entries []*Entry

	for {
		e, err := p.parseEntry()
		if err != nil {
			if p.hasRuneToRead() {
				p.seekToNextEol()
				continue
			}
			break
		}
		entries = append(entries, e)
	}

	end, endLine := p.Current, p.Line

	return &HurlFile{
		Position{begin, beginLine},
		Position{end, endLine},
		whitespaces,
		entries,
	}, nil

}

func (p *Parser) parseEntry() (*Entry, error) {

	begin, beginLine := p.Current, p.Line

	request, err := p.parseRequest()
	if err != nil {
		return nil, err
	}

	end, endLine := p.Current, p.Line

	return &Entry{
		Position{begin, beginLine},
		Position{end, endLine},
		request,
	}, nil
}

func (p *Parser) parseRequest() (*Request, error) {

	begin, beginLine := p.Current, p.Line

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

	end, endLine := p.Current, p.Line

	return &Request{
		Position{begin, beginLine},
		Position{end, endLine},
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
			begin := Position{p.Current, p.Line}
			end := Position{p.Current + count, p.Line}
			_, _ = p.readRunes(count)
			return &Method{begin, end, method}, nil
		}
	}
	return nil, newSyntaxError(p, fmt.Sprintf("method %v is expected", methods))
}

func (p *Parser) parseUrl() (*Url, error) {

	begin, beginLine := p.Current, p.Line

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

	end, endLine := p.Current, p.Line

	return &Url{
		Position{begin, beginLine},
		Position{end, endLine},
		string(url),
	}, nil
}

func (p *Parser) parseEol() (*Eol, error) {

	begin, beginLine := p.Current, p.Line

	eol, err := p.readRunesWhile(func(r rune) bool {
		return IsNewline(r)
	})

	if err != nil && err != io.EOF {
		return nil, newSyntaxError(p, "newline is expected")
	}

	end, endLine := p.Current, p.Line

	return &Eol{
		Position{begin, beginLine},
		Position{end, endLine},
		string(eol),
	}, nil

}


// Specific debug
func (p *Parser) seekToNextEol()  {
	_, _ = p.readRunesWhile(func(r rune) bool {
		return !IsNewline(r)
	})

	_, _ = p.parseWhitespaces()
}