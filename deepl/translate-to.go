package deepl

import (
	"context"

	"github.com/bounoable/deepl"
)

// TranslateTo translates a string from an unknown language using the DeepL API into a requested language (Default: British English).
func TranslateTo(lang string, phrase string, client *deepl.Client) string {
	translated, _, err := client.Translate(
		context.TODO(),
		phrase,
		GetLang(lang),
	)

	if err != nil {
		panic(err)
	}

	return translated
}
