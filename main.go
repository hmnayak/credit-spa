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
	configFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalln("Error reading configuration file:", err)
	}

	var appConfig config.AppConfig
	err = yaml.Unmarshal([]byte(configFile), &appConfig)
	if err != nil {
		log.Fatalln("Error parsing configuration data:", err)
	}
	if appConfig.CustomersPageSize == 0 {
		appConfig.CustomersPageSize = config.DefaultCustomersPageSize
	}

	opt := option.WithCredentialsFile(appConfig.FBServiceFile)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := fbApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	c := controller.Controller{}
	c.Init(appConfig.PGConn, appConfig.AuthSecret, authClient)

	db, err := db.InitDb(appConfig.PGConn)
	if err != nil {
		log.Fatalln("Error InitDb:", err)
	}

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.Use(loggingMiddleware)
	api.Use(rest.AuthMiddleware(authClient, db))

	api.Handle("/customers", rest.UpsertCustomer(db)).Methods("PUT")
	api.Handle("/customers", rest.ListCustomers(db, appConfig.CustomersPageSize)).Methods("GET")
	api.Handle("/customers/{customerid}", rest.GetCustomer(db)).Methods("GET")
	api.Handle("/ping", pingHandler(c))

	if appConfig.StaticDir != "" {
		r.PathPrefix("/").Handler(spaHandler(appConfig.StaticDir))
	}

	log.Println("Starting server on port " + appConfig.Port)
	err = http.ListenAndServe(":"+appConfig.Port, r)
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
