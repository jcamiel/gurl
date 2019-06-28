package ast

type Position struct {
	Offset int // offset in rune, starting at 0
	Line   int // line number in rune, starting at 1
}

type (
	Node interface {
		Begin() Position
		End() Position
	}

	Eol struct {
		begin Position
		end   Position
		Text  string
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
		begin    Position
		end      Position
		Comments *Comments
		Method   *Method
		Spaces0  *Spaces
		Url      *Url
		Spaces1  *Spaces
		Comment  *Comment
		Eol      *Eol
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

	Comments struct {
		begin        Position
		end          Position
		CommentLines []*CommentLine
	}

	Entry struct {
		begin   Position
		end     Position
		Request *Request
	}

	HurlFile struct {
		begin       Position
		end         Position
		Whitespaces *Whitespaces
		Entries     []*Entry
	}
)


// Node not defined in the hurl spec,
type (
	CommentLine struct {
		begin       Position
		end         Position
		Comment     *Comment
		Eol         *Eol
		Whitespaces *Whitespaces
	}
)

func (h *HurlFile) Begin() Position {
	return h.begin
}

func (h *HurlFile) End() Position {
	return h.end
}

func (w *Whitespaces) Begin() Position {
	return w.begin
}

func (w *Whitespaces) End() Position {
	return w.end
}

func (e *Entry) Begin() Position {
	return e.begin
}

func (e *Entry) End() Position {
	return e.end
}

func (r *Request) Begin() Position {
	return r.begin
}

func (r *Request) End() Position {
	return r.end
}

func (c *Comments) Begin() Position {
	return c.begin
}

func (c *Comments) End() Position {
	return c.end
}

func (m *Method) Begin() Position {
	return m.begin
}

func (m *Method) End() Position {
	return m.end
}

func (s *Spaces) Begin() Position {
	return s.begin
}

func (s *Spaces) End() Position {
	return s.end
}

func (c *CommentLine) Begin() Position {
	return c.begin
}

func (c *CommentLine) End() Position {
	return c.end
}

func (c *Comment) Begin() Position {
	return c.begin
}

func (c *Comment) End() Position {
	return c.end
}

func (e *Eol) Begin() Position {
	return e.begin
}

func (e *Eol) End() Position {
	return e.end
}