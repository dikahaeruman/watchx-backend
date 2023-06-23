package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"watchx-backend/model"

	"github.com/gofiber/fiber/v2"
)

func GetTopRatedMovies(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	var movieApiResponse model.MovieApiResponse
	var topRatedURI = fmt.Sprintf("https://api.themoviedb.org/3/movie/top_rated?page=%s&api_key=%s", page, os.Getenv("TMDB_APIKEY"))
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(topRatedURI)

	if err := agent.Parse(); err != nil {
		log.Fatal(err)
	}

	_, body, err := agent.Bytes()

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &movieApiResponse)
	return c.JSON(movieApiResponse)
}

func GetPopularMovies(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	var movieApiResponse model.MovieApiResponse
	var topRatedURI = fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?page=%s&api_key=%s", page, os.Getenv("TMDB_APIKEY"))
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.Header.SetMethod(fiber.MethodGet)
	request.SetRequestURI(topRatedURI)

	if err := agent.Parse(); err != nil {
		log.Fatal(err)
	}

	_, body, err := agent.Bytes()

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &movieApiResponse)
	return c.JSON(movieApiResponse)
}
