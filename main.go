package main

import (
	"backend/database"
	"backend/pkg/mysql"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r)

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Running in localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
