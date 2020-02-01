package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
	assert := assert.New(t)

	assert.True(ValidateURL("http://test.com"))
	assert.True(ValidateURL("https://test.com"))
	assert.False(ValidateURL("htt:/test.com"))
}

func TestValidateUUID(t *testing.T) {
	assert := assert.New(t)

	assert.True(ValidateUUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"))
	assert.True(ValidateUUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"))
	assert.False(ValidateUUID("nb0bd6f5-c3f5-44b2-b677-acd23cdde73c"))
}
