package parser

import "fmt"

func (p *Parser) parseRequest() (*Request, error) {

	begin, beginLine := p.Current, p.Line

	whitespaces, _ := p.tryParseWhitespaces()

	method, err := p.parseMethod()
	if err != nil {
		return nil, err
	}

	spaces, err := p.parseSpaces()
	if err != nil {
		return nil, err
	}

	url, err := p.parseUrl()
	if err != nil {
		return nil, err
	}

	end, endLine := p.Current, p.Line

	return &Request{
		Position{begin, beginLine},
		Position{end, endLine},
		whitespaces,
		method,
		spaces,
		url}, nil
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
		return nil, newSyntaxError(p, "url expected")
	}

	end, endLine := p.Current, p.Line

	return &Url{
		Position{begin, beginLine},
		Position{end, endLine},
		string(url)}, nil
}
