package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	rice "github.com/GeertJohan/go.rice"
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
	log.Println("loading configuration...")
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

// Handler return a http.Handler that supports Vue Router app with history mode
func VueHandler(dist string) http.Handler {
	box := rice.MustFindBox("../../ui/dist/").HTTPBox() //look repeating param but is ok. It must literal
	handler := http.FileServer(box)

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.String()
		// log.Println("check url")
		// static files
		if strings.Contains(url, ".") || url == "/" {
			handler.ServeHTTP(w, req)
			return
		}

		// the all 404 gonna be served as root
		// http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
		http.ServeFile(w, req, path.Join(dist, "/index.html"))
	})
}
func main() {
	log.Println("starting server...")
	log.Printf("address: %s", config.address)
	log.Printf("port: %v", config.port)

	r := mux.NewRouter()

	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("../../ui/dist/")))
	r.PathPrefix("/assets").Handler(http.FileServer(rice.MustFindBox("../../ui/src/").HTTPBox()))
	// r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("../../ui/dist/").HTTPBox()))
	// box := rice.MustFindBox("../../ui/dist/").HTTPBox()
	r.PathPrefix("/").Handler(VueHandler("../../ui/dist/"))
	conf := fmt.Sprintf("%s:%v", config.address, config.port)

	srv := &http.Server{
		Handler: r,
		Addr:    conf,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
