package jptel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	assert := assert.New(t)

	// valid telephone number
	result, err := Split("0312345678")
	assert.Nil(err)
	assert.Equal("03", result.AreaCode)
	assert.Equal("1234", result.CityCode)
	assert.Equal("5678", result.SubscriberCode)

	result, err = Split("03-1234-5678")
	assert.Nil(err)
	assert.Equal("03", result.AreaCode)
	assert.Equal("1234", result.CityCode)
	assert.Equal("5678", result.SubscriberCode)

	result, err = Split("０３１２３４５６７８")
	assert.Nil(err)
	assert.Equal("03", result.AreaCode)
	assert.Equal("1234", result.CityCode)
	assert.Equal("5678", result.SubscriberCode)

	result, err = Split("０３１-２３４５ー６７８")
	assert.Nil(err)
	assert.Equal("03", result.AreaCode)
	assert.Equal("1234", result.CityCode)
	assert.Equal("5678", result.SubscriberCode)

	result, err = Split("0222522222")
	assert.Nil(err)
	assert.Equal("022", result.AreaCode)
	assert.Equal("252", result.CityCode)
	assert.Equal("2222", result.SubscriberCode)

	result, err = Split("0997123456")
	assert.Nil(err)
	assert.Equal("0997", result.AreaCode)
	assert.Equal("12", result.CityCode)
	assert.Equal("3456", result.SubscriberCode)

	result, err = Split("0996912345")
	assert.Nil(err)
	assert.Equal("09969", result.AreaCode)
	assert.Equal("1", result.CityCode)
	assert.Equal("2345", result.SubscriberCode)

	result, err = Split("09012345678")
	assert.Nil(err)
	assert.Equal("090", result.AreaCode)
	assert.Equal("1234", result.CityCode)
	assert.Equal("5678", result.SubscriberCode)

	result, err = Split("0120123456")
	assert.Nil(err)
	assert.Equal("0120", result.AreaCode)
	assert.Equal("123", result.CityCode)
	assert.Equal("456", result.SubscriberCode)

	result, err = Split("08001234567")
	assert.Nil(err)
	assert.Equal("0800", result.AreaCode)
	assert.Equal("123", result.CityCode)
	assert.Equal("4567", result.SubscriberCode)

	// invalid telephone number
	result, err = Split("00000000000")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)

	result, err = Split("００００００００００")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)

	result, err = Split("0312")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)

	result, err = Split("０３１２")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)

	result, err = Split("090")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)

	result, err = Split("hogefugapiyo")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)

	result, err = Split("あいうえお")
	assert.Error(err)
	assert.Equal("", result.AreaCode)
	assert.Equal("", result.CityCode)
	assert.Equal("", result.SubscriberCode)
}
