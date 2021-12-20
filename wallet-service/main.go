package main

import (
	"outsource-dao/wallet-dapp/handler"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/genToken/:clientID", handler.GenTokenController)
	app.Post("/transferToken", handler.TransferTokenController)
	app.Post("/getTokenBalance", handler.TokenBalanceController)
	app.Post("/swap", handler.SwapTokenController)
	app.Post("/approve", handler.ApproveController)

	app.Listen(":4551")
}
