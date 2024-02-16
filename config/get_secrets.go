package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func getSecrets(secretName string) (secrets string, err error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	client := secretsmanager.NewFromConfig(awsCfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	if os.Getenv("AWS_RUN") == "local" {
		result, err := client.GetSecretValue(context.TODO(), input, func(o *secretsmanager.Options) {
			o.Credentials = credentials.NewStaticCredentialsProvider("accessKeyId", "accessKeySecret", "")
			o.BaseEndpoint = aws.String("http://localhost:4566")
		})
		if err != nil {
			return "", err
		}
		return *result.SecretString, nil
	} else {
		result, err := client.GetSecretValue(context.TODO(), input)
		if err != nil {
			return "", err
		}
		return *result.SecretString, nil
	}
}
