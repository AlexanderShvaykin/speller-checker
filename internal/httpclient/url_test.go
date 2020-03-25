package httpclient

import (
	"testing"
)

func TestBuildURL(t *testing.T) {
	expected := "localhost?Foo=Baz"
	base := "localhost"
	params := map[string]string{
		"Foo": "Baz",
	}
	result := BuildURL(base, params)

	if result != expected {
		t.Errorf("unexpected result: got %s, want %s", result, expected)
	}
}
