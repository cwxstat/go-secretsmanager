package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type SecretsManagerCreateSecretAPI interface {
	CreateSecret(ctx context.Context,
		params *secretsmanager.CreateSecretInput,
		optFns ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error)
}

type SecretsManagerDeleteSecretAPI interface {
	DeleteSecret(ctx context.Context,
		params *secretsmanager.DeleteSecretInput,
		optFns ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error)
}

type SecretsManagerGetSecretAPI interface {
	GetSecretValue(ctx context.Context,
		params *secretsmanager.GetSecretValueInput,
		optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

type SecretsManagerListSecretAPI interface {
	ListSecrets(ctx context.Context,
		params *secretsmanager.ListSecretsInput,
		optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)
}

type SecretsManagerUpdateSecretAPI interface {
	UpdateSecret(ctx context.Context,
		params *secretsmanager.UpdateSecretInput,
		optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error)
}
