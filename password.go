package misc

import "golang.org/x/crypto/bcrypt"

// PasswordHash 使用 Bcrypt 算法生成密码哈希值
func PasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// PasswordVerify 比较输入的密码与哈希值是否匹配
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
