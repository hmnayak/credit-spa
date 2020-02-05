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
	DBConfig   db.Config `yaml:"postgresdb"`
	AuthSecret string    `yaml:"authsecret"`
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
	c.Init(config.DBConfig, config.AuthSecret)

	mux := http.NewServeMux()

	mux.Handle("/login/", loginHandler(c))
	mux.Handle("/routes/", authenticate(c, routesHandler(c)))
	mux.Handle("/customers/", authenticate(c, customersHandler(c)))
	mux.Handle("/credits/", authenticate(c, creditsHandler(c)))
	mux.Handle("/payments/", authenticate(c, paymentsHandler(c)))
	mux.Handle("/defaulters/", authenticate(c, defaultersHandler(c)))

	err = http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
	log.Println("Listening on port 8001...")
}

func authenticate(c controller.Controller, h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		t, err := req.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response = ui.CreateResponse(http.StatusUnauthorized,
					"No authentication token present in request cookies", nil)
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusBadRequest,
				"An authentication token needs to be present in request cookies", nil)
			ui.Respond(res, response)
			return
		}
		auth, err := c.ValidateUser(t.Value)
		if err != nil {
			response = ui.CreateResponse(http.StatusUnauthorized,
				"Invalid authentication token", nil)
			ui.Respond(res, response)
			return
		}

		if auth == "r" {
			switch req.Method {
			case "OPTIONS":
			case "PUT":
			case "POST":
			case "DELETE":
			case "PATCH":
				response = ui.CreateResponse(http.StatusUnauthorized,
					"Credentials not authorized to perform the operation", nil)
				ui.Respond(res, response)
				return
			}
		}

		h.ServeHTTP(res, req)
	})
}

func loginHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response)
			return
		}
		if req.Method == "POST" {
			var user model.User
			err := json.NewDecoder(req.Body).Decode(&user)
			if err != nil {
				log.Panicln("Error decoding credentials from payload:", err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured parsing request body: %v", req.Body))

				ui.Respond(res, response)
				return
			}

			token, err := c.Login(user.Username, user.Password)
			if err != nil {
				log.Panicln("Unable to login with credentials:", err)
				response = ui.MakeErrorResponse(
					fmt.Sprintf("Unable to login with credentials: %v", req.Body))

				ui.Respond(res, response)
				return
			}

			response = ui.CreateResponse(http.StatusAccepted, "", nil)

			http.SetCookie(res, &http.Cookie{
				Name:  "token",
				Value: token.Token,
				Path:  "/",
			})
			ui.Respond(res, response)
		}
		if req.Method == "DELETE" {
			var response ui.Response

			t, err := req.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					response = ui.CreateResponse(http.StatusUnauthorized,
						"No authentication token present in request cookies", nil)
					ui.Respond(res, response)
					return
				}
				response = ui.CreateResponse(http.StatusBadRequest,
					"An authentication token needs to be present in request cookies", nil)
				ui.Respond(res, response)
				return
			}

			err = c.Logout(t.Value)
			if err != nil {
				response = ui.CreateResponse(http.StatusInternalServerError,
					"Error logging out user", nil)
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusAccepted, "User logged out successfully", nil)
			ui.Respond(res, response)
			return
		}
	})
}

func routesHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response)
			return
		}
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
			response = ui.CreateResponse(http.StatusOK, "", nil)
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
					name := q.Get("searchname")
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

				ui.Respond(res, response)
				return
			}
			id, err := c.CreateCreditor(creditor)
			if err != nil {
				response = ui.MakeErrorResponse(
					fmt.Sprintf("An error occured creating creditor: %v", err))

				ui.Respond(res, response)
				return
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
			var t model.Payment
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
			var t model.Credit
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

func defaultersHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		if req.Method == "GET" {
			d, err := c.GetAllDefaulters()
			if err != nil {
				response = ui.MakeErrorResponse("An error occurred getting all defaulters")
				ui.Respond(res, response)
				return
			}
			response = ui.CreateResponse(http.StatusOK, "", d)
			ui.Respond(res, response)
		}
	})
}
