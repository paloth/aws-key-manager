package profile

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/bigkevmcd/go-configparser"
)

const (
	credentialFile string = "/.aws/credentials"
)

var (
	home        string
	config      *configparser.ConfigParser
	profileList []string
	err         error
	reToken     = regexp.MustCompile(`[0-9]{6}`)
)

func CheckToken(token string) error {
	if reToken.MatchString(token) {
		return nil
	}
	return fmt.Errorf("The token %s must be composed by six digits", token)
}

func getHomeValue(h *string) {
	*h, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(2)
	}
}

func getConfig() *configparser.ConfigParser {
	getHomeValue(&home)

	cfg, err := configparser.NewConfigParserFromFile(home + credentialFile)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}
	return cfg
}

func CheckProfile(profile string) error {
	config = getConfig()

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
	return fmt.Errorf("The profile %s is not present in your configuration file", profile)
}

func WriteConfigFile(profile string, session *sts.GetSessionTokenOutput) error {
	if exists := config.HasSection(profile + "-tmp"); exists {
		config.Set(profile+"-tmp", "aws_access_key_id", *session.Credentials.AccessKeyId)
		config.Set(profile+"-tmp", "aws_secret_access_key", *session.Credentials.SecretAccessKey)
		config.Set(profile+"-tmp", "aws_session_token", *session.Credentials.SessionToken)
		config.Set(profile+"-tmp", "aws_default_region", "eu-west-1")
	} else {
		fmt.Println("Profile " + profile + "-tmp does not exists.")
		config.AddSection(profile + "-tmp")
		fmt.Println("Profile " + profile + "-tmp has been created.")
		config.Set(profile+"-tmp", "aws_access_key_id", *session.Credentials.AccessKeyId)
		config.Set(profile+"-tmp", "aws_secret_access_key", *session.Credentials.SecretAccessKey)
		config.Set(profile+"-tmp", "aws_session_token", *session.Credentials.SessionToken)
		config.Set(profile+"-tmp", "aws_default_region", "eu-west-1")
	}

	err := config.SaveWithDelimiter(home+credentialFile, "=")
	if err != nil {
		return err
	}

	return nil
}
