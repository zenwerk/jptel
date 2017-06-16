package jptel

import (
	"errors"
	"regexp"
)

var (
	telephoneNumberRegexString           = `^[0-9０-９]{10,11}$` // ハイフン無し
	telephoneNumberRegex                 = regexp.MustCompile(telephoneNumberRegexString)
	telephoneNumberWithHyphenRegexString = `^[0-9０-９]{2,4}[−ー-][0-9０-９]{2,4}[−ー-][0-9０-９]{3,4}$` // ハイフン有り
	telephoneNumberWithHyphenRegex       = regexp.MustCompile(telephoneNumberWithHyphenRegexString)

	ErrInvalidNumber = errors.New("telephone number is invalid")
)

// Validate は入力された文字列が電話番号として有効か検査する
func Validate(src string) error {
	if !telephoneNumberRegex.MatchString(src) && !telephoneNumberWithHyphenRegex.MatchString(src) {
		return ErrInvalidNumber
	}

	number, err := ExtractNumber(src)
	if err != nil {
		return err
	}

	_, err = Split(number)
	return err
}
