package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudhanshu-chauhan/mcrap-web/models/db"
	"github.com/sudhanshu-chauhan/mcrap-web/utils"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func main() {
	utils.LoadEnv()
	db.GetConnection()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
