package query
import "C"

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
import "C"

func EvalXPathXMLC() {
	C.xmlInitParser()
	C.xmlCleanupParser()
}