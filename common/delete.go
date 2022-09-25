package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func DeleteSecret(ctx context.Context, api SecretsManagerDeleteSecretAPI, secretId string) error {
	_, err := api.DeleteSecret(ctx, &secretsmanager.DeleteSecretInput{
		SecretId: aws.String(secretId),
	})
	return err
}
