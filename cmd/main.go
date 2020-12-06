package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/quangkhaidam93/go_train_api/pkg/database"
	"github.com/quangkhaidam93/go_train_api/pkg/handlers"
	"github.com/quangkhaidam93/go_train_api/pkg/routers"
)

func main() {
	fmt.Println("Hello World")
	sqlDb := database.DbConn()
	defer sqlDb.Close()
	gormDb := database.DbConnViaGorm()
	handler := handlers.HandlerInit(sqlDb, gormDb)
	myRouter := routers.RouterInit(handler)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
