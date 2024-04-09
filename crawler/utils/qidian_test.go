package utils

import "testing"

func TestReplacePlaceholer(t *testing.T) {

	asserCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // needed to tell the test suite that this method is  a helper
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("test1", func(t *testing.T) {
		input := "Hello, {name}! You have {count} new messages."
		want := "Hello, 1! You have 2 new messages."
		res := ReplacePlaceholer("{name}", input, "1")
		res = ReplacePlaceholer("{count}", res, "2")
		asserCorrectMessage(t, res, want)
	})

	t.Run("test1", func(t *testing.T) {
		filename := "../test"
		SaveAsJson(filename, "test")
		// res = ReplacePlaceholer("{count}", res, "2")
		// asserCorrectMessage(t, res, want)
	})
}
