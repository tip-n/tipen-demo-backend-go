package pkg

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (hashedPwd string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	hashedPwd = string(hashed)
	return
}

func CompareHash(hashed string, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(pwd),
	); err != nil {
		return false
	}
	return true
}
