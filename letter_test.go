package letter

import "testing"

func TestReturnsMiddleLetter(t *testing.T) {
	got := getMiddle("testing")
	want := "t"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
