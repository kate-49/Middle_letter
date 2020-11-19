package letter

import "testing"

func TestReturnsMiddleLetter(t *testing.T) {

	t.Run("Odd value letters", func(t *testing.T) {
		got := getMiddle("testing")
		want := "t"

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
		// acceptance criteria tests
		{input: "test", want: "es"},
		{input: "middle", want: "dd"},
	}

	for _, tc := range tests {
		got := getMiddle(tc.input)
		if tc.want != got {
			t.Errorf("expected: %s, got: %s", tc.want, got)
		}
	}
}
