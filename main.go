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

	"github.com/hmnayak/credit/config"
	"github.com/hmnayak/credit/controller"
	"github.com/hmnayak/credit/db"
	"github.com/hmnayak/credit/rest"
	"github.com/hmnayak/credit/ui"
)

func main() {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln("Error reading configuration file:", err)
	}

	err = yaml.Unmarshal([]byte(configFile), &config.ApiConfig)
	if err != nil {
		log.Fatalln("Error parsing configuration data:", err)
	}
	if config.ApiConfig.CustomersPageSize == 0 {
		config.ApiConfig.CustomersPageSize = config.DefaultCustomersPageSize
	}

	opt := option.WithCredentialsFile(config.ApiConfig.FBServiceFile)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := fbApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	c := controller.Controller{}
	c.Init(config.ApiConfig.PGConn, config.ApiConfig.AuthSecret, authClient)

	db, err := db.InitDb(config.ApiConfig.PGConn)
	if err != nil {
		log.Fatalln("Error InitDb:", err)
	}

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)
	api.Use(rest.AuthMiddleware(authClient, db))

	api.Handle("/customers", rest.UpsertCustomer(db)).Methods("PUT")
	api.Handle("/customers", rest.ListCustomers(db)).Methods("GET")
	api.Handle("/customers/{customerid}", rest.GetCustomer(db)).Methods("GET")
	api.Handle("/ping", pingHandler(c))

	if config.ApiConfig.StaticDir != "" {
		r.PathPrefix("/").Handler(spaHandler(config.ApiConfig.StaticDir))
	}

	log.Println("Starting server on port " + config.ApiConfig.Port)
	err = http.ListenAndServe(":"+config.ApiConfig.Port, r)
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
		var response ui.Response
		response = ui.CreateResponse(http.StatusOK, nil)
		ui.Respond(res, response)
	})
}
