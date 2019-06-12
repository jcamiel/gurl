package gurl

type Token interface {
	GetPosition() int
}

type EofToken struct {
	Position int
}

func (t *EofToken) GetPosition() int {
	return t.Position
}

type WhitespaceToken struct {
	Position int
	Text     string
}

func (t *WhitespaceToken) GetPosition() int {
	return t.Position
}
