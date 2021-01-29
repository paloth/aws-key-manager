package psession

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func GetTmpSession(awsSession *session.Session, user *string, token *string) sts.GetSessionTokenOutput {
	svcSts := sts.New(awsSession)

	identity, err := svcSts.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Println(awsErr)
		}
	}

	tmpSession, err := svcSts.GetSessionToken(&sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(43200),
		SerialNumber:    aws.String("arn:aws:iam::" + *identity.Account + ":mfa/" + *user),
		TokenCode:       aws.String(*token),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			log.Println(awsErr)
		}
	}
	return *tmpSession
}
