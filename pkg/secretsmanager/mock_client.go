package secretsmanager

// MockSecretsManagerClient is a mock implementation of SecretsManagerClient.
type MockSecretsManagerClient struct{}

func NewMockSecretsManagerClient() (*MockSecretsManagerClient, error) {
	return &MockSecretsManagerClient{}, nil
}

// CreateSecretValue calls the mocked CreateSecretValueFunc.
func (m *MockSecretsManagerClient) CreateSecretValue(secretName, secretValue string) error {
	return nil
}

func (m *MockSecretsManagerClient) ReadSecretValue(secretName string) (string, error) {
	return "mock-secret-value", nil
}
