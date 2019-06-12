package gurl

import (
	"io"
	"io/ioutil"
)

type Parser struct {
	Filename string // filename, if any
	Content  []rune // file content
	Current  int    // position of the buffer, current rune
}

func NewParserFromFile(path string) (*Parser, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewParserFromString(string(dat), path), nil
}

func NewParserFromString(text string, filename string) *Parser {
	runes := []rune(text)
	return &Parser{Filename: filename, Content: runes}
}

func (p *Parser) readRune() (rune, error) {
	r, err := p.nextRune()
	if err != nil {
		return 0, err
	}
	p.Current += 1
	return r, nil
}

func (p *Parser) nextRune() (rune, error) {
	if p.Current >= len(p.Content) {
		return 0, io.EOF
	}
	return p.Content[p.Current], nil
}

func (p *Parser) parseWhiteSpace(skipNewLine bool) (Token, error) {

	start := p.Current

	for {
		r, err := p.nextRune()
		if err != nil {
			if err == io.EOF {
				return &EofToken{start}, nil
			} else {
				return nil, err
			}
		}
		if isWhiteSpace(r) || (skipNewLine && isNewLine(r)) {
			_, _ = p.readRune()
			continue
		}
		if start == p.Current {
			return nil, nil
		}
		whitespace := string(p.Content[start:p.Current])
		return &WhitespaceToken{start, whitespace}, nil
	}
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isNewLine(r rune) bool {
	return r == '\n'
}

