package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Configuration struct {
	address string
	port    int
}

var config = new(Configuration)

func loadConfiguration() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Printf("Fatal error: %s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})

	viper.SetConfigType("yaml")

	config.address = viper.Get("address").(string)
	config.port = viper.Get("port").(int)

	log.Printf("->>%s\n", config.address)
	log.Printf("->>%v\n", config.port)
}

func init() {
	loadConfiguration()
}

func main() {
	log.Println("starting calendar...")
}
