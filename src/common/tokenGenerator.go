package common

import "crypto/rand"

const TOKEN_LENGTH int = 20

func randomCharacter(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return string(b)
}

func RandomCharacterWithPrefix(prefix string) string {
	return prefix + randomCharacter(TOKEN_LENGTH-len([]rune(prefix)))
}
