/*
Copyright © 2021 Paloth

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
	"github.com/spf13/cobra"
)

// rotateCmd represents the rotate command
var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "A brief description of your command",
	Long:  ``,
	RunE:  execRotate,
}

func init() {
	rootCmd.AddCommand(rotateCmd)

	rotateCmd.Flags().StringP("username", "u", "", "AWS user name")
	rotateCmd.Flags().StringP("profile", "p", "default", "AWS user profile (Must be in the aws credentials file)")

	rotateCmd.MarkFlagRequired("username")
	rotateCmd.MarkFlagRequired("token")
}

func execRotate(cmd *cobra.Command, args []string) error {

	//Check user name entry
	userName, err := cmd.Flags().GetString("username")
	if err != nil {
		return err
	}
	if userName == "" {
		return fmt.Errorf("User name cannot be empty! Please provide a user name")
	}

	//Check profile entry
	userProfile, err := cmd.Flags().GetString("profile")
	if err != nil {
		return err
	}
	err = profile.CheckProfile(userProfile)
	if err != nil {
		return err
	}

	session := profile.GetAwsSession(userProfile, userName, userToken)

	err = profile.WriteConfigFile(userProfile, &session)
	if err != nil {
		return err
	}

	return nil
}
