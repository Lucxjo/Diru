package deepl

import (
	"context"

	"github.com/bounoable/deepl"
)

// AutoTranslate translates a string from an unknown language using the DeepL API into English (British).
func AutoTranslate(phrase string, client *deepl.Client) string {
	translated, _, err := client.Translate(
		context.TODO(),
		phrase,
		deepl.EnglishBritish,
	)

	if err != nil {
		panic(err)
	}

	return translated
}