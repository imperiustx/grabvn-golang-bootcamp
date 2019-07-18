package main

import "testing"

func TestArithmetic(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test add", func(t *testing.T) {
		got := Arithmetic("2 + 3")
		want := "2 + 3 = 5"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test sub", func(t *testing.T) {
		got := Arithmetic("999 - 123")
		want := "999 - 123 = 876"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test mul", func(t *testing.T) {
		got := Arithmetic("9 * 11")
		want := "9 * 11 = 99"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test div", func(t *testing.T) {
		got := Arithmetic("10 / 5")
		want := "10 / 5 = 2"
		assertCorrectMessage(t, got, want)
	})

}

func BenchmarkArithmetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Arithmetic("1 + 2")
	}
}
