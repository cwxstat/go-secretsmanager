package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// This creates a secret and returns out its ARN.
func CreateSecret(ctx context.Context,
	api SecretsManagerCreateSecretAPI,
	name string,
	value string,
	description string) (string, error) {

	result, err := api.CreateSecret(ctx, &secretsmanager.CreateSecretInput{
		Name: aws.String(name),
		// descriptions are optional
		Description: aws.String(description),
		// You must provide either SecretString or SecretBytes.
		// Both is considered invalid.
		SecretString: aws.String(value),
	})

	if err != nil {
		return "", err
	}

	return *result.ARN, err

}
