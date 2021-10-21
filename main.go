package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gookit/ini"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

var profiles = []string{"default"}

func main() {

	log.Debug("parsing Keycloak profile from environment")
	awsrole, found := os.LookupEnv("AWS_KEYCLOAK_PROFILE")

	if found {
		profiles = append(profiles, awsrole)
	}


	log.Debug("parsing aws credentials from default chain")
	sess := session.Must(session.NewSession())

	creds, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("open AWS credentials file")
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	credentialsFile := path.Join(home, ".aws", "credentials")

	awsCredentialsFile, err := ini.LoadExists(credentialsFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("writing session credentials to file")

	for _, s := range profiles {
		awsCredentialsFile.SetSection(s, map[string]string{
			"aws_access_key_id": creds.AccessKeyID,
			"aws_secret_access_key"	: creds.SecretAccessKey,
			"aws_session_token"	: creds.SessionToken,
		})
	}

	awsCredentialsFile.WriteToFile(credentialsFile)

	log.Info("Write AWS Session to file successful!")
}

