package parser

import "fmt"

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

	end, endLine := p.Current, p.Line

	return &Request{
		Position{begin, beginLine},
		Position{end, endLine},
		whitespaces,
		method,
		spaces}, nil
}
