package parser

type Position struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
}

type Node interface {
	Start() Position
	End() Position
}

type Eof struct {
	start Position
	end   Position
}

func (t *Eof) Start() Position {
	return t.start
}

func (t *Eof) End() Position {
	return t.end
}

type Whitespace struct {
	start Position
	end Position
	Text  string
}

func (t *Whitespace) Start() Position {
	return t.start
}

func (t *Whitespace) End() Position {
	return t.end
}
