package StringCase

import "testing"

func Benchmark_SnakeCaseToSnakeCase(b *testing.B) {
	origin := "snake_case"
	except := "snake_case"
	for i := 0; i < b.N; i++ {
		if s := ToSnakeCase(origin); s != except {
			b.Fatalf("snack case to snack case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
	}
	b.ReportAllocs()
}

func Benchmark_LowerCamelToSnakeCase(b *testing.B) {
	origin := "lowerCamelCase"
	except := "lower_camel_case"
	for i := 0; i < b.N; i++ {
		if s := ToSnakeCase(origin); s != except {
			b.Fatalf("lower camel case to snack case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
	}
	b.ReportAllocs()
}

func Benchmark_UpperCamelToSnakeCase(b *testing.B) {
	origin := "UpperCamelCase"
	except := "upper_camel_case"
	for i := 0; i < b.N; i++ {
		if s := ToSnakeCase(origin); s != except {
			b.Fatalf("upper camel case to snack case\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
	}
	b.ReportAllocs()
}
