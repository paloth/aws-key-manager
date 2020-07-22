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
	home    string
	config  *configparser.ConfigParser
	err     error
	reToken = regexp.MustCompile(`[0-9]{6}`)
)

func CheckToken(token string) error {
	if reToken.MatchString(token) {
		return nil
	}
	return fmt.Errorf("The token %s must be composed by six digits", token)
}

func getHomeValue() error {
	home, err = os.UserHomeDir()
	if err != nil {
		return err
	}
	return nil
}

func getConfig() error {
	err := getHomeValue()
	if err != nil {
		return err
	}

	config, err = configparser.NewConfigParserFromFile(home + credentialFile)
	if err != nil {
		return err
	}
	return nil
}

func CheckProfile(profile string) error {
	err = getConfig()
	if err != nil {
		return err
	}

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
	if exists := !config.HasSection(profile + "-tmp"); exists {
		fmt.Println("Profile " + profile + "-tmp does not exists.")
		config.AddSection(profile + "-tmp")
		fmt.Println("Profile " + profile + "-tmp has been created.")
	}
	//Set new profile values in ConfigParser
	config.Set(profile+"-tmp", "aws_access_key_id", *session.Credentials.AccessKeyId)
	config.Set(profile+"-tmp", "aws_secret_access_key", *session.Credentials.SecretAccessKey)
	config.Set(profile+"-tmp", "aws_session_token", *session.Credentials.SessionToken)
	config.Set(profile+"-tmp", "aws_default_region", "eu-west-1")

	//Write profile in file
	err := config.SaveWithDelimiter(home+credentialFile, "=")
	if err != nil {
		return err
	}

	fmt.Println("The profile " + profile + "-tmp has been set up and will expire on " + session.Credentials.Expiration.Format("Mon Jan 2") + " at " + session.Credentials.Expiration.Format("15:04:05"))

	return nil
}
