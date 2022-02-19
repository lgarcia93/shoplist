package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretData struct {
	MySQLUser         string `json:"mysql_user"`
	MySQLPass         string `json:"mysql_pass"`
	MySQLPort         string `json:"mysql_port"`
	MySQLAddress      string `json:"mysql_address"`
	MySQLDatabaseName string `json:"mysql_db"`
}

var (
	secretName   string = "rds_mysql_shoplist"
	region       string = "us-east-1"
	versionStage string = "AWSCURRENT"
)

func GetSecret() SecretData {
	svc := secretsmanager.New(
		session.New(),
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
