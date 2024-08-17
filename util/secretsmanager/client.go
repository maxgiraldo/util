package secretsmanager

type SecretsManagerClient interface {
	CreateSecretValue(secretName, secretValue string) error
	ReadSecretValue(secretName string) (string, error)
}
