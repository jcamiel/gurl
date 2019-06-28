package ast

type Position struct {
	Offset int // offset in rune, starting at 0
	Line   int // line number in rune, starting at 1
}

type (
	Node interface {
		GetBegin() Position
		GetEnd() Position
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

func (h *HurlFile) GetBegin() Position {
	return h.Begin
}

func (h *HurlFile) GetEnd() Position {
	return h.End
}

func (w *Whitespaces) GetBegin() Position {
	return w.Begin
}

func (w *Whitespaces) GetEnd() Position {
	return w.End
}

func (e *Entry) GetBegin() Position {
	return e.Begin
}

func (e *Entry) GetEnd() Position {
	return e.End
}

func (r *Request) GetBegin() Position {
	return r.Begin
}

func (r *Request) GetEnd() Position {
	return r.End
}

func (c *Comments) GetBegin() Position {
	return c.Begin
}

func (c *Comments) GetEnd() Position {
	return c.End
}

func (m *Method) GetBegin() Position {
	return m.Begin
}

func (m *Method) GetEnd() Position {
	return m.End
}

func (s *Spaces) GetBegin() Position {
	return s.Begin
}

func (s *Spaces) GetEnd() Position {
	return s.End
}

func (c *CommentLine) GetBegin() Position {
	return c.Begin
}

func (c *CommentLine) GetEnd() Position {
	return c.End
}

func (c *Comment) GetBegin() Position {
	return c.Begin
}

func (c *Comment) GetEnd() Position {
	return c.End
}

func (e *Eol) GetBegin() Position {
	return e.Begin
}

func (e *Eol) GetEnd() Position {
	return e.End
}

func (u *Url) GetBegin() Position {
	return u.Begin
}

func (u *Url) GetEnd() Position {
	return u.End
}