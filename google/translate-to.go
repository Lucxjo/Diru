package google

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func TranslateTo(target string, phrase string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(target)

	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)

	if err != nil {
		return "", err
	}

	defer client.Close()

	resp, err := client.Translate(ctx, []string{phrase}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", phrase)
	}
	return resp[0].Text, nil

}
