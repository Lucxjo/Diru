package deepl

import (
	"strings"

	"github.com/bounoable/deepl"
)

// GetLang returns the deepl language for a given language code.
func GetLang(langCode string) deepl.Language {
	trimmed := strings.TrimSpace(strings.ToUpper(langCode))
	switch trimmed {
		case "BG":
			return deepl.Bulgarian
		case "CS":
			return deepl.Czech
		case "DA":
			return deepl.Danish
		case "DE":
			return deepl.German
		case "EL":
			return deepl.Greek
		case "EN":
			return deepl.EnglishBritish
		case "EN-GB":
			return deepl.EnglishBritish
		case "EN-US":
			return deepl.EnglishAmerican
		case "ES":
			return deepl.Spanish
		case "ET":
			return deepl.Estonian
		case "FI":
			return deepl.Finnish
		case "FR":
			return deepl.French
		case "HU":
			return deepl.Hungarian
		case "IT":
			return deepl.Italian
		case "JA":
			return deepl.Japanese
		case "LT":
			return deepl.Lithuanian
		case "LV":
			return deepl.Latvian
		case "NL":
			return deepl.Dutch
		case "PL":
			return deepl.Polish
		case "PT":
			return deepl.PortuguesePortugal
		case "PT-BR":
			return deepl.PortugueseBrazil
		case "PT-PT":
			return deepl.PortuguesePortugal
		case "RO":
			return deepl.Romanian
		case "RU":
			return deepl.Russian
		case "SK":
			return deepl.Slovak
		case "SL":
			return deepl.Slovenian
		case "SV":
			return deepl.Swedish
		case "ZH":
			return deepl.Chinese
		default:
			return deepl.EnglishBritish
	}
}

// CheckCode checks if a language code is valid.
func CheckCode(langCode string) bool {
	trimmed := strings.TrimSpace(strings.ToUpper(langCode))

	switch trimmed {
		case "BG":
			return true
		case "CS":
			return true
		case "DA":
			return true
		case "DE":
			return true
		case "EL":
			return true
		case "EN":
			return true
		case "EN-GB":
			return true
		case "EN-US":
			return true
		case "ES":
			return true
		case "ET":
			return true
		case "FI":
			return true
		case "FR":
			return true
		case "HU":
			return true
		case "IT":
			return true
		case "JA":
			return true
		case "LT":
			return true
		case "LV":
			return true
		case "NL":
			return true
		case "PL":
			return true
		case "PT":
			return true
		case "PT-BR":
			return true
		case "PT-PT":
			return true
		case "RO":
			return true
		case "RU":
			return true
		case "SK":
			return true
		case "SL":
			return true
		case "SV":
			return true
		case "ZH":
			return true
		default:
			return false
	}

}
