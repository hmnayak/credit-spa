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

	. "github.com/hmnayak/credit/config"
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

	err = yaml.Unmarshal([]byte(configFile), &ApiConfig)
	if err != nil {
		log.Fatalln("Error parsing configuration data:", err)
	}

	opt := option.WithCredentialsFile(ApiConfig.FBServiceFile)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := fbApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	c := controller.Controller{}
	c.Init(ApiConfig.PGConn, ApiConfig.AuthSecret, authClient)

	db, err := db.InitDb(ApiConfig.PGConn)
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

	api.Handle("/items", rest.UpsertItem(db)).Methods("PUT")
	api.Handle("/items", rest.ListItems(db)).Methods("GET")
	api.Handle("/items/{itemid}", rest.GetItem(db)).Methods("GET")

	api.Handle("/ping", pingHandler(c))

	if ApiConfig.StaticDir != "" {
		r.PathPrefix("/").Handler(spaHandler(ApiConfig.StaticDir))
	}

	log.Println("Starting server on port " + ApiConfig.Port)
	err = http.ListenAndServe(":"+ApiConfig.Port, r)
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
		response = ui.Response{HTTPStatus: http.StatusOK}
		ui.Respond(res, response)
	})
}
