package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"

	"github.com/hmnayak/credit/controller"
	"github.com/hmnayak/credit/db"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/rest"
	"github.com/hmnayak/credit/ui"
)

// AppConfig is a container of api configuration data
type AppConfig struct {
	Port          string `yaml:"port"`
	PGConn        string `yaml:"pg_conn"`
	StaticDir     string `yaml:"static_dir"`
	AuthSecret    string `yaml:"authsecret"`
	FBServiceFile string `yaml:"service_file_location"`
}

func main() {

	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln("Error reading configuration file:", err)
	}

	var config AppConfig
	err = yaml.Unmarshal([]byte(configFile), &config)
	if err != nil {
		log.Fatalln("Error parsing configuration data:", err)
	}

	opt := option.WithCredentialsFile(config.FBServiceFile)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := fbApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	c := controller.Controller{}
	c.Init(config.PGConn, config.AuthSecret, authClient)

	db, err := db.InitDb(config.PGConn)
	if err != nil {
		log.Fatalln("Error InitDb:", err)
	}

	mdl := model.New(db)
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)
	api.Use(rest.AuthMiddleware(authClient, mdl))

	api.Handle("/customers", rest.UpsertCustomer(mdl)).Methods("PUT")
	api.Handle("/ping", pingHandler(c))

	if config.StaticDir != "" {
		r.PathPrefix("/").Handler(spaHandler(config.StaticDir))
	}

	log.Println("Starting server on port " + config.Port)
	err = http.ListenAndServe(":"+config.Port, r)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func spaHandler(staticDir string) http.Handler {
	fileServer := http.FileServer(http.Dir(staticDir))
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !strings.Contains(req.URL.Path, ".") {
			req.URL.Path = "/"
		}
		fileServer.ServeHTTP(w, req)
	})
}

func pingHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		origin := req.Header.Get("Origin")
		var response ui.Response
		response = ui.CreateResponse(http.StatusOK, "OK", nil)
		ui.Respond(res, response, origin)
	})
}
