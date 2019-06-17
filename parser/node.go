package parser

type Position struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
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

func (t *Method) Begin() Position {
	return t.begin
}

func (t *Method) End() Position {
	return t.end
}
