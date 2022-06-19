package deepl

import (
	"os"
	"testing"

	deeplgo "github.com/bounoable/deepl"
)

func TestTranslateTo(t *testing.T) {
	dpl_token := os.Getenv("DEEPL_TOKEN")

	client := deeplgo.New(dpl_token)
	phrase := "Hola Mundo! Cómo estáis?"
	expected := "Hello World! How are you?"
	actual := TranslateTo("en", phrase, client)
	if actual != expected {
		t.Errorf("AutoTranslate: expected %s, got %s", expected, actual)
	} else {
		t.Logf("AutoTranslate: expected %s, got %s", expected, actual)
	}
}
