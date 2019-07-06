package ast

type Position struct {
	Offset int // offset in rune, starting at 0
	Line   int // line number in rune, starting at 1
	Column int // column number in rune, starting at 1
}

type Node struct {
	Begin Position
	End   Position
}

type (
	Noder interface {
		GetBegin() Position
		GetEnd() Position
	}

	Cookies struct {
		Node
		Comments      *Comments
		SectionHeader *SectionHeader
		Spaces        *Spaces
		Eol           *Eol
		Cookies       []*Cookie
	}

	Cookie struct {
		Node
		Comments    *Comments
		Key         *Key
		Spaces0     *Spaces
		Colon       *Colon
		Spaces1     *Spaces
		CookieValue *CookieValue
		Spaces2     *Spaces
		Comment     *Comment
		Eol         *Eol
	}

	CookieValue struct {
		Node
		Value string
	}

	Headers struct {
		Node
		Headers []*KeyValue
	}

	Colon struct {
		Node
		Value string
	}

	KeyValue struct {
		Node
		Comments *Comments
		Key      *Key
		Spaces0  *Spaces
		Colon    *Colon
		Spaces1  *Spaces
		Value    *Value
		Spaces2  *Spaces
		Comment  *Comment
		Eol      *Eol
	}

	Key struct {
		Node
		KeyString  *KeyString
		JsonString *JsonString
		Value      string
	}

	Value struct {
		Node
		ValueString *ValueString
		JsonString  *JsonString
		Value       string
	}

	JsonString struct {
		Node
		Text  string
		Value string
	}

	KeyString struct {
		Node
		Value string
	}

	ValueString struct {
		Node
		Value string
	}

	Eol struct {
		Node
		Value string
	}

	Spaces struct {
		Node
		Value string
	}

	Whitespaces struct {
		Node
		Value string
	}

	Request struct {
		Node
		Comments   *Comments
		Method     *Method
		Spaces0    *Spaces
		Url        *Url
		Spaces1    *Spaces
		Comment    *Comment
		Eol        *Eol
		Headers    *Headers
		Cookies    *Cookies
		QsParams   *QsParams
		FormParams *FormParams
		Body       *Body
	}

	Method struct {
		Node
		Value string
	}

	Url struct {
		Node
		Value string
	}

	Comment struct {
		Node
		Value string
	}

	Comments struct {
		Node
		CommentLines []*CommentLine
	}

	Entry struct {
		Node
		Request  *Request
		Whitespaces *Whitespaces
		Response *Response
	}

	HurlFile struct {
		Node
		Whitespaces *Whitespaces
		Entries     []*Entry
	}

	QsParams struct {
		Node
		Comments      *Comments
		SectionHeader *SectionHeader
		Spaces        *Spaces
		Eol           *Eol
		Params        []*KeyValue
	}

	FormParams struct {
		Node
		Comments      *Comments
		SectionHeader *SectionHeader
		Spaces        *Spaces
		Eol           *Eol
		Params        []*KeyValue
	}

	Body struct {
		Node
		Text  string
		Value []byte
	}

	Response struct {
		Node
		Comments *Comments
		Version  *Version
		Spaces0  *Spaces
		Status   *Status
		Spaces1  *Spaces
		Comment  *Comment
		Eol      *Eol
		Headers  *Headers
	}

	Version struct {
		Node
		Value string
	}

	Status struct {
		Node
		Text  string
		Value int
	}
)

// Node not defined in the hurl spec,
type (
	CommentLine struct {
		Node
		Comment     *Comment
		Eol         *Eol
		Whitespaces *Whitespaces
	}

	SectionHeader struct {
		Node
		Value string
	}

	Json = interface{}

	Xml = interface{}
)

func (n *Node) GetBegin() Position {
	return n.Begin
}

func (n *Node) GetEnd() Position {
	return n.End
}
