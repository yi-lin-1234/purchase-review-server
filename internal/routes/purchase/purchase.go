package purchaseRoutes

import (
	"github.com/gofiber/fiber/v2"
	purchaseHandler "github.com/yi-lin-1234/purchase-review-backend/internal/handlers/purchase"
)

func SetupPurchaseRoutes(router fiber.Router) {
	purchase := router.Group("/purchase")

	//=======================( POST )===========================
	// create a purchase
	purchase.Post("/", purchaseHandler.CreateNewPurchase)

	//=======================( GET )===========================

	// read all purchase
	purchase.Get("/", purchaseHandler.GetAllPurchases)
	// read one purchase
	purchase.Get("/:id", purchaseHandler.GetPurchaseById)
	// read purchases by category
	purchase.Get("category/:category", purchaseHandler.GetPurchaseByCategory)
	// read purchases by evaluation
	purchase.Get("evaluation/:evaluation", purchaseHandler.GetPurchaseByEvaluation)
	// read all distinct category
	purchase.Get("distinct/category", purchaseHandler.GetDistinctCategory)
	// read all distinct evaluation
	purchase.Get("distinct/evaluation", purchaseHandler.GetDistinctEvaluation)
	// read all distinct evaluation
	purchase.Get("search/:term", purchaseHandler.GetPurchasesByTerm)

	//=======================( PUT )===========================
	// update one purchase
	purchase.Put("/:id", purchaseHandler.UpdatePurchase)

	//=======================( DELETE )===========================
	// delete one purchase
	purchase.Delete("/:id", purchaseHandler.DeletePurchaseById)
}
