package cli

import (
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := Configuration{
		Keywords: []string{"autumn"},
		Span:     true,
	}
	got := Parse([]string{"autumn", "Span"})

	if !reflect.DeepEqual(expected.Keywords, got.Keywords) {
		t.Errorf("expected Keywords %s but got %s", expected.Keywords, got.Keywords)
	}

	if got.Span != expected.Span {
		t.Errorf("expected Span %t but got %t", expected.Span, got.Span)
	}
}
