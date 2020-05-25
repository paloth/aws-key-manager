/*
Copyright © 2020 Paloth

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userName string
var token string
var profile string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a temporary acces key with mfa",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&userName, "username", "u", "", "The AWS user name")
	generateCmd.Flags().StringVarP(&profile, "profile", "p", "", "The AWS user profile")
	generateCmd.Flags().StringVarP(&token, "token", "t", "", "The user token")

	generateCmd.MarkFlagRequired("username")
	generateCmd.MarkFlagRequired("token")
	generateCmd.MarkFlagRequired("profile")
}

const (
	credentialFile string = "/.aws/credentials"
)

var (
	home        string
	profileList []string
)

func run() {
	profile.getConfig()
	getConfig()
}

func getHomeValue(h *string) {
	var err error
	*h, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
}

func getConfig() {
	getHomeValue(&home)
	viper.SetConfigFile("awsProfile")
	viper.SetConfigType("toml")
	viper.AddConfigPath(home + credentialFile)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	viper.Get("")
}
