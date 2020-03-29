package StringCase

import "testing"

func TestToSnakeCaseSnakeToSnake(t *testing.T) {
	origin := "snake_case"
	except := "snake_case"
	if s := ToSnakeCase(origin); s != except {
		t.Fatalf("snack to snack but string changed\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func TestToSnakeCaseLowerCamelToSnake(t *testing.T) {
	origin := "lowerCamelCase"
	except := "lower_camel_case"
	if s := ToSnakeCase(origin); s != except {
		t.Fatalf("lower camel case to snack but string changed\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}

func TestToSnakeCaseUpperCamelToSnake(t *testing.T) {
	origin := "UpperCamelCase"
	except := "upper_camel_case"
	if s := ToSnakeCase(origin); s != except {
		t.Fatalf("upper camel case to snack but string changed\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
	}
}
