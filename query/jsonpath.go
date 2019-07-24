package query

import (
	"encoding/json"
	"errors"
	"github.com/PaesslerAG/jsonpath"
	"strings"
)

func EvalJSONPath(expr string, body [] byte) (interface{}, error) {

	// jsonpath supports only double quoted expression, so we
	// replace any single quote by a double quote
	expr = strings.ReplaceAll(expr, `'`, `"`)

	var data interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	v, err := jsonpath.Get(expr, data)
	switch v.(type) {
	case []interface{}:
		return nil, errors.New("node set not supported")
	case string, bool, float64:
		return v, nil
	default:
		return nil, errors.New("unsupported jsonpath eval result")
	}
}

func IsJSON(body [] byte) bool {
	var data interface{}
	err := json.Unmarshal(body, &data)
	return err == nil
}