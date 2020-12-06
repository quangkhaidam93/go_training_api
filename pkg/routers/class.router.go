package routers

import (
	"github.com/gorilla/mux"
	"github.com/quangkhaidam93/go_train_api/pkg/handlers"
)

func classRouter(router *mux.Router, handler *handlers.Handler) {
	r := router.PathPrefix("/class").Subrouter()

	r.HandleFunc("", handler.ClassHandler.GetClassHandler).Methods("GET")
	r.HandleFunc("", handler.ClassHandler.CreateClassHandler).Methods("POST")
}
