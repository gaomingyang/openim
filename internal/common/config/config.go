package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(workpath string, environment string) {
	log.Println("config Init,environment:", environment)
	var configFileName string
	var env string
	// 优先使用传入的变量
	if environment != "" {
		env = environment
	} else {
		osEnv := os.Getenv("APP_ENV")
		log.Printf("ENV=%s\n", osEnv)
		env = osEnv
	}

	switch env {
	case "dev":
		configFileName = "app_dev.yaml"
	case "test":
		configFileName = "app_test.yaml"
	case "online":
		configFileName = "app_online.yaml"
	default:
		log.Println("empty env,use default file app.yaml")
		configFileName = "app.yaml"
	}
	configPath := filepath.Join(workpath, "conf", configFileName)
	// log.Println("configPath:", configPath)
	err := initConfig(configPath)
	if err != nil {
		panic(err)
	}

}

func initConfig(filepath string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Printf("config file %s not exist", filepath)
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	return err
}
