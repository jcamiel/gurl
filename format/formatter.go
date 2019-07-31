package format

import "gurl/ast"

type Formatter interface {
	Format(hurlFile *ast.HurlFile) string
}
