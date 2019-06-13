package gurl

import (
	"io"
	"io/ioutil"
)

const (
	space   = '\u0020'
	tab     = '\u0009'
	newLine = '\u000a'
)

type Parser struct {
	Filename string // filename, if any
	Buffer   []rune // file content
	Current  int    // position of the buffer, current rune
	Line     int    // current line
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
	return &Parser{Filename: filename, Buffer: runes, Line: 1}
}

func (p *Parser) readRune() (rune, error) {
	r, err := p.nextRune()
	if err != nil {
		return 0, err
	}
	p.Current += 1
	if r == newLine {
		p.Line += 1
	}
	return r, nil
}

func (p *Parser) nextRune() (rune, error) {
	if p.Current >= len(p.Buffer) {
		return 0, io.EOF
	}
	return p.Buffer[p.Current], nil
}

func (p *Parser) parseWhiteSpace(skipNewLine bool) (Node, error) {

	startCur := p.Current
	startLine := p.Line

	for {
		r, err := p.nextRune()
		if err != nil || !isWhiteSpace(r) || (isNewLine(r) && !skipNewLine) {
			break
		}
		_, _ = p.readRune()
	}

	if startCur == p.Current {
		return nil, nil
	}
	whitespace := string(p.Buffer[startCur:p.Current])
	position := Position{startCur, startLine}
	return &Whitespace{position, whitespace}, nil
}

func isWhiteSpace(r rune) bool {
	return r == space || r == tab
}

func isNewLine(r rune) bool {
	return r == newLine
}

func (p *Parser) Parse() error {
	return nil
}
