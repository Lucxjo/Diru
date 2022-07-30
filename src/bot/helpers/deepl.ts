import deeplLangs from '../../shared/deepl.json' assert { type: 'json' };

type DeepLTranslateOptions = typeof deeplLangs;
const deepLTranslateOptions: DeepLTranslateOptions = deeplLangs;

export const autoLanguage = (usrLocale: string, targetLanguage: string) => {
	let langCode: string | null = null;

	if (targetLanguage === 'XX') {
		if (usrLocale.includes('-')) {
			let _langCode: string[] | string = usrLocale.split('-');
			if (_langCode[0].toUpperCase() === 'EN') {
				if (_langCode[1].toUpperCase() === 'US') {
					langCode = 'EN-US';
				} else {
					langCode = 'EN-GB';
				}
			} else if (_langCode[0].toUpperCase() === 'PT') {
				if (_langCode[1].toUpperCase() === 'BR') {
					langCode = 'PT-BR';
				} else {
					langCode = 'PT-PT';
				}
			} else {
				_langCode = _langCode[0].toUpperCase();
				deepLTranslateOptions.target.forEach((lang) => {
					if (lang.code === _langCode) {
						langCode = lang.code;
					}
				});
				if (!langCode) {
					langCode = 'EN-GB';
				}
			}
		}
	}

	return langCode;
};
