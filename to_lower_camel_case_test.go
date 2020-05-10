package StringCase

import (
	"testing"
)

func Test_SnakeCaseToLowerCamelCase(t *testing.T) {
	origin := "snake_case"
	except := "snakeCase"
	if s := ToLowerCamelCase(origin); s != except {
		t.Fatalf("snack to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func Test_LowerCamelToLowerCamelCase(t *testing.T) {
	origin := "lowerCamelCase"
	except := "lowerCamelCase"
	if s := ToLowerCamelCase(origin); s != except {
		t.Fatalf("lower camel case to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func Test_UpperCamelToLowerCamelCase(t *testing.T) {
	origin := "UpperCamelCase"
	except := "upperCamelCase"
	if s := ToLowerCamelCase(origin); s != except {
		t.Fatalf("upper camel case to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}
