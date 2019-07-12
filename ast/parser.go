package ast

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Parser struct {
	filename string // filename, if any
	buffer   []rune // file content
	pos      Position
	err      error
	errs     []error
}

type SyntaxError struct {
	Msg string   // description of error
	Pos Position // error occurred after reading Offset bytes
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
	return &Parser{filename: filename, buffer: runes, pos: Position{0, 1, 1}}
}

func (p *Parser) Parse() *HurlFile {
	return p.parseHurlFile()
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) Errs() []error {
	return p.errs
}

func (p *Parser) readRune() (rune, error) {
	r, err := p.nextRune()
	if err != nil {
		return 0, err
	}
	p.pos.Offset += 1
	if !isCombining(r) {
		p.pos.Column += 1
	}
	if r == '\n' {
		p.pos.Line += 1
		p.pos.Column = 1
	}
	return r, nil
}

func (p *Parser) readRunes(count int) ([]rune, error) {
	offset := p.pos.Offset
	for i := 0; i < count; i++ {
		_, err := p.readRune()
		if err != nil {
			return nil, err
		}
	}
	return p.buffer[offset:p.pos.Offset], nil
}

func (p *Parser) readRunesWhile(f func(rune) bool) ([]rune, error) {
	offset := p.pos.Offset
	for {
		r, err := p.nextRune()
		if err != nil {
			// We can't read any data any more (EOF), if we haven't been able to read any rune, we
			// return an EOF error, otherwise we return the read slice.
			if offset == p.pos.Offset {
				return nil, err
			}
			break
		}
		if f(r) {
			_, _ = p.readRune()
		} else {
			break
		}
	}
	return p.buffer[offset:p.pos.Offset], nil
}

func (p *Parser) nextRune() (rune, error) {
	if p.pos.Offset >= len(p.buffer) {
		return 0, io.EOF
	}
	return p.buffer[p.pos.Offset], nil
}

func (p *Parser) nextRunes(count int) ([]rune, error) {
	end := p.pos.Offset + count
	if end > len(p.buffer) {
		return nil, io.EOF
	}
	return p.buffer[p.pos.Offset:end], nil
}

func (p *Parser) nextRunesMax(count int) ([]rune, error) {
	end := p.pos.Offset + count
	if end > len(p.buffer) {
		return p.buffer[p.pos.Offset:], nil
	}
	return p.buffer[p.pos.Offset:end], nil
}

// Number of runes left to parse
func (p *Parser) left() int {
	return len(p.buffer) - p.pos.Offset
}

func (p *Parser) newSyntaxError(text string) error {
	return &SyntaxError{text, p.pos}
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("[%d:%d] %s", e.Pos.Line, e.Pos.Column, e.Msg)
}

func (p *Parser) rewindTo(pos Position) {
	if p.err != nil {
		p.errs = append(p.errs, p.err)
		p.err = nil
	}
	p.pos = pos
}
