package deepl

import (
	"context"

	"github.com/bounoable/deepl"
)


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