package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Actual:", teste)
	}
}

func TestSomaErro(t *testing.T) {
	teste := soma(1, 2, 3)
	resultado := 7
	if teste == resultado {
		t.Error("Expected:", resultado, "Actual:", teste)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma(1, 1)
	}
}
