package routes

import (
	"watchx-backend/handler"

	"github.com/gofiber/fiber/v2"
)

// type Coba struct {
// 	Data    string `json:"data"`
// 	Message string `json:"message"`
// }

func AccountRoutes(route fiber.Router) {
	route.Get("/", handler.GetAllAccounts)
	// r.Get("/user", func(ctx *fiber.Ctx) error {
	// 	return ctx.JSON(fiber.Map{
	// 		"data":    "Hello, World",
	// 		"message": "Hahahaha",
	// 	})
	// })

	// r.Get("/cobaclient", func(ctx *fiber.Ctx) error {
	// 	var coba1 Coba
	// 	a := fiber.AcquireAgent()
	// 	req := a.Request()
	// 	req.Header.SetMethod(fiber.MethodGet)
	// 	req.SetRequestURI("http://localhost:3000/user")

	// 	if err := a.Parse(); err != nil {
	// 		panic(err)
	// 	}

	// 	_, body, _ := a.Bytes()

	// 	json.Unmarshal(body, &coba1)
	// 	return ctx.JSON(coba1)
	// })
}
