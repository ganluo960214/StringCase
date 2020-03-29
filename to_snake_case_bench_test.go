package StringCase

import "testing"

func BenchmarkToSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		origin := "UpperCamelCase"
		except := "upper_camel_case"
		if s := ToSnakeCase(origin); s != except {
			b.Fatalf("upper camel case to snack but string changed\norigin:%s\nexcept:%s\nnow:%s", origin, except, s)
		}
		b.ReportAllocs()
	}
}
