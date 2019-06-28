package format

import (
	"encoding/json"
	"gurl/ast"
)

type JSONFormatter struct {
}

func (p *JSONFormatter) ToText(hurlFile *ast.HurlFile) string {
	b, _ := json.Marshal(hurlFile)
	return string(b)
}
