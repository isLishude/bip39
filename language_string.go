// Code generated by "stringer -type=Language"; DO NOT EDIT.

package bip39

import "strconv"

const _Language_name = "ChineseSimplifiedChineseTraditionalEnglishFrenchItalianJapaneseKoreanSpanishCzech"

var _Language_index = [...]uint8{0, 17, 35, 42, 48, 55, 63, 69, 76, 81}

func (i Language) String() string {
	if i >= Language(len(_Language_index)-1) {
		return "Language(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Language_name[_Language_index[i]:_Language_index[i+1]]
}
