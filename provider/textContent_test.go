package provider

import "testing"

func TestCanSetTextContent(t *testing.T) {
	content := &plainTextContent{}

	if err := content.SetBody("some text"); err != nil {
		t.Error(err)
	}
}

func TestCannotSetNonTextContent(t *testing.T) {
	content := &plainTextContent{}

	if err := content.SetBody([]string{"1", "2"}); err == nil {
		t.Error("expected to not fail in setting non text value")
	}
}
