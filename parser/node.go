package parser

type Position struct {
	Offset int // offset in rune, starting at 0
	Line   int // line number in rune, starting at 1
	Column int // column number in rune, starting at 1
}

type (
	Node interface {
		Begin() Position
		End() Position
	}

	Eof struct {
		begin Position
		end   Position
	}

	Whitespace struct {
		begin Position
		end   Position
		Text  string
	}

	UnquotedString struct {
		begin Position
		end   Position
		Text  string
	}

	Method struct {
		begin  Position
		end    Position
		Method string
	}
)

func (t *Eof) Begin() Position {
	return t.begin
}

func (t *Eof) End() Position {
	return t.end
}

func (t *Whitespace) Begin() Position {
	return t.begin
}

func (t *Whitespace) End() Position {
	return t.end
}

func (t *UnquotedString) Begin() Position {
	return t.begin
}

func (t *UnquotedString) End() Position {
	return t.end
}

func (t *Method) Begin() Position {
	return t.begin
}

func (t *Method) End() Position {
	return t.end
}
