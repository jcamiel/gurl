package ast

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
)

func (p *Parser) parseWhitespaces() *Whitespaces {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	w, err := p.readRunesWhile(isWhitespace)
	if err != nil || len(w) == 0 {
		p.err = p.newSyntaxError("space, tab or newline is expected in whitespaces")
		return nil
	}

	return &Whitespaces{Node{pos, p.pos}, string(w)}
}

func (p *Parser) parseSpaces() *Spaces {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	spaces, err := p.readRunesWhile(isSpace)
	if err != nil || len(spaces) == 0 {
		p.err = p.newSyntaxError("space or tab is expected in spaces")
		return nil
	}

	return &Spaces{Node{pos, p.pos}, string(spaces)}
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

	return &Comment{Node{pos, p.pos}, string(comment)}
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
	return &CommentLine{Node{pos, p.pos}, comment, eol, whitespaces}
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
	return &Comments{Node{pos, p.pos}, lines}
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

	r, err := p.readRune()
	if p.err = err; err != nil {
		return nil
	}
	if r != '"' {
		p.err = p.newSyntaxError("\" is expected at json-string beginning")
		return nil
	}

	var pr rune
	for {
		r, err := p.readRune()
		if p.err = err; err != nil {
			return nil
		}
		if r == '"' && pr != '\\' {
			break
		}
		if isNewLine(r) {
			p.err = p.newSyntaxError("newline not allowed in json-string")
			return nil
		}
		pr = r
	}

	var value string
	rs := p.buffer[pos.Offset:p.pos.Offset]
	text := string(rs)
	bs := []byte(text)
	err = json.Unmarshal(bs, &value)
	if p.err = err; err != nil {
		return nil
	}
	return &JsonString{Node{pos, p.pos}, text, value}
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
		p.err = p.newSyntaxError("char is expected in key-string")
		return nil
	}

	return &KeyString{Node{pos, p.pos}, string(key)}
}

func (p *Parser) parseKey() *Key {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var k *KeyString
	var j *JsonString
	var value string
	if k = p.tryParseKeyString(); k != nil {
		value = k.Value
	} else if j = p.parseJsonString(); j != nil {
		value = j.Value
	} else {
		p.err = p.newSyntaxError("key-string or json-string is expected in key")
		return nil
	}
	return &Key{Node{pos, p.pos}, k, j, value}
}

func (p *Parser) parseValue() *Value {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var v *ValueString
	var j *JsonString
	var value string
	if v = p.tryParseValueString(); v != nil {
		value = v.Value
	} else if j = p.parseJsonString(); j != nil {
		value = j.Value
	} else {
		p.err = p.newSyntaxError("key-string or json-string is expected in key")
	}
	return &Value{
		Node{pos, p.pos},
		v,
		j,
		value,
	}
}

func (p *Parser) parseColon() *Colon {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	if r, err := p.readRune(); err != nil || r != ':' {
		p.err = p.newSyntaxError(": is expected")
		return nil
	}

	return &Colon{Node{pos, p.pos}, ":"}
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
		Node{pos, p.pos},
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
		p.err = p.newSyntaxError("key-value is expected")
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
	return &ValueString{Node{pos, p.pos}, string(value)}
}

func (p *Parser) parseJson() (value Json, text string) {
	if p.err != nil {
		return nil, ""
	}

	rs := p.buffer[p.pos.Offset:]
	bs := []byte(string(rs))
	reader := bytes.NewReader(bs)
	dec := json.NewDecoder(reader)
	err := dec.Decode(&value)
	if p.err = err; err != nil {
		return nil, ""
	}

	// Count the number of bytes read by the json parsing
	// and convert the number of bytes read to the number of rune
	// read so we can advance the parser.
	remaining, err := ioutil.ReadAll(dec.Buffered())
	if p.err = err; err != nil {
		return nil, ""
	}
	readBytes := len(bs) - reader.Len() - len(remaining)
	text = string(bs[:readBytes])
	runes := []rune(text)
	_, _ = p.readRunes(len(runes))
	return
}

func (p *Parser) parseXml() string {
	if p.err != nil {
		return ""
	}

	// Basic check to discard non xml text : the xml stream decoder read the stream
	// until a valid xml is found ('abcd<xml' will be decoded), and we want to avoid
	// this.
	r, err := p.nextRune()
	if err != nil || r != '<' {
		p.err = p.newSyntaxError("< is expected at xml beginning")
		return ""
	}

	var value interface{}
	rs := p.buffer[p.pos.Offset:]
	bs := []byte(string(rs))
	reader := bytes.NewReader(bs)
	dec := xml.NewDecoder(reader)
	err = dec.Decode(&value)
	if p.err = err; err != nil {
		return ""
	}

	// Count the number of bytes read by the xml parsing
	// and convert the number of bytes read to the number of rune
	// read so we can advance the parser.
	readBytes := dec.InputOffset()
	text := string(bs[:readBytes])
	runes := []rune(text)
	_, _ = p.readRunes(len(runes))
	return text
}

