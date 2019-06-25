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

	Eol struct {
		begin Position
		end   Position
		Text string
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

	Request struct {
		begin   Position
		end     Position
		Method  *Method
		Spaces0 *Spaces
		Url     *Url
		Spaces1 *Spaces
		Comment *Comment
		Eol		*Eol
	}

	Method struct {
		begin Position
		end   Position
		Value string
	}

	Url struct {
		begin Position
		end   Position
		Text  string
	}

	Comment struct {
		begin Position
		end   Position
		Text  string
	}
)

func (t *Eof) Begin() Position {
	return t.begin
}

func (t *Eof) End() Position {
	return t.end
}

func (t *Eol) Begin() Position {
	return t.begin
}

func (t *Eol) End() Position {
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

func (t *Request) Begin() Position {
	return t.begin
}

func (t *Request) End() Position {
	return t.end
}

func (t *Url) Begin() Position {
	return t.begin
}

func (t *Url) End() Position {
	return t.end
}

func (t *Comment) Begin() Position {
	return t.begin
}

func (t *Comment) End() Position {
	return t.end
}