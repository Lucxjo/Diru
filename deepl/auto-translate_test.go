package deepl

import (
	deeplgo "github.com/bounoable/deepl"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"testing"
)

var phrases = []string{
	"Hola mundo! Cómo estáis?",
	"Hallo verden! Hvordan har du det?",
	"Bonjour le monde! Comment allez-vous?",
	"Ciao mondo! Come stai?",
	"Olá mundo! Como vai você?",
	"こんにちは世界！元気ですか？",
}

func TestAutoTranslate(t *testing.T) {
	t.Parallel()
	dpl_token := os.Getenv("DEEPL_TOKEN")

	client := deeplgo.New(dpl_token)
	expected := "Hello world! How are you?"

	for _, phrase := range phrases {
		actual := AutoTranslate(phrase, client)
		if actual != expected {
			t.Errorf("AutoTranslate: expected %s, got %s", expected, actual)
		} else {
			t.Logf("AutoTranslate: expected %s, got %s", expected, actual)
		}
	}
}
