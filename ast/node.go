package ast

type Position struct {
	Offset int // offset in rune, starting at 0
	Line   int // line number in rune, starting at 1
}

type (
	Noder interface {
		Node() (Position, Position)
	}

	JsonString struct {
		Begin Position
		End   Position
		Text  string
		Value string
	}

	Eol struct {
		Begin Position
		End   Position
		Text  string
	}

	Spaces struct {
		Begin Position
		End   Position
		Text  string
	}

	Whitespaces struct {
		Begin Position
		End   Position
		Text  string
	}

	Request struct {
		Begin    Position
		End      Position
		Comments *Comments
		Method   *Method
		Spaces0  *Spaces
		Url      *Url
		Spaces1  *Spaces
		Comment  *Comment
		Eol      *Eol
	}

	Method struct {
		Begin Position
		End   Position
		Value string
	}

	Url struct {
		Begin Position
		End   Position
		Text  string
	}

	Comment struct {
		Begin Position
		End   Position
		Text  string
	}

	Comments struct {
		Begin        Position
		End          Position
		CommentLines []*CommentLine
	}

	Entry struct {
		Begin   Position
		End     Position
		Request *Request
	}

	HurlFile struct {
		Begin       Position
		End         Position
		Whitespaces *Whitespaces
		Entries     []*Entry
	}
)

// Node not defined in the hurl spec,
type (
	CommentLine struct {
		Begin       Position
		End         Position
		Comment     *Comment
		Eol         *Eol
		Whitespaces *Whitespaces
	}
)

func (h *HurlFile) Node() (Position, Position) {
	return h.Begin, h.End
}

func (w *Whitespaces) Node() (Position, Position) {
	return w.Begin, w.End
}

func (e *Entry) Node() (Position, Position) {
	return e.Begin, e.End
}

func (r *Request) Node() (Position, Position) {
	return r.Begin, r.End
}

func (c *Comments) Node() (Position, Position) {
	return c.Begin, c.End
}

func (m *Method) Node() (Position, Position) {
	return m.Begin, m.End
}

func (s *Spaces) Node() (Position, Position) {
	return s.Begin, s.End
}

func (c *CommentLine) Node() (Position, Position) {
	return c.Begin, c.End
}

func (c *Comment) Node() (Position, Position) {
	return c.Begin, c.End
}

func (e *Eol) Node() (Position, Position) {
	return e.Begin, e.End
}

func (u *Url) Node() (Position, Position) {
	return u.Begin, u.End
}
