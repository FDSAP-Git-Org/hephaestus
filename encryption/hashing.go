package encryption

import "golang.org/x/crypto/bcrypt"

// Hash string ...
func GenerateHash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(bytes), err
}

// Validate hash string...
func ValidateHash(text, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}

// Hide text in console
func HidePassword(password string) string {
	passLen := len(password)
	newLen := passLen - 3
	tempPass := ""
	for newLen >= 0 {
		tempPass += "x"
		newLen--
	}
	tempPass += password[passLen-3 : passLen]
	return tempPass
}
