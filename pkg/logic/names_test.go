package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpString(t *testing.T) {
	assert.Equal(t, true, RegexpString("CoMeRcIaL", "ComerCIAL"))
	assert.Equal(t, true, RegexpString("ConvencionaL", "convenCional"))
	assert.Equal(t, true, RegexpString("Industrial-", "Industrial"))
}
