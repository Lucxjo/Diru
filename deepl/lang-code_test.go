package deepl

import (
	"testing"

	deeplgo "github.com/bounoable/deepl"
)

type testCodeStruct struct {
	langCode     string
	expected     bool
	expectedLang deeplgo.Language
}

var langs = []testCodeStruct{
	{"BG", true, deeplgo.Bulgarian},
	{"CS", true, deeplgo.Czech},
	{"DA", true, deeplgo.Danish},
	{"DE", true, deeplgo.German},
	{"EL", true, deeplgo.Greek},
	{"EN", true, deeplgo.EnglishBritish},
	{"EN-GB", true, deeplgo.EnglishBritish},
	{"EN-US", true, deeplgo.EnglishAmerican},
	{"ES", true, deeplgo.Spanish},
	{"ET", true, deeplgo.Estonian},
	{"FI", true, deeplgo.Finnish},
	{"FR", true, deeplgo.French},
	{"HU", true, deeplgo.Hungarian},
	{"IT", true, deeplgo.Italian},
	{"JA", true, deeplgo.Japanese},
	{"LT", true, deeplgo.Lithuanian},
	{"LV", true, deeplgo.Latvian},
	{"NL", true, deeplgo.Dutch},
	{"PL", true, deeplgo.Polish},
	{"PT", true, deeplgo.PortuguesePortugal},
	{"PT-BR", true, deeplgo.PortugueseBrazil},
	{"PT-PT", true, deeplgo.PortuguesePortugal},
	{"RO", true, deeplgo.Romanian},
	{"RU", true, deeplgo.Russian},
	{"SK", true, deeplgo.Slovak},
	{"SL", true, deeplgo.Slovenian},
	{"SV", true, deeplgo.Swedish},
	{"ZH", true, deeplgo.Chinese},
	{"EO", false, deeplgo.EnglishBritish},
	{"NB", true, NorwegianBokmal},
	{"ID", true, Indonesian},
	{"KO", true, Korean},
	{"TR", true, Turkish},
	{"UK", true, Ukrainian},
}

func TestCheckCode(t *testing.T) {

	for _, lang := range langs {
		if CheckCode(lang.langCode) != lang.expected {
			t.Errorf("CheckCode: expected %t, got %t", lang.expected, CheckCode(lang.langCode))
		} else {
			t.Logf("CheckCode: expected %t, got %t", lang.expected, CheckCode(lang.langCode))
		}
	}
}

func TestGetLang(t *testing.T) {

	for _, lang := range langs {
		if GetLang(lang.langCode) != lang.expectedLang {
			t.Errorf("GetLang: expected %s, got %s", lang.langCode, GetLang(lang.langCode))
		} else {
			t.Logf("GetLang: expected %s, got %s", lang.langCode, GetLang(lang.langCode))
		}
	}
}
