package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:     "127.0.0.1",
		DBPort:     "13306",
		DBUser:     "root",
		DBPassword: "1234",
		DBName:     "zero",
		JWTSecret:  "SECRET",
	}
}
