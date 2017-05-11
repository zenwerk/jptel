package jptel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assert := assert.New(t)

	// valid telephone number
	assert.Nil(Validate("0123456789"))
	assert.Nil(Validate("0222522222"))
	assert.Nil(Validate("022-252-2222"))
	assert.Nil(Validate("０１２３４５６７８９"))
	assert.Nil(Validate("０１２３-４５ー６７８９"))
	assert.Nil(Validate("０12３-４５ー６78９"))
	assert.Nil(Validate("０１２３−４５−6789"))

	// invalid telephone number
	assert.Error(Validate("022252-2222"))
	assert.Error(Validate("０１２ー３４５６７８-９"))
	assert.Error(Validate("０１２３４５ー６７８９"))
	assert.Error(Validate("-0９０-１２３ー4５６-７８-"))
	assert.Error(Validate("０９０あ−７a９３３ー−４０４３1234567890-"))
	assert.Error(Validate("--"))
}
