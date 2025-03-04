package servercfg

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
	DbConfig = DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "backend",
	}
}
