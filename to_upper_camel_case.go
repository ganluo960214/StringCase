package StringCase

import (
	"strings"
)

func ToUpperCamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	rs := []rune(s)
	b := strings.Builder{}
	b.WriteString(strings.ToUpper(string(rs[0])))
	rs = rs[1:]

	isSnakeCase := false
	for i := 0; i < len(rs); i++ {
		if isSnakeCase && IsLower(rs[i]) {
			b.WriteString(strings.ToUpper(string(rs[i])))
			isSnakeCase = false
			continue
		}
		if rs[i] == '_' { // find snake case
			isSnakeCase = true
		} else {
			isSnakeCase = false
			b.WriteRune(rs[i])
		}
	}

	return b.String()
}
