package main

import (
	"database/sql"
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
	mux.Handle("/creditors/", authenticate(c, customersHandler(c)))
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
		origin := req.Header.Get("Origin")
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response, req.Header.Get("Origin"))
			return
		}
		t, err := req.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response = ui.CreateResponse(http.StatusUnauthorized,
					"No authentication token present in request cookies", nil)
				ui.Respond(res, response, origin)
				return
			}
			response = ui.CreateResponse(http.StatusBadRequest,
				"An authentication token needs to be present in request cookies", nil)
			ui.Respond(res, response, origin)
			return
		}
		auth, err := c.ValidateUser(t.Value)
		if err != nil {
			response = ui.CreateResponse(http.StatusUnauthorized,
				"Invalid authentication token", nil)
			ui.Respond(res, response, origin)
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
				ui.Respond(res, response, origin)
				return
			}
		}
		h.ServeHTTP(res, req)
	})
}

func loginHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		origin := req.Header.Get("Origin")
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response, origin)
			return
		}
		if req.Method == "POST" {
			var user model.User
			err := json.NewDecoder(req.Body).Decode(&user)
			if err != nil {
				log.Println("Error decoding credentials from payload:", err)
				response = ui.MakeErrorResponse(http.StatusBadRequest,
					fmt.Sprintf("Error decoding credentials from request body: %v", req.Body))

				ui.Respond(res, response, origin)
				return
			}

			token, err := c.Login(user.Username, user.Password)
			if err != nil {
				log.Printf("Unable to login with credentials: %v %v \nError: %v",
					user.Username, user.Password, err)
				response = ui.MakeErrorResponse(http.StatusUnauthorized,
					fmt.Sprintf("Unable to login with credentials"))

				ui.Respond(res, response, origin)
				return
			}

			response = ui.CreateResponse(http.StatusAccepted, "", nil)

			http.SetCookie(res, &http.Cookie{
				Name:  "token",
				Value: token.Token,
				Path:  "/",
			})
			ui.Respond(res, response, origin)
		}
		if req.Method == "DELETE" {
			var response ui.Response

			t, err := req.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					response = ui.CreateResponse(http.StatusUnauthorized,
						"No authentication token present in request cookies", nil)
					ui.Respond(res, response, origin)
					return
				}
				response = ui.CreateResponse(http.StatusBadRequest,
					"An authentication token needs to be present in request cookies", nil)
				ui.Respond(res, response, origin)
				return
			}

			err = c.Logout(t.Value)
			if err != nil {
				response = ui.CreateResponse(http.StatusInternalServerError,
					"Error logging out user", nil)
				ui.Respond(res, response, origin)
				return
			}
			response = ui.CreateResponse(http.StatusAccepted, "User logged out successfully", nil)
			ui.Respond(res, response, origin)
			return
		}
	})
}

func routesHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		origin := req.Header.Get("Origin")
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response, req.Header.Get("Origin"))
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
					response = ui.MakeErrorResponse(http.StatusInternalServerError,
						"An error occurred getting all routes")
				}
				response = ui.CreateResponse(http.StatusOK, "", routes)
			case 2:
				// routes/{route}/
				route := pathParams[1]
				customers, err := c.GetCreditorsOnRoute(route)
				if err != nil {
					switch err {
					case sql.ErrNoRows:
						ui.RespondError(res, http.StatusNoContent,
							fmt.Sprintln("No creditors found on route:", route))
					default:
						ui.RespondError(res, http.StatusInternalServerError,
							fmt.Sprintln("An error occured getting customers on route:", route))
					}
				}
				response = ui.CreateResponse(http.StatusOK, "", customers)
			}
		}
		ui.Respond(res, response, origin)
	})
}

func customersHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		origin := req.Header.Get("Origin")
		if req.Method == "OPTIONS" {
			response = ui.CreateResponse(http.StatusOK, "", nil)
			ui.RespondWithOptions(res, response, req.Header.Get("Origin"))
			return
		}
		// extract path parameters
		cleanPath := path.Clean(req.URL.Path)
		pathParams := strings.Split(cleanPath[1:], "/")

		if req.Method == "GET" {
			if len(pathParams) == 1 {
				// customers/
				q := req.URL.Query()
				switch len(q) {
				case 0:
					// no query params
					creditors, err := c.GetAllCreditors()
					if err != nil {
						ui.RespondError(res, http.StatusInternalServerError,
							"An error occurred getting all creditors")
						return
					}
					response = ui.CreateResponse(http.StatusOK, "", creditors)
					ui.Respond(res, response, origin)
				case 2:
					// customers?route={r}&name={n}
					route := q.Get("route")
					name := q.Get("searchname")
					creditor, err := c.GetCreditorByNameRoute(route, name)
					if err != nil {
						ui.RespondError(res, http.StatusInternalServerError,
							fmt.Sprintf("An error occured getting creditor %v on route %v",
								route, name))
					}
					response = ui.CreateResponse(http.StatusOK, "", creditor)
					ui.Respond(res, response, origin)
				}
			} else if len(pathParams) == 2 {
				// customers/{id}
				id, err := strconv.ParseInt(pathParams[1], 10, 64)
				if err != nil {
					ui.RespondError(res, http.StatusBadRequest,
						fmt.Sprintf("An error occured parsing creditor id %v",
							pathParams[1]))
					return
				}
				creditor, err := c.GetCreditorByID(id)
				if err != nil {
					ui.RespondError(res, http.StatusNoContent,
						fmt.Sprintf("An error occured getting creditor with id %v:",
							id, err))
					return
				}
				response = ui.CreateResponse(http.StatusOK, "", creditor)
				ui.Respond(res, response, origin)
			} else if len(pathParams) == 3 {
				// customers/{id}/credits or customers/{id}/payments
				id, err := strconv.ParseInt(pathParams[1], 10, 64)
				if err != nil {
					ui.RespondError(res, http.StatusBadRequest,
						fmt.Sprintf("An error occured parsing creditor id %v",
							pathParams[1]))
					return
				}
				if pathParams[2] == "credits" {
					credits, err := c.GetCreditsByCreditor(id)
					if err != nil {
						ui.RespondError(res, http.StatusInternalServerError,
							fmt.Sprintf("An error occured getting credits by creditor with id %v:",
								id, err))
						return
					}
					response = ui.CreateResponse(http.StatusOK, "", credits)
					ui.Respond(res, response, origin)
				} else if pathParams[2] == "payments" {
					payments, err := c.GetPaymentsByCreditor(id)
					if err != nil {
						ui.RespondError(res, http.StatusInternalServerError,
							fmt.Sprintf("An error occured getting payments by creditor with id %v:",
								id, err))
						return
					}
					response = ui.CreateResponse(http.StatusOK, "", payments)
					ui.Respond(res, response, origin)
				}
			}
		} else if req.Method == "POST" {
			if len(pathParams) == 1 {
				// creditors/
				var creditor model.Customer
				err := json.NewDecoder(req.Body).Decode(&creditor)
				if err != nil {
					log.Panicln("Error decoding body:", err)
					response = ui.MakeErrorResponse(http.StatusBadRequest,
						fmt.Sprintf("An error occured parsing request body: %v", req.Body))

					ui.Respond(res, response, origin)
					return
				}
				id, err := c.CreateCreditor(creditor)
				if err != nil {
					response = ui.MakeErrorResponse(http.StatusInternalServerError,
						fmt.Sprintf("An error occured creating creditor: %v", err))

					ui.Respond(res, response, origin)
					return
				}
				response = ui.CreateResponse(http.StatusCreated, "", id)
				ui.Respond(res, response, origin)
				return
			} else if len(pathParams) == 3 {
				// creditors/{id}/credits or creditors/{id}/payments
				if pathParams[2] == "credits" {
					var credit model.Credit
					err := json.NewDecoder(req.Body).Decode(&credit)
					if err != nil {
						log.Panicln("Error decoding body:", err)
						ui.RespondError(res, http.StatusBadRequest,
							fmt.Sprintf("An error occured parsing credit: %v", req.Body))
						return
					}
					err = c.CreateCredit(credit)
					if err != nil {
						ui.RespondError(res, http.StatusInternalServerError,
							fmt.Sprintf("An error occured creating credit: %v", err))
						return
					}
					response = ui.CreateResponse(http.StatusCreated, "", nil)
					ui.RespondWithOptions(res, response, req.Header.Get("Origin"))
					return
				} else if pathParams[2] == "payments" {
					var payment model.Payment
					err := json.NewDecoder(req.Body).Decode(&payment)
					if err != nil {
						log.Panicln("Error decoding body:", err)
						response = ui.MakeErrorResponse(http.StatusBadRequest,
							fmt.Sprintf("An error occured parsing payment: %v", req.Body))
						ui.Respond(res, response, origin)
						return
					}
					err = c.CreatePayment(payment)
					if err != nil {
						response = ui.MakeErrorResponse(http.StatusInternalServerError,
							fmt.Sprintf("An error occured creating payment: %v", err))
						ui.Respond(res, response, origin)
						return
					}
					response = ui.CreateResponse(http.StatusCreated, "", nil)
					ui.RespondWithOptions(res, response, req.Header.Get("Origin"))
					return
				}
			}
		}
	})
}

func defaultersHandler(c controller.Controller) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var response ui.Response
		origin := req.Header.Get("Origin")
		if req.Method == "GET" {
			d, err := c.GetAllDefaulters()
			if err != nil {
				response = ui.MakeErrorResponse(http.StatusInternalServerError,
					"An error occurred getting all defaulters")
				ui.Respond(res, response, origin)
				return
			}
			response = ui.CreateResponse(http.StatusOK, "", d)
			ui.Respond(res, response, origin)
		}
	})
}
