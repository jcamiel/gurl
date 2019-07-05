package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (p *Parser) parseWhitespaces() *Whitespaces {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	whitespaces, err := p.readRunesWhile(isWhitespace)
	if err != nil || len(whitespaces) == 0 {
		p.err = p.newSyntaxError("space, tab or newline is expected at whitespaces beginning")
		return nil
	}

	return &Whitespaces{pos, p.pos, string(whitespaces)}
}

func (p *Parser) parseSpaces() *Spaces {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	spaces, err := p.readRunesWhile(isSpace)
	if err != nil || len(spaces) == 0 {
		p.err = p.newSyntaxError("space or tab is expected at spaces beginning")
		return nil
	}

	return &Spaces{pos, p.pos, string(spaces)}
}

func (p *Parser) parseComment() *Comment {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	r, err := p.nextRune()
	if p.err = err; err != nil {
		return nil
	}
	if r != hash {
		p.err = p.newSyntaxError("# is expected at comment beginning")
		return nil
	}
	comment, _ := p.readRunesWhile(isNotNewLine)

	return &Comment{pos, p.pos, string(comment)}
}

func (p *Parser) parseCommentLine() *CommentLine {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comment := p.parseComment()
	eol := p.parseEol()
	whitespaces := p.tryParseWhitespaces()

	if p.err != nil {
		return nil
	}
	return &CommentLine{
		pos,
		p.pos,
		comment,
		eol,
		whitespaces,
	}
}

func (p *Parser) parseComments() *Comments {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	lines := p.parseNCommentLine()

	if p.err != nil {
		return nil
	}
	return &Comments{pos, p.pos, lines}
}

func (p *Parser) parseNCommentLine() []*CommentLine {
	if p.err != nil {
		return nil
	}
	cls := make([]*CommentLine, 0)
	for {
		c := p.tryParseCommentLine()
		if c == nil {
			break
		}
		cls = append(cls, c)
	}
	if len(cls) == 0 {
		p.err = p.newSyntaxError("At least one comment-line is expected")
		return nil
	}
	return cls
}

func (p *Parser) parseJsonString() *JsonString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var value string
	rs := p.buffer[p.pos.Offset:]
	bs := []byte(string(rs))
	r := bytes.NewReader(bs)
	dec := json.NewDecoder(r)
	err := dec.Decode(&value)
	if p.err = err; err != nil {
		return nil
	}

	// We have decoded a valid json string, now we can safely
	// skip the read runes and start and end string delimiter "
	count := len([]rune(value)) + 2
	_, _ = p.readRunes(count)

	return &JsonString{
		pos,
		p.pos,
		string(p.buffer[pos.Offset:p.pos.Offset]),
		value,
	}
}

func (p *Parser) parseKeyString() *KeyString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	key, err := p.readRunesWhile(func(r rune) bool {
		return !isWhitespace(r) && r != ':' && r != '"' && r != '#'
	})
	if err != nil || len(key) == 0 {
		p.err = p.newSyntaxError("char is expected at key-string beginning")
		return nil
	}

	return &KeyString{pos, p.pos, string(key)}
}

func (p *Parser) parseKey() *Key {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var keyString *KeyString
	var jsonString *JsonString

	keyString = p.tryParseKeyString()
	if keyString == nil {
		jsonString = p.parseJsonString()
		if p.err != nil {
			p.err = p.newSyntaxError("key-string or json-string is expected at key beginning")
			return nil
		}
	}

	var value string
	if keyString != nil {
		value = keyString.Value
	}
	if jsonString != nil {
		value = jsonString.Value
	}
	return &Key{
		pos,
		p.pos,
		keyString,
		jsonString,
		value,
	}
}

func (p *Parser) parseValue() *Value {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var valueString *ValueString
	var jsonString *JsonString

	valueString = p.tryParseValueString()
	if valueString == nil {
		jsonString = p.parseJsonString()
		if p.err != nil {
			p.err = p.newSyntaxError("key-string or json-string is expected at key beginning")
			return nil
		}
	}

	var value string
	if valueString != nil {
		value = valueString.Value
	}
	if jsonString != nil {
		value = jsonString.Value
	}
	return &Value{
		pos,
		p.pos,
		valueString,
		jsonString,
		value,
	}
}

func (p *Parser) parseColon() *Colon {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	r, err := p.readRune()
	if err != nil || r != ':' {
		p.err = p.newSyntaxError(": is expected")
		return nil
	}

	return &Colon{pos, p.pos, ":"}
}

func (p *Parser) parseKeyValue() *KeyValue {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	comments := p.tryParseComments()
	key := p.parseKey()
	spaces0 := p.tryParseSpaces()
	colon := p.parseColon()
	spaces1 := p.tryParseSpaces()
	value := p.parseValue()
	space2 := p.tryParseSpaces()
	comment := p.tryParseComment()
	eol := p.parseEol()

	if p.err != nil {
		return nil
	}
	return &KeyValue{
		pos,
		p.pos,
		comments,
		key,
		spaces0,
		colon,
		spaces1,
		value,
		space2,
		comment,
		eol,
	}
}

func (p *Parser) parseNKeyValue() []*KeyValue {
	if p.err != nil {
		return nil
	}
	kvs := make([]*KeyValue, 0)
	for {
		k := p.tryParseKeyValue()
		if k == nil {
			break
		}
		kvs = append(kvs, k)
	}
	if len(kvs) == 0 {
		p.err = p.newSyntaxError("At least one key-value is expected")
		return nil
	}
	return kvs
}

func (p *Parser) parseValueString() *ValueString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	value := make([]rune, 0)
	for {
		v, err := p.readRunesWhile(func(r rune) bool {
			return !isWhitespace(r) && r != '#'
		})
		if err != nil {
			break
		}
		value = append(value, v...)
		n, err := p.nextRune()
		if err != nil || isNewLine(n) || n == '#' {
			break
		}
		if !isSpace(n) {
			r, _ := p.readRune()
			value = append(value, r)
			continue
		}

		// Next spaces can be either part of the trailing spaces (with optional comment)
		// or part if the value-string. In the first case, we must not consume it, and
		// in the second we must continue and keep parsing value-string.
		if p.isTrailingSpaces() {
			break
		}
		s, _ := p.readRunesWhile(isSpace)
		value = append(value, s...)
	}

	if len(value) == 0 {
		p.err = p.newSyntaxError("# or whitespaces is forbidden at value-string beginning")
		return nil
	}
	return &ValueString{pos, p.pos, string(value)}
}

// must start with spaces
func (p *Parser) isTrailingSpaces() bool {
	offset := p.pos.Offset
	_, err := p.readRunesWhile(isSpace)
	if err != nil {
		p.pos.Offset = offset
		return true
	}
	n, err := p.nextRune()
	if err != nil || isNewLine(n) || n == '#' {
		p.pos.Offset = offset
		return true
	}
	p.pos.Offset = offset
	return false
}

func (p *Parser) parseSectionHeader(section string) *SectionHeader {
	if p.err != nil {
		return nil
	}
	pos := p.pos
	s := fmt.Sprintf("[%s]", section)
	if !p.tryParseString(s) {
		p.err = p.newSyntaxError(fmt.Sprintf("[%s] is expected", section))
		return nil
	}
	return &SectionHeader{pos, p.pos, s}
}
