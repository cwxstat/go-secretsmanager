package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"testing"
)

type mockUpdateSecret func(ctx context.Context, params *secretsmanager.UpdateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error)

func (m mockUpdateSecret) UpdateSecret(ctx context.Context, params *secretsmanager.UpdateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error) {
	return m(ctx, params, optFns...)
}

func TestUpdateSecret(t *testing.T) {
	cases := []struct {
		client       func(t *testing.T) SecretsManagerUpdateSecretAPI
		name         string
		secretId     string
		secretString string
		expect       []byte
	}{
		{
			client: func(t *testing.T) SecretsManagerUpdateSecretAPI {
				return mockUpdateSecret(func(ctx context.Context, params *secretsmanager.UpdateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error) {
					t.Helper()
					if params.SecretId == nil || params.SecretString == nil {
						t.Errorf("expected secretID and SecretString to be set")
					}
					if e, a := "example", *params.SecretId; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &secretsmanager.UpdateSecretOutput{
						ARN: aws.String("arn:aws:secretsmanager:us-west-2:123456789012:secret:example-123456"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			err := UpdateSecret(ctx, c.client(t), "example", "value")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}
