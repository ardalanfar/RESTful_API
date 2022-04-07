package conf

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Charset  string
}

//Set configs
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "admin",
			Password: "123456",
			Dbname:   "service_restful",
			Charset:  "utf8",
		},
	}
}
