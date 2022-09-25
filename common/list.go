package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func ListSecrets(ctx context.Context, api SecretsManagerListSecretAPI) ([]string, error) {

	secrets := make([]string, 0)
	result, err := api.ListSecrets(ctx, &secretsmanager.ListSecretsInput{})
	if err != nil {
		return secrets, err
	}
	for _, secret := range result.SecretList {
		secrets = append(secrets, *secret.ARN)
	}

	return secrets, err
}
