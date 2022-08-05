import { Static, Type } from '@sinclair/typebox';

export const DeepL = Type.Object({
	text: Type.String() || Type.Undefined(),
	KEY: Type.String() || Type.Undefined(),
	LANG_CODE: Type.String() || Type.Undefined(),
});

export type DeepLType = Static<typeof DeepL>;
