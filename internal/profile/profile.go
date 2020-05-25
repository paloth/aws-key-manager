package profile

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
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

func GetConfig() {
	getHomeValue(&home)

	filePath := home + credentialFile
	viper.SetConfigFile("awsProfile")
	viper.SetConfigType("toml")
	viper.AddConfigPath(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	viper.Get("")
}
