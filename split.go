package jptel

import (
	"errors"
	"strings"
)

// from: http://www.soumu.go.jp/main_sosiki/joho_tsusin/top/tel_number/number_shitei.html
//  0 + [xxxx]  ~  [xxxx] + [xxxx]
//  1 + <- total 5 num -> = 6

const (
	totalCodeLen = 6

	mobileCodePrefixLen = 3
	mobileCodeLen       = 4
	freeDialPrefixLen   = 4
	freeDialCodeLen     = 3
	otherCodePrefixLen  = 4
	otherCodeLen        = 3
)

var mobileCode = []string{
	"050", // IP電話
	"070", // 携帯電話/PHS
	"080", // 携帯電話
	"090", // 携帯電話
	"020", // その他
}

var freeDialCode = []string{
	"0120", // フリーダイヤル 0120-xxx-xxx
	"0800", // フリーダイヤル 0800-xxx-xxxx
}

var otherCode = []string{
	"0570", // ナビダイヤル 0570-xxx-xxx
	"0990", // ダイヤルQ2 0990-xxx-xxx
}

var areaCodes = [][]string{
	areaCode5,
	areaCode4,
	areaCode3,
	areaCode2,
}

var (
	ErrShort = errors.New("telephone number is too short")
	ErrMatch = errors.New("telephone number does not match any area code")
)

func generateJPTelephoneNumber(number string, areaCodeLength, cityCodeLength int) (JPTelephoneNumber, error) {
	var totalCodeLength = areaCodeLength + cityCodeLength
	if len(number) < totalCodeLength {
		return JPTelephoneNumber{}, ErrShort
	}

	return JPTelephoneNumber{
		AreaCode:       number[:areaCodeLength],
		CityCode:       number[areaCodeLength:totalCodeLength],
		SubscriberCode: number[totalCodeLength:],
	}, nil
}

// Split は入力された文字列を JPTelephoneNumber 構造体にして返す
func Split(src string) (JPTelephoneNumber, error) {
	number, err := ExtractNumber(src)
	if err != nil {
		return JPTelephoneNumber{}, err
	}

	// 固定電話
	for _, areaCode := range areaCodes {
		for _, code := range areaCode {
			if strings.HasPrefix(number, code) {
				areaCodeLength := len(code)
				cityCodeLength := totalCodeLen - areaCodeLength
				return generateJPTelephoneNumber(number, areaCodeLength, cityCodeLength)
			}
		}
	}

	// フリーダイヤル
	for _, code := range freeDialCode {
		if strings.HasPrefix(number, code) {
			return generateJPTelephoneNumber(number, freeDialPrefixLen, freeDialCodeLen)
		}
	}

	// 携帯番号
	for _, code := range mobileCode {
		if strings.HasPrefix(number, code) {
			return generateJPTelephoneNumber(number, mobileCodePrefixLen, mobileCodeLen)
		}
	}

	// その他番号
	for _, code := range otherCode {
		if strings.HasPrefix(number, code) {
			return generateJPTelephoneNumber(number, otherCodePrefixLen, otherCodeLen)
		}
	}

	return JPTelephoneNumber{}, ErrMatch
}
