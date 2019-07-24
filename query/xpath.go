package query

/*
#cgo CFLAGS: -I${SRCDIR}/../clib/darwin/amd64/include/libxml2
#cgo LDFLAGS: ${SRCDIR}/../clib/darwin/amd64/lib/libxml2.a
#cgo LDFLAGS: ${SRCDIR}/../clib/darwin/amd64/lib/libz.a
#cgo LDFLAGS: ${SRCDIR}/../clib/darwin/amd64/lib/libiconv.a
#cgo LDFLAGS: ${SRCDIR}/../clib/darwin/amd64/lib/liblzma.a
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <assert.h>

#include <libxml/tree.h>
#include <libxml/parser.h>
#include <libxml/xpath.h>
#include <libxml/xpathInternals.h>
#include <libxml/HTMLparser.h>

void _nilGenericErrorHandler(void *ctx, const char *msg, ...) {}

void _xmlSilentParseErrors() {
	xmlThrDefSetGenericErrorFunc(NULL, _nilGenericErrorHandler);
	xmlSetGenericErrorFunc(NULL, _nilGenericErrorHandler);
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

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

// Reference http://www.xmlsoft.org/examples/xpath1.c
// Code snippet from https://github.com/lestrrat-go/libxml2

// Evaluate returns the result of the expression.
// The result type of the expression is one of the follow: bool,float64,string,[]string (representing collections)).
func EvalXPathHTML(expr string, body [] byte) (interface{}, error) {

	C.xmlInitParser()
	defer C.xmlCleanupParser()

	C._xmlSilentParseErrors()

	exp := stringToXMLChar(expr)
	defer C.free(unsafe.Pointer(exp))

	enc := C.CString("UTF-8")
	defer C.free(unsafe.Pointer(enc))

	// Load HTLM document.
	doc := C.htmlParseDoc(bytesToXMLChar(body), enc)
	if doc == nil {
		return nil, errors.New("unable to parse HTML")
	}
	defer C.xmlFreeDoc(doc)

	// Create xpath evaluation context.
	ctx := C.xmlXPathNewContext(doc)
	if ctx == nil {
		return nil, errors.New("context creation failed")
	}
	defer C.xmlXPathFreeContext(ctx)

	// Evaluate xpath expression
	xobj := C.xmlXPathEvalExpression(exp, ctx)
	if xobj == nil {
		return nil, errors.New("XPath evaluation error")
	}
	defer C.xmlXPathFreeObject(xobj)

	switch t := XPathObjectType(xobj._type); t {
	case XPathStringType:
		return stringValue(xobj), nil
	case XPathBooleanType:
		return boolValue(xobj), nil
	case XPathNumberType:
		return float64Value(xobj), nil
	case XPathNodeSetType:
		return nil, errors.New("node set not supported")
	default:
		return nil, errors.New("unsupported xpath eval result")
	}
}

// Evaluate returns the result of the expression.
// The result type of the expression is one of the follow: bool,float64,string,[]string (representing collections)).
func EvalXPathXML(expr string, body []byte) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// stringToXMLChar creates a new *C.xmlChar from a Go string.
// Remember to always free this data, as C.CString creates a copy
// of the byte buffer contained in the string
func stringToXMLChar(s string) *C.xmlChar {
	return (*C.xmlChar)(unsafe.Pointer(C.CString(s)))
}

func xmlCharToString(s *C.xmlChar) string {
	return C.GoString((*C.char)(unsafe.Pointer(s)))
}

func bytesToXMLChar(b []byte) *C.xmlChar {
	return (*C.xmlChar)(unsafe.Pointer(C.CBytes(b)))
}

func stringValue(obj C.xmlXPathObjectPtr) string {
	return xmlCharToString(obj.stringval)
}

func boolValue(obj C.xmlXPathObjectPtr) bool {
	return obj.boolval != 0
}

func float64Value(obj C.xmlXPathObjectPtr) float64 {
	return float64(obj.floatval)
}
