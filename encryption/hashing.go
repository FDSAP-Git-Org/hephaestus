package encryption

import "golang.org/x/crypto/bcrypt"

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
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
