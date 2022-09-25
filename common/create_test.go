package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"testing"
)

type mockCreateSecret func(ctx context.Context, params *secretsmanager.CreateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error)

func (m mockCreateSecret) CreateSecret(ctx context.Context, params *secretsmanager.CreateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error) {
	return m(ctx, params, optFns...)
}

func TestCreateSecret(t *testing.T) {
	cases := []struct {
		client       func(t *testing.T) SecretsManagerCreateSecretAPI
		name         string
		description  string
		secretString string
		expect       []byte
	}{
		{
			client: func(t *testing.T) SecretsManagerCreateSecretAPI {
				return mockCreateSecret(func(ctx context.Context, params *secretsmanager.CreateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error) {
					t.Helper()
					if params.Name == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "example", *params.Name; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &secretsmanager.CreateSecretOutput{
						ARN: aws.String("arn:aws:secretsmanager:us-west-2:123456789012:secret:example-123456"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := CreateSecret(ctx, c.client(t), "example", "value", "description")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
