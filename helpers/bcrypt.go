package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(str string) string {
	salt := 8
	byteToHash := []byte(str)

	hash, err := bcrypt.GenerateFromPassword(byteToHash, salt)
	if err != nil {
		panic("failed to hash string")
	}
	return string(hash)
}

func CompareHash(hash, valueToCompare string) bool {
	hashByte, valueByte := []byte(hash), []byte(valueToCompare)

	err := bcrypt.CompareHashAndPassword(hashByte, valueByte)
	return err == nil
}
