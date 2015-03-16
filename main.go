package main

import (
  "net/http"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/unrolled/render"
  "os"
)

func main() {
  n := negroni.Classic()
  router := mux.NewRouter()
  router.Handle("/", http.FileServer(http.Dir("./views")))

  r := render.New()

  api := router.PathPrefix("/api").Subrouter()
  api.HandleFunc("/services", func(resp http.ResponseWriter, req *http.Request) {
  	r.JSON(resp, http.StatusOK, map[string]string{
  		"hello": "json",
  	})
  }).Methods("GET")

  n.UseHandler(router)

  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

  n.Run(":" + port)
}

