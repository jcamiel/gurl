package template

import (
	"errors"
	"fmt"
	"regexp"
)

// Render the variable inside the template string. If the template
// string doesn't contains any tempating, the template string is returned.
// If there is any undefined variable ({{x}} where x is not a key in variables)
// and error is returned.
func Render(template string, variables map[string]string) (string, error) {
	r := regexp.MustCompile(`\{\{(.*?)\}\}`)
	s := template
	for {
		m := r.FindStringSubmatchIndex(s)
		if m == nil {
			// No more template variable to render, we exit and return the current string.
			break
		}
		begin, end := m[2], m[3]
		name := s[begin:end]
		value, ok := variables[name]
		if !ok {
			return "", errors.New(fmt.Sprintf("undefined variable '%s'", name))
		}
		s = fmt.Sprintf("%s%s%s", s[:begin-2], value, s[end+2:])
	}
	return s, nil
}
