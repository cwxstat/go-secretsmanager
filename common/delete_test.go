package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"testing"
)

type mockDeleteSecret func(ctx context.Context, params *secretsmanager.DeleteSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error)

func (m mockDeleteSecret) DeleteSecret(ctx context.Context, params *secretsmanager.DeleteSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error) {
	return m(ctx, params, optFns...)
}

func TestDeleteSecret(t *testing.T) {
	cases := []struct {
		client   func(t *testing.T) SecretsManagerDeleteSecretAPI
		secretId string
		expect   []byte
	}{
		{
			client: func(t *testing.T) SecretsManagerDeleteSecretAPI {
				return mockDeleteSecret(func(ctx context.Context, params *secretsmanager.DeleteSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.DeleteSecretOutput, error) {
					t.Helper()
					if params.SecretId == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "example", *params.SecretId; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &secretsmanager.DeleteSecretOutput{
						ARN: aws.String("arn:aws:secretsmanager:us-west-2:123456789012:secret:example-123456"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.secretId, func(t *testing.T) {
			ctx := context.TODO()
			err := DeleteSecret(ctx, c.client(t), "example")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}
