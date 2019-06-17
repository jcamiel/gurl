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
	Column   int    // current column number in rune, starting at 1
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
	return &Parser{Filename: filename, Buffer: runes, Line: 1, Column: 1}
}

func (p *Parser) readRune() (rune, error) {
	r, err := p.nextRune()
	if err != nil {
		return 0, err
	}
	p.Current += 1
	p.Column += 1
	if r == newLine {
		p.Line += 1
		p.Column = 1
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

	begin, beginLine, beginColumn := p.Current, p.Line, p.Column
	end, endLine, endColumn := begin, beginLine, beginColumn

	for {
		r, err := p.nextRune()
		if err != nil {
			break
		}
		if isWhiteSpace(r) || (isNewLine(r) && skipNewLine) {
			_, _ = p.readRune()
			end, endLine, endColumn = p.Current, p.Line, p.Column
		} else {
			break
		}
	}

	if begin == end {
		return nil, nil
	}
	beginPos := Position{begin, beginLine, beginColumn}
	endPos := Position{end, endLine, endColumn}
	whitespace := string(p.Buffer[begin:end])
	return &Whitespace{beginPos, endPos, whitespace}, nil
}

func (p *Parser) parseUnquotedString() (Node, error) {

	begin, beginLine, beginColumn := p.Current, p.Line, p.Column
	end, endLine, endColumn := begin, beginLine, beginColumn

	r, err := p.nextRune()
	if err != nil {
		return nil, err
	}
	if r == quote || isWhiteSpace(r) {
		return nil, newSyntaxError(p, "unquoted string should begin with space or \"")
	}

	for {
		r, err := p.nextRune()
		if err != nil {
			break
		}

		// TODO: manage unicode literal
		if !isWhiteSpace(r) && !isNewLine(r) && !isHash(r) {
			_, _ = p.readRune()
			end, endLine, endColumn = p.Current, p.Line, p.Column
		} else {
			break
		}
	}

	beginPos := Position{begin, beginLine, beginColumn}
	endPos := Position{end, endLine, endColumn}
	text := string(p.Buffer[begin:end])
	return &UnquotedString{beginPos, endPos, text}, nil
}

func newSyntaxError(p *Parser, text string) error {
	pos := Position{p.Current, p.Line, p.Current}
	return &SyntaxError{text, pos}
}

type SyntaxError struct {
	msg string   // description of error
	Pos Position // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("[%d:%d] %s", e.Pos.Line, e.Pos.Column, e.msg)
}
