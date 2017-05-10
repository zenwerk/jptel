package jptelsplit

import (
	"strings"
	"github.com/pkg/errors"
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
	OtherCodePrefixLen  = 4
	OtherCodeLen        = 3
)

var MobileCode = []string{
	"050", // IP電話
	"070", // 携帯電話/PHS
	"080", // 携帯電話
	"090", // 携帯電話
	"020", // その他
}

var FreeDialCode = []string{
	"0120", // フリーダイヤル 0120-xxx-xxx
	"0800", // フリーダイヤル 0800-xxx-xxxx
}

var OtherCode = []string{
	"0570", // ナビダイヤル 0570-xxx-xxx
	"0990", // ダイヤルQ2 0990-xxx-xxx
}

var AreaCodes = [][]string{
	AreaCode5,
	AreaCode4,
	AreaCode3,
	AreaCode2,
}

var (
	shortError = errors.New("telephone number is too short.")
	matchError = errors.New("telephone number does not match any area code.")
)

// JpTelSplit splits japaneses telephone number to a slice consist of AreaCode, CityCode, SubscriberNumber.
func JpTelSplit(tel string) ([]string, error) {
	// 固定電話
	for _, areaCode := range AreaCodes {
		for _, code := range areaCode {
			if strings.HasPrefix(tel, code) {
				if len(tel) < totalCodeLen {
					return []string{"", "", ""}, shortError
				}
				codeLen := len(code)
				return []string{
					tel[:codeLen],             // AreaCode
					tel[codeLen:totalCodeLen], // CityCode
					tel[totalCodeLen:],        // SubscriberCode
				}, nil
			}
		}
	}

	// フリーダイヤル
	for _, code := range FreeDialCode {
		if strings.HasPrefix(tel, code) {
			if len(tel) < freeDialPrefixLen + freeDialCodeLen {
				return []string{"", "", ""}, shortError
			}
			return []string{
				tel[:freeDialPrefixLen],
				tel[freeDialPrefixLen : freeDialPrefixLen+freeDialCodeLen],
				tel[freeDialPrefixLen+freeDialCodeLen:],
			}, nil
		}
	}

	// 携帯番号
	for _, code := range MobileCode {
		if strings.HasPrefix(tel, code) {
			if len(tel) < mobileCodePrefixLen + mobileCodeLen {
				return []string{"", "", ""}, shortError
			}
			return []string{
				tel[:mobileCodePrefixLen],
				tel[mobileCodePrefixLen : mobileCodePrefixLen+mobileCodeLen],
				tel[mobileCodePrefixLen+mobileCodeLen:],
			}, nil
		}
	}

	// その他番号1
	for _, code := range OtherCode {
		if strings.HasPrefix(tel, code) {
			if len(tel) < OtherCodePrefixLen + OtherCodeLen {
				return []string{"", "", ""}, shortError
			}
			return []string{
				tel[:OtherCodePrefixLen],
				tel[OtherCodePrefixLen : OtherCodePrefixLen+OtherCodeLen],
				tel[OtherCodePrefixLen+OtherCodeLen:],
			}, nil
		}
	}

	return []string{"", "", ""}, matchError
}
