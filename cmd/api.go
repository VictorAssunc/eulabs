package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"eulabs/pkg/api"
	"eulabs/pkg/repository"
	"eulabs/pkg/service"
)

func main() {
	db, err := sql.Open("mysql", "dev:5up3r53cr37@tcp(localhost:8000)/eulabs")
	if err != nil {
		log.Fatal(err)
	}

	handler := api.NewHandler(service.NewProduct(repository.NewProduct(db)))

	e := echo.New()
	e.Use(middleware.Logger())

	g := e.Group("/products")
	g.GET("/:id", handler.GetProduct)

	admin := g.Group("", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}

		return false, nil
	}))
	admin.POST("/", handler.CreateProduct)
	admin.PUT("/:id", handler.UpdateProduct)
	admin.DELETE("/:id", handler.DeleteProduct)

	s := http.Server{
		Addr:        ":8001",
		Handler:     e,
		ReadTimeout: time.Second,
	}
	log.Println("Server is running on http://localhost:8001")
	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
