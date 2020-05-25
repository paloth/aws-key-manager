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
	"regexp"

	"github.com/paloth/aws-key-manager/internal/profile"
	"github.com/spf13/cobra"
)

var (
	userName    string
	userToken   string
	userProfile string
	reToken     = regexp.MustCompile(`\d{6}`)
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&userName, "username", "u", "", "The AWS user name")
	generateCmd.Flags().StringVarP(&userProfile, "profile", "p", "", "The AWS user profile")
	generateCmd.Flags().StringVarP(&userToken, "token", "t", "", "The user token")

	generateCmd.MarkFlagRequired("username")
	generateCmd.MarkFlagRequired("token")
	generateCmd.MarkFlagRequired("profile")
}

func checkToken(userToken string) error {

	if match := reToken.MatchString(userToken); match {
		return nil
	}
	return fmt.Errorf("The token %s must be composed by six digits")
}

func run() {
	err := checkToken(userToken)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	config := profile.GetConfig()
	profile.CheckProfile(userProfile, config)
	profile.Main()
	// config := profile.GetConfig()
}
