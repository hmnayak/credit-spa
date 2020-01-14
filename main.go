package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/hmnayak/credit/db"
	"github.com/hmnayak/credit/model"
	"github.com/rs/cors"
)

func main() {
	pgConfig := db.Config{
		ConnectionString: "host=localhost port=5432 user=mnayak dbname=credit sslmode=disable",
	}
	pgDb, err := db.InitDb(pgConfig)
	if err != nil {
		log.Panicln("Error InitDb: %v", err)
	}
	m := model.New(pgDb)

	mux := http.NewServeMux()
	c := cors.AllowAll()
	mux.Handle("/routes/", c.Handler(routesHandler(m)))
	mux.Handle("/customers/", c.Handler(customersHandler(m)))
	mux.Handle("/credits/", c.Handler(creditsHandler(m)))
	mux.Handle("/payments/", c.Handler(paymentsHandler(m)))
	err = http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}

func routesHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")

			switch len(pathParams) {
			case 1:
				// routes/
				getAllRoutes(m, res, req)
			case 2:
				// routes/{route}/
				getCustomersOnRoute(m, pathParams[1], res, req)
			}
		}
	})
}

func customersHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
					getAllCustomers(m, res, req)
				case 2:
					// customers?route={r}&name={n}
					getCustomerByNameRoute(m, q.Get("route"), q.Get("name"), res, req)
				}
			} else if len(pathParams) == 2 {
				// customers/{id}
				id, err := strconv.ParseInt(pathParams[1], 10, 64)
				if err != nil {
					getCustomerByID(m, id, res, req)
				}
			}
		} else if req.Method == "POST" {
			createCustomer(m, res, req)
		}
	})
}

func paymentsHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			createPayment(m, res, req)
		} else if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")
			id, err := strconv.ParseInt(pathParams[1], 10, 64)
			if err != nil {
				log.Println(err)
			}
			getPaymentsByCustomer(m, id, res, req)
		}
	})
}

func creditsHandler(m *model.Model) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Method == "POST" {
			createCredit(m, res, req)
		} else if req.Method == "GET" {
			// remove trailing slash
			cleanPath := path.Clean(req.URL.Path)
			pathParams := strings.Split(cleanPath[1:], "/")
			id, err := strconv.ParseInt(pathParams[1], 10, 64)
			if err != nil {
				log.Println(err)
			}
			getCreditsByCustomer(m, id, res, req)
		}
	})
}

func getAllRoutes(m *model.Model, res http.ResponseWriter, req *http.Request) {
	r, err := m.Db.GetRoutes()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error GetRoutes: %v", err)
	}
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(r)
}

func getCustomersOnRoute(m *model.Model, route string, res http.ResponseWriter, req *http.Request) {
	c, err := m.Db.GetCustomersOnRoute(route)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error getCustomersOnRoute: %v", err)
	}
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(c)
}

func getAllCustomers(m *model.Model, res http.ResponseWriter, req *http.Request) {
	customers, err := m.Db.GetAllCustomers()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error SelectCustomers: %v", err)
	}

	for _, c := range customers {
		c.CalculateDueAmount()
	}

	res.Header().Set("content-type", "application/json")
	json.NewEncoder(res).Encode(customers)
}

func getCustomerByID(m *model.Model, id int64, res http.ResponseWriter, req *http.Request) {
	customer, err := m.Db.GetCustomerByID(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error SelectCustomers: %v", err)
	}
	// customer.CalculateDueAmount()
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(customer)
}

func getCustomerByNameRoute(m *model.Model, route string, name string, res http.ResponseWriter, req *http.Request) {
	customer, err := m.Db.GetCustomerByNameRoute(route, name)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error SelectCustomers: %v", err)
	}
	// customer.CalculateDueAmount()
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(customer)
}

func createCustomer(m *model.Model, res http.ResponseWriter, req *http.Request) {
	var c model.Customer
	err := json.NewDecoder(req.Body).Decode(&c)
	if err != nil {
		log.Panicln("Error decoding body:", err)
	}
	id, err := m.Db.CreateCustomer(c)
	if err != nil {
		log.Panicln("Error CreateCustomer:", err)
	}
	customer, _ := m.Db.GetCustomerByID(id)
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(customer)
}

func createCredit(m *model.Model, res http.ResponseWriter, req *http.Request) {
	var t model.Transaction
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		log.Fatalln("Error decoding body:", err)
	}
	err = m.Db.CreateCredit(t)
	if err != nil {
		log.Fatalln("Error CreateCredit:", err)
	}

	res.WriteHeader(http.StatusCreated)
}

func createPayment(m *model.Model, res http.ResponseWriter, req *http.Request) {
	var t model.Transaction
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		log.Println("Error decoding body:", err)
	}
	log.Println(t)
	err = m.Db.CreatePayment(t)
	if err != nil {
		log.Println("Error CreatePayment:", err)
	}

	res.WriteHeader(http.StatusCreated)
}

func getPaymentsByCustomer(m *model.Model, id int64, res http.ResponseWriter, req *http.Request) {
	p, err := m.Db.GetPaymentsByCustomer(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error getCustomersOnRoute: %v", err)
	}
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(p)
}

func getCreditsByCustomer(m *model.Model, id int64, res http.ResponseWriter, req *http.Request) {
	c, err := m.Db.GetCreditsByCustomer(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		log.Panicln("Error getCustomersOnRoute: %v", err)
	}
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(res).Encode(c)
}
