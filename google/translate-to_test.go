package google

import "testing"

func TestTranslateTo(t *testing.T) {
	t.Parallel()
	phrase := "Hola Mundo! Cómo estáis?"
	target := "en"
	expected := "Hello World! How are you?"
	actual, err := TranslateTo(target, phrase)
	if err != nil {
		t.Errorf("TranslateTo: %v", err)
	} else if actual != expected {
		t.Errorf("TranslateTo: expected %s, got %s", expected, actual)
	} else {
		t.Logf("TranslateTo: expected %s, got %s", expected, actual)
	}
}