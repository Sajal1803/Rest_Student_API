package main

import (
	"simple-rest-api/datastore"
	"simple-rest-api/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/students", h.GetByID)
	app.POST("/students", h.Create)
	app.PUT("/students/{id}", h.Update)
	app.DELETE("/students/{id}", h.Delete)

	app.Start()
}
