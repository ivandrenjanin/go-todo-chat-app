package cfg

type Config struct {
	AppConfig AppConfig
	PgConfig  PostgresqlConfig
	JwtConfig JwtConfig
}

type (
	PostgresqlConfig struct {
		Host     string `env:"DB_HOST, required"`
		Port     int    `env:"DB_PORT, required"`
		User     string `env:"DB_USER, required"`
		Password string `env:"DB_PASSWORD, required"`
		DbName   string `env:"DB_NAME, required"`
		SslMode  string `env:"DB_SSL_MODE, required"`
	}
	AppConfig struct {
		Host string `env:"APP_HOST, required"`
		Port int    `env:"APP_PORT, required"`
	}
	JwtConfig struct {
		Secret string `env:"JWT_SECRET_STRING, required"`
	}
)

func (jc JwtConfig) GetJwtSecret() []byte {
	return []byte(jc.Secret)
}
