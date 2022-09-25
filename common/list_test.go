package common

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"testing"
)

type mockListSecret func(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)

func (m mockListSecret) ListSecrets(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
	return m(ctx, params, optFns...)
}

func TestListSecrets(t *testing.T) {
	cases := []struct {
		client func(t *testing.T) SecretsManagerListSecretAPI
		name   string
		expect []byte
	}{
		{
			client: func(t *testing.T) SecretsManagerListSecretAPI {
				return mockListSecret(func(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
					t.Helper()

					return &secretsmanager.ListSecretsOutput{}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := ListSecrets(ctx, c.client(t))
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}
