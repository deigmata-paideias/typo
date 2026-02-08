package common

import (
	"strings"
)

type SwitchLangRule struct{}

func (r *SwitchLangRule) ID() string { return "switch_lang" }

var (
	qwerty = "qwertyuiop[]asdfghjkl;'zxcvbnm,./QWERTYUIOP{}ASDFGHJKL:\"ZXCVBNM<>?"

	// Russian
	ru = "йцукенгшщзхъфывапролджэячсмитьбю.ЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮ,"

	// Greek
	gr = "ςερτυθιοπ[]ασδφγηξκλ΄ζχψωβνμ,./:΅ΕΡΤΥΘΙΟΠ{}ΑΣΔΦΓΗΞΚΛ¨\"ΖΧΨΩΒΝΜ<>?"

	// Korean
	ko = "ㅂㅈㄷㄱㅅㅛㅕㅑㅐㅔ[]ㅁㄴㅇㄹㅎㅗㅓㅏㅣ;'ㅋㅌㅊㅍㅠㅜㅡ,./ㅃㅉㄸㄲㅆㅛㅕㅑㅒㅖ{}ㅁㄴㅇㄹㅎㅗㅓㅏㅣ:\"ㅋㅌㅊㅍㅠㅜㅡ<>?"
)

func (r *SwitchLangRule) Match(command string, output string) bool {
	layouts := []string{ru, gr, ko}
	for _, layout := range layouts {
		if isLayout(command, layout) {
			return true
		}
	}
	return false
}

func isLayout(cmd string, layout string) bool {
	hasLayoutChar := false
	for _, r := range cmd {
		if strings.ContainsRune(" -_0123456789", r) {
			continue
		}
		if !strings.ContainsRune(layout, r) {
			return false
		}
		hasLayoutChar = true
	}
	// Must have at least one character from the layout to prevent matching just "123"
	return hasLayoutChar
}

func (r *SwitchLangRule) GetNewCommand(command string, output string) string {
	if isLayout(command, ru) {
		return convertLayout(command, ru, qwerty)
	}
	if isLayout(command, gr) {
		return convertLayout(command, gr, qwerty)
	}
	if isLayout(command, ko) {
		return convertLayout(command, ko, qwerty)
	}
	return command
}

func convertLayout(s, src, dst string) string {
	// Create map
	m := make(map[rune]rune)
	srcRunes := []rune(src)
	dstRunes := []rune(dst)

	limit := len(srcRunes)
	if len(dstRunes) < limit {
		limit = len(dstRunes)
	}

	for i := 0; i < limit; i++ {
		m[srcRunes[i]] = dstRunes[i]
	}

	var res strings.Builder
	for _, r := range s {
		if target, ok := m[r]; ok {
			res.WriteRune(target)
		} else {
			res.WriteRune(r)
		}
	}
	return res.String()
}
