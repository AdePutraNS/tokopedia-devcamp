package config

import (
	"log"
	"sync"

	"github.com/tokopedia/tdk/go/config"
)

type Config struct {
	Appname string
}

var cfg Config
var once = sync.Once{}

func GetConfig() Config {
	once.Do(func() {
		err := config.Read(&cfg, "config/generated_apps.{TKPENV}.yaml", "/etc/generated_apps/generated_apps.{TKPENV}.yaml")
		if err != nil {
			log.Fatal(err)
		}
	})
	return cfg
}
