package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

// Config is a config :)
type Config struct {
	PgURL                string `env:"PG_URL"`
	PgProto              string `env:"PG_PROTO"`
	PgAddr               string `env:"PG_ADDR"`
	PgDb                 string `env:"PG_DB"`
	PgUser               string `env:"PG_USER"`
	PgPassword           string `env:"PG_PASSWORD"`
	CallRoomEndpoint     string `env:"CALL_ROOM_ENDPOINT"`
	SchedulerLocation    string `env:"SCHEDULER_LOCATION"`
	SchedulerTimeZone    string `env:"SCHEDULER_TIMEZONE"`
	SchedulerMaxRetryCount  int `env:"SCHEDULER_MAX_RETRY_COUNT"`
	SchedulerRetryPeriod string `env:"SCHEDULER_RETRY_PERIOD"`
	CallEndpoint         string `env:"CALL_ENDPOINT"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment
func Get() *Config {
	once.Do(func() {
		envType := os.Getenv("ENV")
		if envType == "" {
			envType = "dev"
		}
		if err := configor.New(&configor.Config{Environment: envType}).Load(&config, "config.json"); err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
