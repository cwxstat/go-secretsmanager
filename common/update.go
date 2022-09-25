package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func UpdateSecret(ctx context.Context, api SecretsManagerUpdateSecretAPI, secretId string, newValue string) error {
	_, err := api.UpdateSecret(ctx, &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(secretId),
		SecretString: aws.String(newValue),
	})
	return err
}
