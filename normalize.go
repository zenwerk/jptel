package jptel

import (
	"errors"
	"strings"
	"regexp"
)

var replacer = strings.NewReplacer(
	"０", "0",
	"１", "1",
	"２", "2",
	"３", "3",
	"４", "4",
	"５", "5",
	"６", "6",
	"７", "7",
	"８", "8",
	"９", "9",
	"ー", "-",
	"−", "-",
)

const (
	numberAndHyphenRegexpString = `^[−ー0-9０-９-]+$`
)

var (
	numberAndHyphenRegexp = regexp.MustCompile(numberAndHyphenRegexpString)

	ErrContainsInvalidCharactor = errors.New("src string must be number or hyphen")
)

func zenkakuToHankaku(src string) string {
	return replacer.Replace(src)
}

// ExtractNumber は全角半角やハイフンが混じった文字列を半角数字のみの文字列にして返す
func ExtractNumber(src string) (string, error) {
	if !numberAndHyphenRegexp.MatchString(src) {
		return "", ErrContainsInvalidCharactor
	}
	return strings.Replace(zenkakuToHankaku(src), "-", "", -1), nil
}

// Normalize は入力された文字列をハイフン区切りの電話番号にして返す
func Normalize(src string) (string, error) {
	number, err := ExtractNumber(src)
	if err != nil {
		return "", err
	}

	phoneNumber, err := Split(number)
	if err != nil {
		return "", err
	}

	return phoneNumber.Format(), nil
}
