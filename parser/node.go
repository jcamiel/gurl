package parser

type Position struct {
	Offset int // offset in rune, starting at 0
	Line   int // line number in rune, starting at 1
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

	Spaces struct {
		begin Position
		end   Position
		Text  string
	}

	Whitespaces struct {
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

func (t *Spaces) Begin() Position {
	return t.begin
}

func (t *Spaces) End() Position {
	return t.end
}

func (t *Whitespaces) Begin() Position {
	return t.begin
}

func (t *Whitespaces) End() Position {
	return t.end
}

func (t *Method) Begin() Position {
	return t.begin
}

func (t *Method) End() Position {
	return t.end
}
