package parser

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Parser struct {
	Filename string // filename, if any
	Buffer   []rune // file content
	Current  int    // start of the buffer, current rune
	Line     int    // current line number in rune, starting at 1
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

func (p *Parser) readRunesWhile(f func(rune) bool) ([]rune, error) {
	begin := p.Current
	end := begin

	for {
		r, err := p.nextRune()
		if err != nil {
			break
		}
		if f(r) {
			_, _ = p.readRune()
			end = p.Current
		} else {
			break
		}
	}
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

func (p *Parser) tryParse(parse func() (Node, error)) (Node, error) {

	begin, beginLine := p.Current, p.Line

	node, err := parse()

	if err != nil {
		p.Current, p.Line = begin, beginLine
		return nil, err
	}

	return node, nil
}

func newSyntaxError(p *Parser, text string) error {
	pos := Position{p.Current, p.Line}
	return &SyntaxError{text, pos}
}

type SyntaxError struct {
	msg string   // description of error
	Pos Position // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Pos.Line, e.msg)
}
