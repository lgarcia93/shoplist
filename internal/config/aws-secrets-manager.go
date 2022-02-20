package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretData struct {
	MySQLUser               string `json:"username"`
	MySQLPass               string `json:"password"`
	MySQLPort               int    `json:"port"`
	MySQLAddress            string `json:"host"`
	MySQLDatabaseName       string `json:"dbname"`
	MySQLInstanceIdentifier string `json:"dbInstanceIdentifier"`
}

func GetSecret() SecretData {
	secretName := "dev/shopList/mysql"
	region := "us-east-1"
	versionStage := "AWSCURRENT"

	sess, err := session.NewSession()
	svc := secretsmanager.New(
		sess,
		aws.NewConfig().WithRegion(region),
	)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		panic(err.Error())
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	var secretData SecretData
	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		panic(err.Error())
	}

	return secretData

}
