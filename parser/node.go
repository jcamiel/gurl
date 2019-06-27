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
