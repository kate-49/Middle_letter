package letter

import "testing"

func TestReturnsMiddleLetter(t *testing.T) {

	t.Run("Odd value letters", func(t *testing.T) {
		got := getMiddle("A")
		want := "A"

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	})
}
func TestEvenValueLetters(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// even tests
		{input: "test", want: "es"},
		{input: "middle", want: "dd"},
		// edgecases tests
		{input: "A", want: "A"},
		{input: "of", want: "of"},
	}

	for _, tc := range tests {
		got := getMiddle(tc.input)
		if tc.want != got {
			t.Errorf("expected: %s, got: %s", tc.want, got)
		}
	}
}
