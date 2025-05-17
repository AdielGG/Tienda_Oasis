package config

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

var (
	DbConfig DatabaseConfig
)

func InitDatabaseConfig() {
	// func InitDatabaseConfig(namedb, pasword string) {
	DbConfig = DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "adiel",
		Password: "Adiel123",
		Database: "oasis",
	}
}
