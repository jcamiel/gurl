package gurl

type Position struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
}

type Node interface {
	Position() Position
}

type Eof struct {
	position Position
}

func (t *Eof) Position() Position {
	return t.position
}

type Whitespace struct {
	position Position
	Text     string
}

func (t *Whitespace) Position() Position {
	return t.position
}
