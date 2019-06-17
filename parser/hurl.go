package parser

import "fmt"

func (p *Parser) parseMethod() (Node, error) {
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
			begin := Position{p.Current, p.Line, p.Column}
			end := Position{p.Current + count, p.Line, p.Column + count}
			_, _ = p.readRunes(count)
			return &Method{begin, end, method}, nil
		}
	}
	return nil, newSyntaxError(p, fmt.Sprintf("method %v is expected", methods))
}

func (p *Parser) parseUrl() (Node, error) {
	return nil, nil
}