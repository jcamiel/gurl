package parser

import (
	"errors"
	"fmt"
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
	Current  int    // start of the buffer, current rune
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

func (p *Parser) readRunes(count int) ([]rune, error) {
	begin := p.Current
	for i := 0; i < count; i++ {
		_, err := p.readRune()
		if err != nil {
			return nil, err
		}
	}
	end := p.Current
	return p.Buffer[begin:end], nil
}

func (p *Parser) nextRune() (rune, error) {
	if p.Current >= len(p.Buffer) {
		return 0, io.EOF
	}
	return p.Buffer[p.Current], nil
}

func (p *Parser) nextRunes(count int) ([]rune, error) {
	end := p.Current + count
	if end > len(p.Buffer) {
		return nil, io.EOF
	}
	return p.Buffer[p.Current:end], nil
}

func (p *Parser) isNext(text string) bool {
	runes := []rune(text)
	next, err := p.nextRunes(len(runes))
	if err != nil {
		return false
	}
	return Equal(runes, next)
}

func (p *Parser) parseWhiteSpace(skipNewLine bool) (Node, error) {

	begin := p.Current
	beginLine := p.Line
	end := begin
	endLine := beginLine

	for {
		r, err := p.nextRune()
		if err != nil {
			break
		}
		if isWhiteSpace(r) || (isNewLine(r) && skipNewLine) {
			_, _ = p.readRune()
			end = p.Current
			endLine = p.Line
		} else {
			break
		}
	}

	if begin == end {
		return nil, nil
	}
	beginPos := Position{begin, beginLine}
	endPos := Position{end, endLine}
	whitespace := string(p.Buffer[begin:end])
	return &Whitespace{beginPos, endPos, whitespace}, nil
}

func isWhiteSpace(r rune) bool {
	return r == space || r == tab
}

func isNewLine(r rune) bool {
	return r == newLine
}

func newSyntaxError(p *Parser, text string) error {
	textError := fmt.Sprintf("Line %d column %d, %s", p.Line, 0, text)
	return errors.New(textError)
}

