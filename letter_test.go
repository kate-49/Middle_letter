package letter

import "testing"

func TestReturnsMiddleLetter(t *testing.T) {
	got := getMiddle("test")
	want := "es"

	if got != want {
		t.Errorf("expected %s, got %s", got, want)
	}
}
