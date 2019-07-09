package run

import (
	"fmt"
	"gurl/ast"
)

type HttpRunner struct {
	text string
}

func NewHttpRunner() *HttpRunner {
	return &HttpRunner{}
}

func (h *HttpRunner) Run(hurlFile *ast.HurlFile) {
	ast.Walk(h, hurlFile)
}

func (h *HttpRunner) Visit(node ast.Noder) ast.Visitor {
	switch n := node.(type) {
	case *ast.Request:
		fmt.Printf("Do request %v\n", n)
		return nil
	}
	return h
}