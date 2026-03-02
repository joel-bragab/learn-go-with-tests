package iteration

import (
	"strings"
	"testing"
)

const repeatCount = 5

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func Repeat(character string) string {
	var repeated strings.Builder
	// alternative for range repeatCount
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
