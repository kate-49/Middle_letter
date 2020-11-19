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

	t.Run("Even value letters", func(t *testing.T) {
		got := getMiddle("test")
		want := "es"

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	})

}
