package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/hmnayak/credit/controller"
	"github.com/hmnayak/credit/db"
	"github.com/hmnayak/credit/model"
	"github.com/hmnayak/credit/ui"
)

// AppConfig is a container of api configuration data
type AppConfig struct {
	DBConfig db.Config `yaml:"postgresdb"`
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

	c := controller.Controller{}
	c.Init(config.DBConfig)

	mux := http.NewServeMux()

	mux.Handle("/routes/", routesHandler(c))
	mux.Handle("/customers/", customersHandler(c))
	mux.Handle("/credits/", creditsHandler(c))
	mux.Handle("/payments/", paymentsHandler(c))

	err = http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
	log.Println("Listening on port 8001...")
}

func routesHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")

			switch len(pathParams) {
			case 1:
				// routes/
				routes, err := c.GetAllRoutes()
				if err != nil {
					response = ui.MakeErrorResponse("An error occurred getting all routes")
				}
				response = ui.CreateResponse(http.StatusOK, "", routes)
			case 2:
				// routes/{route}/
				customers, err := c.GetCreditorsOnRoute(pathParams[1])
				if err != nil {
					response =
						ui.MakeErrorResponse("An error occured getting customers on route - " +
							pathParams[1])
				}
				response = ui.CreateResponse(http.StatusOK, "", customers)
			}
		}
		ui.Respond(res, response)
	})
}

func customersHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusCreated, "", nil)
			ui.RespondWithOptions(res, response)
			return
		}
		if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")

			if len(pathParams) == 1 {
				// customers/
				q := req.URL.Query()
				switch len(q) {
				case 0:
					// no query params
					creditors, err := c.GetAllCreditors()
					if err != nil {
						response = ui.MakeErrorResponse("An error occurred getting all creditors")
					}
					response = ui.CreateResponse(http.StatusOK, "", creditors)
				case 2:
					// customers?route={r}&name={n}
					route := q.Get("route")
					name := q.Get("name")
					creditor, err := c.GetCreditorByNameRoute(route, name)
					if err != nil {
						response = ui.MakeErrorResponse(
							fmt.Sprintf("An error occured getting creditor %v on route %v",
								route, name))
					}
					response = ui.CreateResponse(http.StatusOK, "", creditor)
				}
			} else if len(pathParams) == 2 {
				// customers/{id}
				id, err := strconv.ParseInt(pathParams[1], 10, 64)
				if err != nil {
					response = ui.MakeErrorResponse(
						fmt.Sprintf("An error occured parsing creditor id %v",
							pathParams[1]))
				}
				creditor, err := c.GetCreditorByID(id)
				if err != nil {
					response = ui.MakeErrorResponse(
						fmt.Sprintf("An error occured getting creditor with id %v:",
							id, err))
				}
				response = ui.CreateResponse(http.StatusOK, "", creditor)
			}
		} else if req.Method == "POST" {
			var creditor model.Customer
			err := json.NewDecoder(req.Body).Decode(&creditor)
			if err != nil {
				log.Panicln("Error decoding body:", err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured parsing request body: %v", req.Body))
			}
			id, err := c.CreateCreditor(creditor)
			if err != nil {
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured creating creditor: %v", err))
			}
			response = ui.CreateResponse(http.StatusCreated, "", id)
		}
		ui.Respond(res, response)
	})
}

func paymentsHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response)
			return
		}
		if req.Method == "POST" {
			var t model.Transaction
			err := json.NewDecoder(req.Body).Decode(&t)
			if err != nil {
				log.Panicln("Error decoding body:", err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured parsing payment: %v", req.Body))
				ui.Respond(res, response)
				return
			}
			err = c.CreatePayment(t)
			if err != nil {
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured creating payment: %v", err))
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusCreated, "", nil)
			ui.RespondWithOptions(res, response)
		} else if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")
			id, err := strconv.ParseInt(pathParams[1], 10, 64)
			if err != nil {
				log.Panicln(err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured parsing creditor id %v",
						pathParams[1]))
				ui.Respond(res, response)
				return
			}
			payments, err := c.GetPaymentsByCreditor(id)
			if err != nil {
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured getting payments by creditor with id %v:",
						id, err))
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusOK, "", payments)
			ui.Respond(res, response)
		}
	})
}

func creditsHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response)
			return
		}
		if req.Method == "POST" {
			var t model.Transaction
			err := json.NewDecoder(req.Body).Decode(&t)
			if err != nil {
				log.Panicln("Error decoding body:", err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured parsing credit: %v", req.Body))
				ui.Respond(res, response)
				return
			}
			err = c.CreateCredit(t)
			if err != nil {
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured creating credit: %v", err))
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusCreated, "", nil)
			ui.RespondWithOptions(res, response)
		} else if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")
			id, err := strconv.ParseInt(pathParams[1], 10, 64)
			if err != nil {
				log.Panicln(err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured parsing creditor id %v",
						pathParams[1]))
				ui.Respond(res, response)
				return
			}
			credits, err := c.GetCreditsByCreditor(id)
			if err != nil {
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured getting credits by creditor with id %v:",
						id, err))
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusOK, "", credits)
			ui.Respond(res, response)
		}
	})
}
