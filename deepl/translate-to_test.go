package deepl

import (
	"os"
	"testing"

	deeplgo "github.com/bounoable/deepl"
	_ "github.com/joho/godotenv/autoload"
)

func TestTranslateTo(t *testing.T) {
	t.Parallel()

	dpl_token := os.Getenv("DEEPL_TOKEN")

	client := deeplgo.New(dpl_token)
	expected := "Hello world! How are you?"

	for _, phrase := range phrases {
		actual := TranslateTo("en", phrase, client)
		if actual != expected {
			t.Errorf("AutoTranslate: expected %s, got %s", expected, actual)
		} else {
			t.Logf("AutoTranslate: expected %s, got %s", expected, actual)
		}
	}
}
