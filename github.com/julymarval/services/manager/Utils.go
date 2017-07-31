package manager

import "net/http"

// Field validator
func (u *User) Validate(req *http.Request) bool {

	return true
}
