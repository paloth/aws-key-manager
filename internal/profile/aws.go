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
