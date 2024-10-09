package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/VuThanhThien/golang-gorm-postgres/gateway/initializers"
	"github.com/VuThanhThien/golang-gorm-postgres/gateway/middleware"
	"github.com/VuThanhThien/golang-gorm-postgres/gateway/pkg/logger"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ServiceURL holds the base URL for each microservice
var ServiceURL = map[string]string{}

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	ServiceURL["auth"] = config.AUTH_SERVICE_HOST
	ServiceURL["users"] = config.AUTH_SERVICE_HOST
	ServiceURL["merchant"] = config.MERCHANT_SERVICE_HOST
	ServiceURL["product"] = config.MERCHANT_SERVICE_HOST
	ServiceURL["inventory"] = config.MERCHANT_SERVICE_HOST
	ServiceURL["categories"] = config.MERCHANT_SERVICE_HOST
	ServiceURL["variants"] = config.MERCHANT_SERVICE_HOST
	ServiceURL["product-images"] = config.MERCHANT_SERVICE_HOST
	ServiceURL["order"] = config.ORDER_SERVICE_HOST
	ServiceURL["payment"] = config.PAYMENT_SERVICE_HOST
}

func main() {
	log := logger.NewLogger()

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("🚀 Could not load environment variables")
	}

	// Create a new mux router
	router := mux.NewRouter()

	// CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:" + config.PORT}),
		handlers.AllowCredentials(),
	)

	// Create logging middleware
	loggingMiddleware := middleware.Logging(log)
	operationIDMiddleware := middleware.OperationID
	// Wrap handlers with logging middleware
	handler := corsMiddleware(operationIDMiddleware(loggingMiddleware(router)))

	router.PathPrefix("/api/auth/").Handler(createReverseProxy("auth"))
	router.PathPrefix("/api/users/").Handler(createReverseProxy("users"))
	router.PathPrefix("/api/merchant/").Handler(createReverseProxy("merchant"))
	router.PathPrefix("/api/product/").Handler(createReverseProxy("product"))
	router.PathPrefix("/api/inventory/").Handler(createReverseProxy("inventory"))
	router.PathPrefix("/api/categories/").Handler(createReverseProxy("categories"))
	router.PathPrefix("/api/variants/").Handler(createReverseProxy("variants"))
	router.PathPrefix("/api/product-images/").Handler(createReverseProxy("product-images"))
	router.PathPrefix("/api/order/").Handler(createReverseProxy("order"))
	router.PathPrefix("/api/payment/").Handler(createReverseProxy("payment"))

	// Start the server
	fmt.Println("Starting gateway on http://localhost:" + config.PORT)
	log.Fatal().Err(http.ListenAndServe(":"+config.PORT, handler)).Msg("🚀 Gateway server started")
}

func createReverseProxy(service string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := url.Parse(ServiceURL[service])
		if err != nil {
			http.Error(w, "Error parsing service URL", http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.URL.Path = "/api/" + service + r.URL.Path[len("/api/"+service):]
			req.Host = url.Host
			req.Header.Set("X-Forwarded-Host", r.Host)
			req.Header.Set("X-Forwarded-Method", r.Method)

			// Copy all headers from the original request
			for key, values := range r.Header {
				req.Header[key] = values
			}
		}

		proxy.ModifyResponse = func(resp *http.Response) error {
			// Allow CORS
			resp.Header.Set("Access-Control-Allow-Origin", "*")
			resp.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			resp.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return nil
		}

		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadGateway)
		}

		proxy.ServeHTTP(w, r)
	}
}
