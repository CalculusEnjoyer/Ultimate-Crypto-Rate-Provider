package main

import (
	"genesis-test-task/services/api/rest"
	"genesis-test-task/services/api/rest/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(rest.Api+rest.Rate, controllers.GetRate)
	r.Post(rest.Api+rest.AddEmails, controllers.AddEmail)
	r.Post(rest.Api+rest.SendEmails, controllers.SendEmails)

	http.ListenAndServe(":8080", r)
}
