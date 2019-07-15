package print

import "gurl/ast"

type Printer interface {
	Print(hurlFile *ast.HurlFile) string
}
