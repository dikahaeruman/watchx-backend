package routes

import (
	"watchx-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func MovieRoutes(route fiber.Router) {
	route.Get("/top_rated", handler.GetTopRatedMovies)
	route.Get("/popular", handler.GetPopularMovies)
}
