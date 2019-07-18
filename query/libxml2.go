package query

// XMLNodeType identifies the type of the underlying C struct
type XMLNodeType int

const (
	ElementNode XMLNodeType = iota + 1
	AttributeNode
	TextNode
	CDataSectionNode
	EntityRefNode
	EntityNode
	PiNode
	CommentNode
	DocumentNode
	DocumentTypeNode
	DocumentFragNode
	NotationNode
	HTMLDocumentNode
	DTDNode
	ElementDecl
	AttributeDecl
	EntityDecl
	NamespaceDecl
	XIncludeStart
	XIncludeEnd
	DocbDocumentNode
)


type XPathObjectType int

const (
	XPathUndefinedType XPathObjectType = iota
	XPathNodeSetType
	XPathBooleanType
	XPathNumberType
	XPathStringType
	XPathPointType
	XPathRangeType
	XPathLocationSetType
	XPathUsersType
	XPathXSLTTreeType
)
