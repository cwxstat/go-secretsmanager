package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(ctx context.Context, api SecretsManagerGetSecretAPI, secretId string) (string, error) {

	result, err := api.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	})

	if err != nil {
		return "", err
	}

	return *result.SecretString, err
}
