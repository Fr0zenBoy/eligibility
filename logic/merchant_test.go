package logic
import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestRegexMatchMenchantName(t *testing.T) {
	moes := "moes"
	moesUpperCase := "Moes"
	theCrab := "The Crab Shed"  
	theCrabLowerCase := "the crab shed"
	centralPerk := "Central Perk"
	centralPerkRandomCase := "cEnTraL PERk" 
	assert.Equal(t, RegexMatchMenchantName(moes, moesUpperCase), true)
	assert.Equal(t, RegexMatchMenchantName(theCrab, theCrabLowerCase), true)
	assert.Equal(t, RegexMatchMenchantName(centralPerk, centralPerkRandomCase), true)
}
