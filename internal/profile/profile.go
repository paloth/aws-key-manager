package profile

import (
	"fmt"
	"log"
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
	home   string
	config *configparser.ConfigParser
	err    error
)

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

func getCurrentAccessKey(profile *string) (string, error) {
	var oldAccessKey string

	err := CheckProfile(profile)
	if err != nil {
		return "", err
	}

	oldAccessKey, err = config.Get(*profile, "aws_access_key_id")
	if err != nil {
		return "", err
	}

	return oldAccessKey, nil
}

func CheckToken(token *string) error {
	var reToken = regexp.MustCompile(`[0-9]{6}`)

	if reToken.MatchString(*token) {
		return nil
	}
	return fmt.Errorf("The token %s must be composed by six digits", *token)
}

func CheckProfile(profile *string) error {
	err = getConfig()
	if err != nil {
		return err
	}

	if exists := config.HasSection(*profile); exists {
		flag, err := config.HasOption(*profile, "aws_access_key_id")
		if err != nil {
			fmt.Println(err)
		}
		if flag == true && !strings.HasSuffix(*profile, "-tmp") {
			return nil
		}
		return fmt.Errorf("You have entered an invalid profile")
	}
	return fmt.Errorf("The profile %s is not present in your configuration file", profile)
}

func WriteConfigFile(profile *string, session *sts.GetSessionTokenOutput) error {
	if exists := !config.HasSection(*profile + "-tmp"); exists {
		fmt.Println("Profile " + *profile + "-tmp does not exists.")
		config.AddSection(profile + "-tmp")
		fmt.Println("Profile " + *profile + "-tmp has been created.")
	}
	//Set new profile values in ConfigParser
	config.Set(*profile+"-tmp", "aws_access_key_id", *session.Credentials.AccessKeyId)
	config.Set(*profile+"-tmp", "aws_secret_access_key", *session.Credentials.SecretAccessKey)
	config.Set(*profile+"-tmp", "aws_session_token", *session.Credentials.SessionToken)
	config.Set(*profile+"-tmp", "aws_default_region", "eu-west-1")

	//Write profile in file
	err := config.SaveWithDelimiter(home+credentialFile, "=")
	if err != nil {
		return err
	}

	log.Println("The profile " + *profile + "-tmp has been set up and will expire on " + session.Credentials.Expiration.Format("Mon Jan 2") + " at " + session.Credentials.Expiration.Format("15:04:05"))

	return nil
}

func RenewLocalProfile(profile *string, session *sts.GetSessionTokenOutput) error {
	if exists := !config.HasSection(*profile); exists {
		log.Fatalln("Profile " + *profile + " does not exists.")
	}
	//Set new profile values in ConfigParser
	config.Set(*profile, "aws_access_key_id", *session.Credentials.AccessKeyId)
	config.Set(*profile, "aws_secret_access_key", *session.Credentials.SecretAccessKey)

	//Write profile in file
	err := config.SaveWithDelimiter(home+credentialFile, "=")
	if err != nil {
		return err
	}

	log.Println("The profile's access key " + *profile + " has been updated")

	return nil
}
