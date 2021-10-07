package common

import (
	"crypto/rand"
	"encoding/hex"
)

const TOKEN_LENGTH int = 12

func randomCharacter(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func RandomCharacterWithPrefix(prefix string) string {
	return prefix + randomCharacter(TOKEN_LENGTH-len([]rune(prefix)))
}
