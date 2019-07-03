package ast

import (
	"fmt"
	"io"
	"io/ioutil"
)

type Parser struct {
	filename string // filename, if any
	buffer   []rune // file content
	current  int    // start of the buffer, current rune
	line     int    // current line number in rune, starting at 1
	err      error
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
	return &Parser{filename: filename, buffer: runes, line: 1}
}

func (p *Parser) Parse() *HurlFile {
	return p.parseHurlFile()
}

func (p *Parser) Err() error {
	return p.err
}

func (p *Parser) readRune() (rune, error) {
	r, err := p.nextRune()
	if err != nil {
		return 0, err
	}
	p.current += 1
	if r == '\n' {
		p.line += 1
	}
	return r, nil
}

func (p *Parser) readRunes(count int) ([]rune, error) {
	begin := p.current
	for i := 0; i < count; i++ {
		_, err := p.readRune()
		if err != nil {
			return nil, err
		}
	}
	end := p.current
	return p.buffer[begin:end], nil
}

func (p *Parser) readRunesWhile(f func(rune) bool) ([]rune, error) {
	begin, end := p.current, p.current
	for {
		r, err := p.nextRune()
		if err != nil {
			// We can't read any data any more (EOF), if we haven't been able to read any rune, we
			// return an EOF error, otherwise we return the read slice.
			if begin == end {
				return nil, err
			}
			break
		}
		if f(r) {
			_, _ = p.readRune()
			end = p.current
		} else {
			break
		}
	}
	return p.buffer[begin:end], nil
}

func (p *Parser) nextRune() (rune, error) {
	if p.current >= len(p.buffer) {
		return 0, io.EOF
	}
	return p.buffer[p.current], nil
}

func (p *Parser) nextRunes(count int) ([]rune, error) {
	end := p.current + count
	if end > len(p.buffer) {
		return nil, io.EOF
	}
	return p.buffer[p.current:end], nil
}

func newSyntaxError(p *Parser, text string) error {
	pos := Position{p.current, p.line}
	return &SyntaxError{text, pos}
}

type SyntaxError struct {
	msg string   // description of error
	Pos Position // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Pos.Line, e.msg)
}

func (p *Parser) hasRuneToRead() bool {
	return p.current < len(p.buffer)
}
