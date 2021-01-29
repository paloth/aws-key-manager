package pkeys

import (
	"fmt"
	"log"
	"runtime"

	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/manifoldco/promptui"
)

func CreateKey(session *iam.IAM, profile string, user *string) (iam.CreateAccessKeyOutput, error) {
	err := tooManyKeys(session, user)

	newAccessKey, err := session.CreateAccessKey(&iam.CreateAccessKeyInput{
		UserName: user,
	})
	if err != nil {
		return *newAccessKey, err
	}

	return *newAccessKey, nil
}

func DeleteKey(session *iam.IAM, keyId *string, user *string) error {
	_, err := session.DeleteAccessKey(&iam.DeleteAccessKeyInput{
		AccessKeyId: keyId,
		UserName:    user,
	})
	if err != nil {
		return err
	}

	return nil
}

func tooManyKeys(session *iam.IAM, user *string) error {
	accessKeys, err := session.ListAccessKeys(&iam.ListAccessKeysInput{
		UserName: user,
	})
	if err != nil {
		return err
	}

	if len(accessKeys.AccessKeyMetadata) >= 2 {
		log.Println("Too much accessKeys for the user %s", user)
		key, err := KeySelection(accessKeys.AccessKeyMetadata)
		if err != nil {
			return err
		}

		err = DeleteKey(session, &key, user)
		if err != nil {
			return err
		}
	}
	return nil
}

func KeySelection(keys []*iam.AccessKeyMetadata) (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "", nil

	default:
		prompt := promptui.Select{
			Label: "Select AccessKey to delete",
			Items: keys,
		}

		_, result, err := prompt.Run()

		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return "", err
		}

		fmt.Printf("You choose %q\n", result)
		return result, nil
	}
}
