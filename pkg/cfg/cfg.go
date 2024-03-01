package cfg

type Config struct {
	AppConfig    AppConfig
	PgConfig     PostgresqlConfig
	JwtConfig    JwtConfig
	MailerConfig MailerConfig
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
	MailerConfig struct {
		Host     string `env:"MAILER_HOST, required"`
		Port     int    `env:"MAILER_PORT, required"`
		Username string `env:"MAILER_USERNAME"`
		Password string `env:"MAILER_PASSWORD"`
	}
)

func (jc JwtConfig) GetJwtSecret() []byte {
	return []byte(jc.Secret)
}
