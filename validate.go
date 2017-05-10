package jptel

import (
	"errors"
	"regexp"
)

var (
	telephoneNumberRegexString           = `^[0-9０-９]{10,11}$` // ハイフン無し
	telephoneNumberRegex                 = regexp.MustCompile(telephoneNumberRegexString)
	telephoneNumberWithHyphenRegexString = `^[0-9０-９]{2,4}[ー-][0-9０-９]{2,4}[ー-][0-9０-９]{3,4}$` // ハイフン有り
	telephoneNumberWithHyphenRegex       = regexp.MustCompile(telephoneNumberWithHyphenRegexString)
)

// Validate validate telephone number
func Validate(src string) error {
	if !telephoneNumberRegex.MatchString(src) && !telephoneNumberWithHyphenRegex.MatchString(src) {
		return errors.New("telephone number is invalid")
	}

	number, err := extractNumber(src)
	if err != nil {
		return err
	}

	_, err = Split(number)
	return err
}