func (p *Parser) parseBase64() (value []byte, text string) {
	if p.err != nil {
		return nil, ""
	}
	pos := p.pos

	if !p.tryParseString("base64,") {
		p.err = p.newSyntaxError("base64, is expected at base64 body start")
		return nil, ""
	}
	runes, err := p.readRunesWhile(func(r rune) bool {
		return r != ';'
	})
	bs, err := base64.StdEncoding.DecodeString(string(runes))
	if p.err = err; err != nil {
		return nil, ""
	}
	r, err := p.readRune()
	if err != nil || r != ';' {
		p.err = p.newSyntaxError("; is expected at base64 body end")
		return nil, ""
	}
	return bs, string(p.buffer[pos.Offset:p.pos.Offset])
}

func (p *Parser) parseNatural() *Natural {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	digits, err := p.readRunesWhile(isDigit)
	if err != nil || len(digits) == 0 {
		p.err = p.newSyntaxError("0-9 is expected")
		return nil
	}

	text := string(digits)
	value, err := strconv.Atoi(text)
	if p.err = err; p.err != nil {
		return nil
	}
	return &Natural{Node{pos, p.pos}, text, value}
}

func (p *Parser) parseInteger() *Integer {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	r, err := p.nextRune()
	if p.err = err; err != nil {
		return nil
	}
	if r == '+' || r == '-' {
		_, _ = p.readRune()
	} else if !isDigit(r) {
		p.err = p.newSyntaxError("+-0-9 is expected")
		return nil
	}
	digits, err := p.readRunesWhile(isDigit)
	if err != nil || len(digits) == 0 {
		p.err = p.newSyntaxError("+-0-9 is expected")
		return nil
	}
	text := string(p.buffer[pos.Offset:p.pos.Offset])
	value, err := strconv.Atoi(text)
	if err != nil {
		p.err = p.newSyntaxError("+-0-9 is expected")
		return nil
	}
	return &Integer{Node{pos, p.pos}, text, value}
}

func (p *Parser) parseFloat() *Float {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	r, err := p.nextRune()
	if p.err = err; err != nil {
		return nil
	}
	if r == '+' || r == '-' {
		_, _ = p.readRune()
	} else if !isDigit(r) {
		p.err = p.newSyntaxError("+-[0-9].[0-9] is expected")
		return nil
	}
	digits, err := p.readRunesWhile(isDigit)
	if err != nil || len(digits) == 0 {
		p.err = p.newSyntaxError("+-[0-9].[0-9] is expected")
		return nil
	}
	r, err = p.readRune()
	if err != nil || r != '.' {
		p.err = p.newSyntaxError(". is expected")
	}
	digits, err = p.readRunesWhile(isDigit)
	if err != nil || len(digits) == 0 {
		p.err = p.newSyntaxError("+-[0-9].[0-9] is expected")
		return nil
	}
	text := string(p.buffer[pos.Offset:p.pos.Offset])
	value, err := strconv.ParseFloat(text, 64)
	if p.err = err; err != nil {
		return nil
	}
	return &Float{Node{pos, p.pos}, text, value}
}

func (p *Parser) parseBool() *Bool {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	var text string
	var value bool
	if p.tryParseString("true") {
		text = "true"
		value = true
	} else if p.tryParseString("false") {
		text = "false"
		value = false
	} else {
		p.err = p.newSyntaxError("true or false is expected")
		return nil
	}
	return &Bool{Node{pos, p.pos}, text, value}
}

func (p *Parser) parseQueryString() *QueryString {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	r, err := p.nextRune()
	if err != nil {
		p.err = p.newSyntaxError("char is expected in query")
		return nil
	}
	if r == '"' {
		p.err = p.newSyntaxError("\" is not allowed at query-string beginning")
		return nil
	}

	s, err := p.readRunesWhile(func(r rune) bool {
		return r >= '!' && r <= '~'
	})
	if err != nil {
		p.err = p.newSyntaxError("ascii char is expected in query")
		return nil
	}

	return &QueryString{Node{pos, p.pos}, string(s)}
}

func (p *Parser) parseQueryType() *QueryType {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	types := []string{"status", "header", "body", "xpath", "jsonpath", "regex"}
	for _, t := range types {
		if p.tryParseString(t) {
			return &QueryType{Node{pos, p.pos}, t}
		}
	}
	p.err = p.newSyntaxError("query type is expected")
	return nil
}

func (p *Parser) parsePredicateType() *PredicateType {
	if p.err != nil {
		return nil
	}
	pos := p.pos

	types := []string{"equals", "matches", "startsWith", "contains"}
	for _, t := range types {
		if p.tryParseString(t) {
			return &PredicateType{Node{pos, p.pos}, t}
		}
	}
	p.err = p.newSyntaxError("predicate type is expected")
	return nil
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
	return &SectionHeader{Node{pos, p.pos}, s}
}
