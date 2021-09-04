package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	createMovie(r)
	readMovie(r)
	updateMovie(r)
	deleteMovie(r)
	readAll(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}

func createMovie(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/movie").Handler(handler.CreateHandler())
}

func readMovie(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/movie/{slug}").Handler(handler.ReadHandler())
}

func readAll(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/movie").Handler(handler.ReadAll())
}

func updateMovie(router *mux.Router) {
	router.Methods(http.MethodPut).Path("/movie/{slug}").Handler(handler.UpdateHandler())
}

func deleteMovie(router *mux.Router) {
	router.Methods(http.MethodDelete).Path("/movie/{slug}").Handler(handler.DeleteHandler())
}
