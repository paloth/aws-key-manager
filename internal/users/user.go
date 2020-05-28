package users

import (
	"fmt"
	"regexp"
)

type Users struct {
	name    string
	profile string
	token   string
}

var (
	reToken = regexp.MustCompile(`[0-9]{6}`)
)

func (u *Users) Init(name string, profile string, token string) {
	u.name = name
	u.profile = profile
	u.token = token
}

func (u *Users) CheckToken() error {
	if reToken.MatchString(u.token) {
		return nil
	}
	return fmt.Errorf("The token %s must be composed by six digits", u.token)
}
