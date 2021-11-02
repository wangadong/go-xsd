package util

import (
	"bytes"
	"strings"
	"unicode"
)

//	Prepends `prefix + sep` to `v` only if `prefix` isn't empty.
func PrefixWithSep(prefix, sep, v string) string {
	if len(prefix) > 0 {
		return prefix + sep + v
	}
	return v
}

//	Returns an empty slice is `v` is emtpy, otherwise like `strings.Split`
func Split(v, sep string) (sl []string) {
	if len(v) > 0 {
		sl = strings.Split(v, sep)
	}
	return
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifs(cond bool, ifTrue, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns `ifTrue` if `cond` is `true`, otherwise returns `ifFalse`.
func Ifm(cond bool, ifTrue, ifFalse map[string]string) map[string]string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Prepends `p` to `s` only if `s` doesn't already have that prefix.
func PrependIf(s, p string) string {
	if strings.HasPrefix(s, p) {
		return s
	}
	return p + s
}

//	A most simplistic (not linguistically-correct) English-language pluralizer that may be useful for code or doc generation.
//
//	If `s` ends with "s", only appends "es": bus -> buses, mess -> messes etc.
//
//	If `s` ends with "y" (but not "ay", "ey", "oy", "uy" or "iy"), removes "y" and appends "ies": autonomy -> autonomies, dictionary -> dictionaries etc.
//
//	Otherwise, appends "s": gopher -> gophers, laptop -> laptops etc.
func Pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s + "es"
	}
	if (len(s) > 1) && strings.HasSuffix(s, "y") && !IsOneOf(s[(len(s)-2):], "ay", "ey", "oy", "uy", "iy") {
		return s[0:(len(s)-1)] + "ies"
	}
	return s + "s"
}

//	Returns whether `s` is in `all`.
func IsOneOf(s string, all ...string) bool {
	for _, a := range all {
		if s == a {
			return true
		}
	}
	return false
}

//	Returns the position of `val` in `slice`.
func StrAt(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

//	Returns whether `one` and `two` only contain identical values, regardless of ordering.
func StrEquivalent(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	for _, v := range one {
		if StrAt(two, v) < 0 {
			return false
		}
	}
	return true
}

//	Replaces in `str` all occurrences of all `repls` hash-map keys with their respective associated (mapped) value.
func Replace(str string, repls map[string]string) string {
	for k, v := range repls {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}

//	Creates a Pascal-cased "identifier" version of the specified string.
func SafeIdentifier(s string) string {
	var (
		isL, isD, last bool
		buf            bytes.Buffer
	)
	for i, r := range s {
		if isL, isD = unicode.IsLetter(r), unicode.IsDigit(r); isL || isD || ((r == '_') && (i == 0)) {
			if (i > 0) && (isL != last) {
				buf.WriteRune(' ')
			}
			buf.WriteRune(r)
		} else {
			buf.WriteRune(' ')
		}
		last = isL
	}
	words := Split(strings.Title(buf.String()), " ")
	for i, w := range words {
		if (len(w) > 1) && IsUpper(w) {
			words[i] = strings.Title(strings.ToLower(w))
		}
	}
	return strings.Join(words, "")
}

//	Returns whether all `unicode.IsLetter` runes in `s` are upper-case.
func IsUpper(s string) bool {
	for _, r := range s {
		if !((unicode.IsLetter(r) && unicode.IsUpper(r)) || unicode.IsSpace(r) || unicode.IsDigit(r) || unicode.IsNumber(r)) {
			return false
		}
	}
	return true
}
