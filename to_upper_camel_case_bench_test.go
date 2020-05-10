package StringCase

import "testing"

func Benchmark_SnakeCaseToUpperCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		origin := "snake_case"
		except := "snakeCase"
		if s := ToUpperCamelCase(origin); s != except {
			b.Fatalf("snack to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
	}
}

func Benchmark_LowerCamelToUpperCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		origin := "lowerCamelCase"
		except := "lowerCamelCase"
		if s := ToUpperCamelCase(origin); s != except {
			b.Fatalf("lower camel case to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
	}
}

func Benchmark_UpperCamelToUpperCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		origin := "UpperCamelCase"
		except := "upperCamelCase"
		if s := ToUpperCamelCase(origin); s != except {
			b.Fatalf("upper camel case to lower camel case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
	}
}
