package run

import (
	"errors"
	"fmt"
	"gurl/ast"
)

type AssertResult struct {
	ok  bool
	msg string
}

func (a *AssertResult) String() string {
	if a.ok {
		return "success"
	} else {
		return fmt.Sprintf("failed, %s", a.msg)
	}
}

func float64Val(p *ast.Predicate) (float64, error) {
	if p.Float != nil {
		return p.Float.Value, nil
	} else if p.Integer != nil {
		return float64(p.Integer.Value), nil
	}
	return 0, errors.New("invalid conversion to float64")
}

func boolVal(p *ast.Predicate) (bool, error) {
	if p.Bool != nil {
		return p.Bool.Value, nil
	}
	return false, errors.New("invalid conversion to bool")
}

func stringVal(p *ast.Predicate) (string, error) {
	if p.String != nil {
		return p.String.Value, nil
	}
	return "", errors.New("invalid conversion to string")
}

func val(p *ast.Predicate) interface{} {
	if v := p.Integer; v != nil {
		return v.Value
	}
	if v := p.Float; v != nil {
		return v.Value
	}
	if v := p.Bool; v != nil {
		return v.Value
	}
	if v := p.String; v != nil {
		return v.Value
	}
	return nil
}

func assertEquals(pred *ast.Predicate, actual interface{}) *AssertResult {
	var ok bool
	var msg string

	switch actual.(type) {
	case float64:
		expected, err := float64Val(pred)
		if err != nil {
			ok = false
		} else {
			ok = actual == expected
		}
	case bool:
		expected, err := boolVal(pred)
		if err != nil {
			ok = false
		} else {
			ok = actual == expected
		}
	case string:
		expected, err := stringVal(pred)
		if err != nil {
			ok = false
		} else {
			ok = actual == expected
		}
	default:
		ok = false
	}
	if !ok {
		msg = fmt.Sprintf("actual: %v expected: %v", actual, val(pred))
	}

	return &AssertResult{ok, msg}
}
