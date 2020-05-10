package StringCase

import "testing"

func Test_SnakeCaseToSnakeCase(t *testing.T) {
	origin := "snake_case"
	except := "snake_case"
	if s := ToSnakeCase(origin); s != except {
		t.Fatalf("snack case to snack case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func Test_LowerCamelToSnakeCase(t *testing.T) {
	origin := "lowerCamelCase"
	except := "lower_camel_case"
	if s := ToSnakeCase(origin); s != except {
		t.Fatalf("lower camel case to snack case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func Test_UpperCamelToSnakeCase(t *testing.T) {
	origin := "UpperCamelCase"
	except := "upper_camel_case"
	if s := ToSnakeCase(origin); s != except {
		t.Fatalf("upper camel case to snack case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}
