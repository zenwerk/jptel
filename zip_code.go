package jptel

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	postCodeLength = 3
)

var (
	zipCodeRegexString           = `^[0-9０-９]{7}$` // ハイフン無し
	zipCodeRegex                 = regexp.MustCompile(zipCodeRegexString)
	zipCodeWithHyphenRegexString = `^[0-9０-９]{3}[−ー-][0-9０-９]{4}$` // ハイフン有り
	zipCodeWIthHyphenRegex       = regexp.MustCompile(zipCodeWithHyphenRegexString)

	ErrInvalidZipCode = errors.New("zip code is invalid")
)

// JPZipCode は郵便番号1件を表す
type JPZipCode struct {
	PostCode string
	AreaCode string
}

// Format は郵便番号をハイフンで区切り返す
func (zc *JPZipCode) Format() string {
	return fmt.Sprintf("%s-%s", zc.PostCode, zc.AreaCode)
}

// SplitZipCode は入力された文字列を JPZipCode 構造体にして返す
func SplitZipCode(src string) (JPZipCode, error) {
	number, err := extractNumber(src)
	if err != nil {
		return JPZipCode{}, err
	}

	if len(number) != 7 {
		return JPZipCode{}, ErrInvalidZipCode
	}

	return JPZipCode{
		PostCode: number[:postCodeLength],
		AreaCode: number[postCodeLength:],
	}, nil
}

// ValidateZipCode は文字列が郵便番号として有効か検査する
func ValidateZipCode(src string) error {
	if !zipCodeRegex.MatchString(src) && !zipCodeWIthHyphenRegex.MatchString(src) {
		return ErrInvalidZipCode
	}

	return nil
}

// NormalizeZipCode は入力された文字列をハイフン区切りの郵便番号にして返す
func NormalizeZipCode(src string) (string, error) {
	number, err := extractNumber(src)
	if err != nil {
		return "", err
	}

	zipCode, err := SplitZipCode(number)
	if err != nil {
		return "", err
	}

	return zipCode.Format(), nil
}
