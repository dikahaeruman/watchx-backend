package handler

import (
	"context"
	"log"
	"math"
	"strconv"
	"time"
	"watchx-backend/config"
	"watchx-backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllAccounts(c *fiber.Ctx) error {
	accountCollection := config.MI.DB.Collection("accounts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var accounts []model.Account

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"account_id": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit int64 = int64(limitVal)

	total, _ := accountCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := accountCollection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Accounts Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var account model.Account
		cursor.Decode(&account)
		accounts = append(accounts, account)
	}

	last := math.Ceil(float64(total) / float64(limit))

	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      accounts,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}
