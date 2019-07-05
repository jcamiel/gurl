package print

import (
	"encoding/json"
	"gurl/ast"
)

type JSONPrinter struct {
}

func NewJSONPrinter() *JSONPrinter {
	return &JSONPrinter{}
}

func (p *JSONPrinter) Print(hurlFile *ast.HurlFile) string {
	b, _ := json.Marshal(hurlFile)
	return string(b)
}
