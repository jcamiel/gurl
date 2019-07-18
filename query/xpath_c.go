package query

// #cgo CFLAGS: -I/opt/local/include/libxml2
// #cgo LDFLAGS: /opt/local/lib/libxml2.a
// #cgo LDFLAGS: /opt/local/lib/libz.a
// #cgo LDFLAGS: /opt/local/lib/libiconv.a
// #cgo LDFLAGS: /opt/local/lib/liblzma.a
// #include <stdlib.h>
// #include <stdio.h>
// #include <string.h>
// #include <assert.h>
//
// #include <libxml/tree.h>
// #include <libxml/parser.h>
// #include <libxml/xpath.h>
// #include <libxml/xpathInternals.h>
// #include <libxml/HTMLparser.h>
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

func EvalXPathXMLC(expr string, body [] byte) error {

	C.xmlInitParser()

	exp := stringToXMLChar(expr)
	defer C.free(unsafe.Pointer(exp))

	// Load HTLM document.
	doc := C.htmlParseDoc(bytesToXMLChar(body), nil)
	if doc == nil {
		return errors.New("unable to parse HTML")
	}
	defer C.xmlFreeDoc(doc)

	// Create xpath evaluation context.
	ctx := C.xmlXPathNewContext(doc)
	if ctx == nil {
		return errors.New("context creation failed")
	}
	defer C.xmlXPathFreeContext(ctx)

	xobj := C.xmlXPathEvalExpression(exp, ctx)
	if xobj == nil {
		return errors.New("Xpath evaluation error")
	}
	defer C.xmlXPathFreeObject(xobj)

	if xobj.nodesetval != nil {
		printXPathNodes(xobj.nodesetval)
	}

	C.xmlCleanupParser()
	return nil
}

// stringToXMLChar creates a new *C.xmlChar from a Go string.
// Remember to always free this data, as C.CString creates a copy
// of the byte buffer contained in the string
func stringToXMLChar(s string) *C.xmlChar {
	return (*C.xmlChar)(unsafe.Pointer(C.CString(s)))
}

func bytesToXMLChar(b []byte) *C.xmlChar {
	return (*C.xmlChar)(unsafe.Pointer(C.CBytes(b)))
}

func printXPathNodes(nodeset C.xmlNodeSetPtr) {

	size := int(nodeset.nodeNr)
	fmt.Printf("Result (%d nodes):\n", size)
	if size == 0 {
		return
	}

	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(nodeset.nodeTab)),
		Len:  size,
		Cap:  size,
	}
	nodes := *(*[]*C.xmlNode)(unsafe.Pointer(&hdr))
	for i := 0; i < size; i++ {
		node := nodes[i]
		fmt.Printf("node %d: type:%d\n",i, node._type)
	}

}
