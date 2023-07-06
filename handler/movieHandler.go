package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"watchx-backend/config"
	"watchx-backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTopRatedMovies(c *fiber.Ctx) error {
	return GetMovies(c, "top_rated")
}

func GetPopularMovies(c *fiber.Ctx) error {
	return GetMovies(c, "popular")
}

func GetMovies(c *fiber.Ctx, urlType string) error {
	userCollection := config.MI.DB.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var userResponse model.User
	var movieApiResponse model.MovieApiResponse

	page := c.Query("page", "1")
	email := c.Get("email")
	var topRatedURI = fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?page=%s&api_key=%s", urlType, page, os.Getenv("TMDB_APIKEY"))
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

	if error := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&userResponse); error != nil {
		return c.JSON(movieApiResponse)
	}

	for i, v := range movieApiResponse.Results {
		for _, w := range userResponse.User_details.Watched_movies {
			if v.ID == w {
				movieApiResponse.Results[i].Is_watched = true
			}
		}
	}

	return c.JSON(movieApiResponse)
}
