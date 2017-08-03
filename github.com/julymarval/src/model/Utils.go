package model

import (
	"log"

	"github.com/badoux/checkmail"
)

/************************* FUNCS ***************************/

// ValidateEmail = func thats validates a email format, host and if it is an active email
func ValidateEmail(email string) bool {

	var err error

	err = checkmail.ValidateFormat(email)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
