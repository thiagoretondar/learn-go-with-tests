package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, got string) {
		t.Helper()

		if expected != got {
			t.Errorf("expected %q but got %q", expected, got)
		}
	}

	t.Run("return same string if count is zero", func(t *testing.T) {
		repeated := Repeat("a", 0)

		expected := "a"

		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("return same string if count is negative", func(t *testing.T) {
		repeated := Repeat("a", -10)

		expected := "a"

		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("return string repeated N times", func(t *testing.T) {
		repeated := Repeat("a", 10)

		expected := "aaaaaaaaaa"

		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
