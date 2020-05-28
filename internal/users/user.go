package users

import (
	"fmt"
	"regexp"
)

type Users struct {
	Name    string
	Profile string
	Token   string
}

var (
	reToken = regexp.MustCompile(`[0-9]{6}`)
)

func (u *Users) Init(name string, profile string, token string) {
	u.Name = name
	u.Profile = profile
	u.Token = token
}

func (u *Users) CheckToken() error {
	if reToken.MatchString(u.Token) {
		return nil
	}
	return fmt.Errorf("The token %s must be composed by six digits", u.Token)
}
