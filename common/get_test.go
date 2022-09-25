package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"testing"
)

type mockGetSecret func(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)

func (m mockGetSecret) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	return m(ctx, params, optFns...)
}

func TestGetSecret(t *testing.T) {
	cases := []struct {
		client   func(t *testing.T) SecretsManagerGetSecretAPI
		name     string
		secretId string
		expect   []byte
	}{
		{
			client: func(t *testing.T) SecretsManagerGetSecretAPI {
				return mockGetSecret(func(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
					t.Helper()
					if params.SecretId == nil {
						t.Errorf("expected name to be set")
					}
					if e, a := "example", *params.SecretId; e != a {
						t.Errorf("expected %v, got %v", e, a)
					}
					return &secretsmanager.GetSecretValueOutput{
						ARN:          aws.String("arn:aws:secretsmanager:us-west-2:123456789012:secret:example-123456"),
						Name:         aws.String("example"),
						SecretString: aws.String("example"),
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			_, err := GetSecret(ctx, c.client(t), "example")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}
