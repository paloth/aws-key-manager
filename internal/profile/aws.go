package profile

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sts"
)

func AwsSession(profile string) *session.Session {
	session, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr)
		}
	}

	return session
}

func GetAwsSession(profile string, user string, token string) sts.GetSessionTokenOutput {
	svcSts := sts.New(AwsSession(profile))

	identity, err := svcSts.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr)
		}
	}

	session, err := svcSts.GetSessionToken(&sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(43200),
		SerialNumber:    aws.String("arn:aws:iam::" + *identity.Account + ":mfa/" + user),
		TokenCode:       aws.String(token),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr)
		}
	}
	return *session
}

func CreateNewAccess(session *iam.IAM, profile string, user *string) (iam.CreateAccessKeyOutput, error) {
	newAccessKey, err := session.CreateAccessKey(&iam.CreateAccessKeyInput{UserName: user})
	if err != nil {
		return *newAccessKey, err
	}

	return *newAccessKey, nil
}

func DeleteOldAccessKey(session *iam.IAM, oldAk *string, user *string) error {
	_, err := session.DeleteAccessKey(&iam.DeleteAccessKeyInput{
		AccessKeyId: oldAk,
		UserName:    user,
	})
	if err != nil {
		return err
	}

	return nil
}

func CheckAkNumber(session *iam.IAM, user *string) (bool, error) {
	accessKeys, err := session.ListAccessKeys(&iam.ListAccessKeysInput{
		UserName: user,
	})
	if err != nil {
		return false, err
	}
	if len(accessKeys.AccessKeyMetadata) >= 2 {
		fmt.Println("Too much accessKeys for the user %s", user)
		selectKeyToDelete(&accessKeys.AccessKeyMetadata)
	}
	return true, nil
}

func selectKeyToDelete(list []*iam.AccessKeyMetadata) (bool, error){
	switch runtime.GOOS {
	case "windows":
		for index, value := range list {
			fmt.Println("%d - %s", index, value)
		}

	default:
		prompt := promptui.Select{
			Label: "Select AccessKey to delete",
			Items: list,
		}
	
		_, result, err := prompt.Run()
	
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	
		fmt.Printf("You choose %q\n", result)
	return true, nil
}
