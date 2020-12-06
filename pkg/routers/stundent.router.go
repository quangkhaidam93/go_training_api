package routers

import (
	"github.com/gorilla/mux"
	"github.com/quangkhaidam93/go_train_api/pkg/handlers"
)

func studentRouter(router *mux.Router, handler *handlers.Handler) {
	r := router.PathPrefix("/student").Subrouter()

	r.HandleFunc("", handler.StudentHandler.GetStudentHandler).Methods("GET")
	r.HandleFunc("", handler.StudentHandler.CreateNewStudentHandler).Methods("POST")
	r.HandleFunc("/{id}", handler.StudentHandler.UpdateStudentHandler).Methods("POST")
	r.HandleFunc("/{id}", handler.StudentHandler.DeleteStudentHandler).Methods("DELETE")
}
