package profile

import (
	"fmt"
	"os"
	"strings"

	"github.com/bigkevmcd/go-configparser"
)

const (
	credentialFile string = "/.aws/credentials"
)

var (
	home        string
	profileList []string
)

func getHomeValue(h *string) {

	var err error
	*h, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}

}

func getConfig() configparser.ConfigParser {

	getHomeValue(&home)

	config, err := configparser.NewConfigParserFromFile(home + credentialFile)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}
	return *config

}

func CheckProfile(profile string) error {

	config := getConfig()
	if exists := config.HasSection(profile); exists {
		flag, err := config.HasOption(profile, "aws_access_key_id")
		if err != nil {
			fmt.Println(err)
		}
		if flag == true && !strings.HasSuffix(profile, "-tmp") {
			return nil
		}
		return fmt.Errorf("You have entered an invalid profile")
	}
	return fmt.Errorf("The profile is not present in your configuration file")
}
