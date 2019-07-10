package template

import (
	"errors"
	"fmt"
	"regexp"
)

func Render(template string, variables map[string]string) (string, error) {
	r := regexp.MustCompile(`\{\{(.*?)\}\}`)
	s := template
	for {
		m := r.FindStringSubmatchIndex(s)
		if m == nil {
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
