package jptel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	assert := assert.New(t)

	result, err := Normalize("0123456789")
	assert.Equal("0123-45-6789", result)
	assert.Nil(err)

	result, err = Normalize("0222522222")
	assert.Equal("022-252-2222", result)
	assert.Nil(err)

	result, err = Normalize("０１２３４５６７８９")
	assert.Equal("0123-45-6789", result)
	assert.Nil(err)

	result, err = Normalize("０１２ー３４５６７８-９")
	assert.Equal("0123-45-6789", result)
	assert.Nil(err)

	result, err = Normalize("-0９０-１２３ー4５６-７８-")
	assert.Equal("090-1234-5678", result)
	assert.Nil(err)

	result, err = Normalize("０９０あ−７a９３３ー−４０４３1234567890-")
	assert.Equal("", result)
	assert.Error(err)

	result, err = Normalize("---")
	assert.Equal("", result)
	assert.Error(err)
}
