package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Kullanıcı adı kullanımda!")
	}

	if strings.Contains(err, "email") {
		return errors.New("Mail adresi kullanımda!")
	}

	if strings.Contains(err, "title") {
		return errors.New("Başlık kullanımda!")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Yanlış şifre!")
	}
	return errors.New("Bir şeyler yanlış!")
}
