package jptel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPTelephoneNumber_Format(t *testing.T) {
	assert := assert.New(t)

	num := JPTelephoneNumber{
		AreaCode: "090",
		CityCode: "1234",
		SubscriberCode: "5678",
	}

	assert.Equal("090-1234-5678", num.Format())
}
