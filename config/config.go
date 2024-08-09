package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBSource     string
	JWTLifeTime  uint
	JWTSecretKey string
	JWTIssuer    string
}

func LoadConfig() *Config {
	DBString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"))
	fmt.Println("DBConnection string:", DBString)
	return &Config{
		//DBSource:     "host=" + host + "user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Tashkent",
		DBSource:     DBString,
		JWTLifeTime:  1 * 60 * 60 * 1000000000,
		JWTSecretKey: "your_secret_key_here",
		JWTIssuer:    "learning_platform",
	}
}
