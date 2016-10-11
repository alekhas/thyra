package gateway

import (
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/uruddarraju/thyra/pkg/api/handlers/restapis"
	"github.com/uruddarraju/thyra/pkg/auth/authn"
	"github.com/uruddarraju/thyra/pkg/auth/authn/tokenfile"
	authnmiddleware "github.com/uruddarraju/thyra/pkg/middleware/authn"
)

type Gateway struct {
	Address       string
	DefaultRouter *mux.Router
	Server        *http.Server
	Authenticator authn.Authenticator
}

type GatewayOpts struct{}

func NewDefaultGateway() *Gateway {
	defaultRouter := mux.NewRouter()
	authn := tokenfile.NewTokenAuthenticator("")
	AddDefaultHandlers(defaultRouter, authn)
	srv := &http.Server{
		Handler:      defaultRouter,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Gateway{
		DefaultRouter: defaultRouter,
		Server:        srv,
	}
}

func (gw *Gateway) Start() {
	gw.Server.ListenAndServe()
	glog.Fatalf("Server quit.....")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Thyra....!\n"))
}

func AddDefaultHandlers(router *mux.Router, authenticator authn.Authenticator) {

	router.HandleFunc("/", authnmiddleware.Authenticate(authenticator, restapis.RestAPIHandler))
	router.HandleFunc("/hello", authnmiddleware.Authenticate(authenticator, HelloHandler))
	router.HandleFunc("/metrics", authnmiddleware.Authenticate(authenticator, HelloHandler))
	router.HandleFunc("/healthz", authnmiddleware.Authenticate(authenticator, HelloHandler))

	router.HandleFunc("/restapis", authnmiddleware.Authenticate(authenticator, restapis.RestAPIHandler))
}
