package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true"`
	Port    int    `yaml:"port" env-required:"true"`
}

type SMTPMail struct {
	Host string `env:"SMTP_HOST" env-required:"true"`
	Port string `env:"SMTP_PORT" env-required:"true"`
	User string `env:"SMTP_USER" env-required:"true"`
	Pass string `env:"SMTP_PASS" env-required:"true"`
	From string `env:"SMTP_FROM" env-required:"true"`
}

type Config struct {
	Env           string `yaml:"env" env-required:"true" env-default:"production`
	DatabaseUri   string `env:"DATABASE_URI" env-required:"true"`
	JwtPrivateKey string `env:"JWT_PRIVATE_KEY env-required:"true"`
	HTTPServer    `yaml:"http_server"`
	SMTPMail
}

func ConfigMustLoad() *Config {
	var configPath string
	godotenv.Load(".env")
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	return &cfg
}
