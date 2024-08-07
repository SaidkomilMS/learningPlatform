package config

type Config struct {
	DBSource     string
	JWTLifeTime  uint
	JWTSecretKey string
	JWTIssuer    string
}

func LoadConfig() *Config {
	return &Config{
		DBSource:     "host=localhost user=learning_platform password=test_password dbname=learning_platform_db port=5432 sslmode=disable TimeZone=Asia/Tashkent",
		JWTLifeTime:  1 * 60 * 60 * 1000000000,
		JWTSecretKey: "your_secret_key_here",
		JWTIssuer:    "learning_platform",
	}
}
