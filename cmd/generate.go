/*
Copyright © 2020 Paloth

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/paloth/aws-key-manager/internal/profile"
	"github.com/paloth/aws-key-manager/internal/users"
	"github.com/spf13/cobra"
)

var (
	userName    string
	userToken   string
	userProfile string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a temporary token",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&userName, "username", "u", "", "AWS user name")
	generateCmd.Flags().StringVarP(&userProfile, "profile", "p", "", "AWS user profile (Must be in the aws credentials file)")
	generateCmd.Flags().StringVarP(&userToken, "token", "t", "", "User's token (Composed by 6 digits)")

	generateCmd.MarkFlagRequired("username")
	generateCmd.MarkFlagRequired("token")
	generateCmd.MarkFlagRequired("profile")
}

func run() {
	user := users.Users{
		Name:    userName,
		Token:   userToken,
		Profile: userProfile,
	}

	err := user.CheckToken()
	if err != nil {
		fmt.Printf("Input error: %s", err)
	}

	err = profile.CheckProfile(user.Profile)
	if err != nil {
		fmt.Printf("Input error: %s", err)
	}
}
