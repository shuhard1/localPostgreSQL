package config

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() *StorageConfig {
	const (
		defaultHost     = "localhost"
		defaultPort     = "5432"
		defaultDatabase = "postgres"
		defaultUsername = "postgres"
		defaultPassword = "1234"
	)

	sc := &StorageConfig{
		Host:     defaultHost,
		Port:     defaultPort,
		Database: defaultDatabase,
		Username: defaultUsername,
		Password: defaultPassword,
	}
	return sc
}
