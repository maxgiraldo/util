package secretsmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// AWSSecretsManagerClient is a concrete implementation of SecretsManagerClient.
type AWSSecretsManagerClient struct {
	client *secretsmanager.Client
}

// NewAWSSecretsManagerClient creates a new AWSSecretsManagerClient.
func NewAWSSecretsManagerClient() (*AWSSecretsManagerClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	return &AWSSecretsManagerClient{client: client}, nil
}

// CreateSecretValue uploads a secret to AWS Secrets Manager.
func (c *AWSSecretsManagerClient) CreateSecretValue(secretName, secretValue string) error {
	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(secretName),
		SecretString: aws.String(secretValue),
	}

	_, err := c.client.CreateSecret(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to create secret, %v", err)
	}

	fmt.Printf("Successfully created secret: %s\n", secretName)
	return nil
}

// ReadSecretValue retrieves a secret from AWS Secrets Manager.
func (c *AWSSecretsManagerClient) ReadSecretValue(secretName string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := c.client.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve secret, %v", err)
	}

	if result.SecretString == nil {
		return "", fmt.Errorf("secret string is nil")
	}

	return *result.SecretString, nil
}
