package StringCase

import (
	"strings"
)

/*
string to snake case
example:
snake_case ToSnakeCase snake_case
lowerCamelCase ToSnakeCase lower_camel_case
UpperCamelCase ToSnakeCase upper_camel_case
*/
func ToSnakeCase(s string) string {
	rs := []rune(s)
	b := strings.Builder{}
	b.Grow(len(rs))

	for i := 0; i < len(rs); i++ {
		if IsUpper(rs[i]) {
			begin := i
			end := i

			for ; i < len(rs); i++ {
				if IsUpper(rs[i]) {
					end = i
				} else {
					break
				}
			}

			if begin != 0 {
				b.WriteRune('_')
			}

			b.WriteString(strings.ToLower(string(rs[begin : end+1])))
			i--
			continue
		}
		b.WriteRune(rs[i])
	}

	return b.String()
}
