package common

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"testing"
	"time"
)

type mockListSecret func(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error)

func (m mockListSecret) ListSecrets(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
	return m(ctx, params, optFns...)
}

func TestListSecrets(t *testing.T) {
	list := []types.SecretListEntry{}
	for i, v := range []string{"a", "b", "c"} {
		list = append(list, types.SecretListEntry{
			ARN:                    aws.String("arn:aws:secretsmanager:us-east-1:123456789012:secret:MyTestDatabaseSecret-ABC123"),
			CreatedDate:            aws.Time(time.Now()),
			DeletedDate:            nil,
			Description:            aws.String("My test database secret: " + v),
			KmsKeyId:               aws.String(fmt.Sprintf("key-id-%d", i)),
			LastAccessedDate:       aws.Time(time.Now()),
			LastChangedDate:        nil,
			LastRotatedDate:        nil,
			Name:                   aws.String("MyTestDatabaseSecret" + v),
			OwningService:          nil,
			PrimaryRegion:          nil,
			RotationEnabled:        false,
			RotationLambdaARN:      nil,
			RotationRules:          nil,
			SecretVersionsToStages: nil,
			Tags:                   nil,
		})
	}
	cases := []struct {
		client func(t *testing.T) SecretsManagerListSecretAPI
		name   string
		expect []byte
	}{
		{
			client: func(t *testing.T) SecretsManagerListSecretAPI {
				return mockListSecret(func(ctx context.Context, params *secretsmanager.ListSecretsInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.ListSecretsOutput, error) {
					t.Helper()

					return &secretsmanager.ListSecretsOutput{
						SecretList: list,
					}, nil
				})
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.TODO()
			list, err := ListSecrets(ctx, c.client(t))
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if len(list) != 3 {
				t.Errorf("expected 3 secrets, got %d", len(list))
			}
		})
	}

}
