package string

import (
	"strings"
	"testing"
)

func BenchmarkPus(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := "H"
		for s := 0; s < 200; s++ {
			a += "H"
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := make([]string, 0, 200)
		for s := 0; s < 200; s++ {
			arr = append(arr, "H")
		}
		strings.Join(arr, "")
	}
}

func BenchmarkBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a := strings.Builder{}
		a.Grow(200)
		for s := 0; s < 200; s++ {
			a.WriteString("H")
		}
		a.String()
	}
}
