package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomCharacterWithPrefix(t *testing.T) {
	// a := randomCharacter(TOKEN_LENGTH)
	b := RandomCharacterWithPrefix("tst_")
	assert := assert.New(t)

	assert.Equal(len([]rune(b)), 20)
}
