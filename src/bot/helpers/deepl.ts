import deeplLangs from '../../shared/deepl.json' assert { type: 'json' };

type DeepLTranslateOptions = typeof deeplLangs;
const deepLTranslateOptions: DeepLTranslateOptions = deeplLangs;

export const autoLanguage = (usrLocale: string) => {
	let langCode: string = '';
	let defaultCode = false;

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
				defaultCode = true;
			}
		}
	} else {
		if (usrLocale.toUpperCase() === 'EN') {
			langCode = 'EN-GB';
		} else if (usrLocale.toUpperCase() === 'PT') {
			langCode = 'PT-PT';
		} else {
			deepLTranslateOptions.target.forEach((lang) => {
				if (lang.code.toUpperCase() === usrLocale.toUpperCase()) {
					langCode = lang.code;
				}
			});
			if (!langCode) {
				langCode = 'EN-GB';
				defaultCode = true;
			}
		}
	}

	return { langCode, defaultCode };
};

export const verifyLanguage = (lang: string) => {
	let l: string = '';
	let langCode: string | null = null;
	let defaultCode = false;
	if (lang.includes('-') || lang.includes('_')) {
		let _langCode = lang.toUpperCase();
		if (_langCode.includes('_')) {
			_langCode = _langCode.replace('_', '-');
		}
		({langCode, defaultCode} = autoLanguage(_langCode));
	} else if (lang.length === 2) {
		({langCode, defaultCode} = autoLanguage(lang.toUpperCase()));
	} else {
		deepLTranslateOptions.target.forEach(({ name, code, aliases }) => {
			if (name.toUpperCase() === lang.toUpperCase()) {
				l = code;
			} else if (aliases) {
				aliases.forEach((alias) => {
					if (alias.toUpperCase() === lang.toUpperCase()) {
						l = code;
					}
				})
			} 
		});
		console.log(`bot/helpers/deepl.ts:62 ${l}`);
		({langCode, defaultCode} = autoLanguage(l));
	}

	return { langCode, defaultCode };
};
