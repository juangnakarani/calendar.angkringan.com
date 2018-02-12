package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		log.Println("test print origin->", origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

}

func main() {
	log.Println("starting calendar...")

	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../../ui/dist/")))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
