package routers

import (
	"github.com/gorilla/mux"
	"github.com/quangkhaidam93/go_train_api/pkg/handlers"
)

func RouterInit(handler *handlers.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	studentRouter(router, handler)
	classRouter(router, handler)

	return router
}
