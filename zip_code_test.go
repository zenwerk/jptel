package jptel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPZipCode_Format(t *testing.T) {
	assert := assert.New(t)
	zipCode := JPZipCode{
		PostCode: "123",
		AreaCode: "4567",
	}
	assert.Equal("123-4567", zipCode.Format())
}

func TestValidateZipCode(t *testing.T) {
	assert := assert.New(t)

	// valid zip code
	assert.Nil(ValidateZipCode("1000000"))
	assert.Nil(ValidateZipCode("100-0000"))
	assert.Nil(ValidateZipCode("１００００００"))
	assert.Nil(ValidateZipCode("１００ー００００"))
	assert.Nil(ValidateZipCode("１00000０"))
	assert.Nil(ValidateZipCode("１００-００００"))

	assert.Error(ValidateZipCode("100000"))
	assert.Error(ValidateZipCode("10-00000"))
	assert.Error(ValidateZipCode("10000000"))
	assert.Error(ValidateZipCode("１０００００"))
	assert.Error(ValidateZipCode("１０００ー０００"))
	assert.Error(ValidateZipCode("100-000a"))
	assert.Error(ValidateZipCode("100-000あ"))
}

func TestSplitZipCode(t *testing.T) {
	assert := assert.New(t)

	// valid zipcode
	zipCode, err := SplitZipCode("1000000")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	zipCode, err = SplitZipCode("100-0000")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	zipCode, err = SplitZipCode("１００００００")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	zipCode, err = SplitZipCode("１００ー００００")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	zipCode, err = SplitZipCode("10０0０0０")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	zipCode, err = SplitZipCode("100ー0000")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	zipCode, err = SplitZipCode("1０0ー0-000")
	assert.Nil(err)
	assert.Equal("100", zipCode.PostCode)
	assert.Equal("0000", zipCode.AreaCode)

	// invalid zipcode
	zipCode, err = SplitZipCode("100000")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)

	zipCode, err = SplitZipCode("10000000")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)

	zipCode, err = SplitZipCode("１０００００")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)

	zipCode, err = SplitZipCode("１０００００００")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)

	zipCode, err = SplitZipCode("１0００-００-0０")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)

	zipCode, err = SplitZipCode("100-000a")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)

	zipCode, err = SplitZipCode("１００ー０００あ")
	assert.Error(err)
	assert.Equal("", zipCode.PostCode)
	assert.Equal("", zipCode.AreaCode)
}

func TestNormalizeZipCode(t *testing.T) {
	assert := assert.New(t)

	// valid zipcode
	zip, err := NormalizeZipCode("1000000")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	zip, err = NormalizeZipCode("100-0000")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	zip, err = NormalizeZipCode("１００００００")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	zip, err = NormalizeZipCode("１００ー００００")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	zip, err = NormalizeZipCode("10００0００")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	zip, err = NormalizeZipCode("10０-０0００")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	zip, err = NormalizeZipCode("1-0-0-0--0-0-0")
	assert.Nil(err)
	assert.Equal("100-0000", zip)

	// invalid zipcode
	zip, err = NormalizeZipCode("100000")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("100-000")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("10000000")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("100-00000")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("100000a")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("１０００００")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("１００ー０００")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("１００ー０００００")
	assert.Error(err)
	assert.Equal("", zip)

	zip, err = NormalizeZipCode("１００ー００００あ")
	assert.Error(err)
	assert.Equal("", zip)
}
