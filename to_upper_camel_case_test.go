package StringCase

import (
	"testing"
)

func Test_SnakeCaseToUpperCamelCase(t *testing.T) {
	origin := "snake_case"
	except := "snakeCase"
	if s := ToUpperCamelCase(origin); s != except {
		t.Fatalf("snack to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func Test_LowerCamelToUpperCamelCase(t *testing.T) {
	origin := "lowerCamelCase"
	except := "lowerCamelCase"
	if s := ToUpperCamelCase(origin); s != except {
		t.Fatalf("lower camel case to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func Test_UpperCamelToUpperCamelCase(t *testing.T) {
	origin := "UpperCamelCase"
	except := "upperCamelCase"
	if s := ToUpperCamelCase(origin); s != except {
		t.Fatalf("upper camel case to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}
