package letter

import "testing"

func TestAcceptanceCriteria(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// even tests
		{input: "test", want: "es"},
		{input: "middle", want: "dd"},
		//odd tests
		{input: "testing", want: "t"},
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
