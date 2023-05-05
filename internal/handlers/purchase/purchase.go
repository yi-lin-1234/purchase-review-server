package purchaseHandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yi-lin-1234/purchase-review-backend/database"
	"github.com/yi-lin-1234/purchase-review-backend/internal/model"
	"strings"
)

func GetAllPurchases(c *fiber.Ctx) error {
	db := database.DB

	// slice of purchases
	var purchases []model.Purchase

	// find all purchases in the database, pass in reference to direct modify it
	db.Find(&purchases)

	// If no purchase is present return an error
	if len(purchases) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no purchases present",
			"data":    nil})
	}

	// Else return all purchases
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "purchases Found",
		"data":    purchases})
}

func CreateNewPurchase(c *fiber.Ctx) error {
	db := database.DB
	purchase := new(model.Purchase)

	// store the body in the purchase and return error if encountered
	err := c.BodyParser(&purchase)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "review your input",
			"data":    err.Error()})
	}
	// add an uuid to the note
	purchase.ID = uuid.New()
	// create the purchase and return error if encountered
	fmt.Println(purchase)
	err = db.Create(&purchase).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create purchase",
			"data":    err.Error()})
	}

	// Return the created note
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "created purchase",
		"data":    purchase})
}

func GetPurchaseById(c *fiber.Ctx) error {
	db := database.DB
	purchase := new(model.Purchase)

	// read the param id
	id := c.Params("id")

	// Find the purchase with the given id
	db.Find(&purchase, "id = ?", id)

	// if no such purchase present return an error
	if purchase.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no purchase exist",
			"data":    nil})
	}

	// return the purchase with the id
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "purchase Found",
		"data":    purchase})
}

func GetPurchaseByCategory(c *fiber.Ctx) error {
	db := database.DB
	var purchases []model.Purchase

	// read the param category
	category := c.Params("category")

	// Find the purchases with the given category
	db.Find(&purchases, "category = ?", category)

	// if no such purchase exist return an error
	if len(purchases) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no purchases present",
			"data":    nil})
	}

	// Else return all purchases with that category
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "purchases found",
		"data":    purchases})
}

func GetPurchaseByEvaluation(c *fiber.Ctx) error {
	db := database.DB
	var purchases []model.Purchase

	// read the param evaluation
	evaluation := c.Params("evaluation")

	// Find the purchases with the given evaluation
	db.Find(&purchases, "evaluation = ?", evaluation)

	// if no such purchase exist return an error
	if len(purchases) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no purchases present",
			"data":    nil})
	}

	// Else return all purchases with that evaluation
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "purchases found",
		"data":    purchases})
}

func GetDistinctCategory(c *fiber.Ctx) error {
	db := database.DB
	var category []string

	// get all distinct category
	db.Model(&model.Purchase{}).Distinct().Pluck("Category", &category)

	// if no such category exist return an error
	if len(category) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no category present",
			"data":    nil})
	}

	// Else return all distinct category
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "distinct category found",
		"data":    category})
}

func GetDistinctEvaluation(c *fiber.Ctx) error {
	db := database.DB
	var evaluation []string

	// get all distinct evaluation
	db.Model(&model.Purchase{}).Distinct().Pluck("Evaluation", &evaluation)

	// if no such evaluation exist return an error
	if len(evaluation) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no evaluation present",
			"data":    nil})
	}

	// Else return all distinct evaluation
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "distinct evaluation found",
		"data":    evaluation})
}

// GetPurchasesByTerm returns purchases that match the search term in their names
func GetPurchasesByTerm(c *fiber.Ctx) error {
	db := database.DB
	var purchases []model.Purchase

	// read the param term
	term := c.Params("term")

	// Check if term is empty or only consists of whitespace
	if strings.TrimSpace(term) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "search term cannot be empty",
			"data":    nil})
	}

	// Find the purchases with the name containing the search term (case-insensitive)
	db.Where("LOWER(name) LIKE LOWER(?)", "%"+term+"%").Find(&purchases)

	// if no such purchase exist return an error
	if len(purchases) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no purchases present",
			"data":    nil})
	}

	// Else return all purchases with that contains search term
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "purchases found",
		"data":    purchases})
}

func UpdatePurchase(c *fiber.Ctx) error {

	db := database.DB

	purchase := new(model.Purchase)

	// read the param id
	id := c.Params("id")

	// find the purchase with the given id
	db.Find(&purchase, "id = ?", id)

	// if no such purchase exist return an error
	if purchase.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no such purchase exist",
			"data":    nil})
	}

	// store the body containing the updated data and return error if encountered
	temp := new(model.Purchase)
	err := c.BodyParser(&temp)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "review your input",
			"data":    err})
	}

	// Edit the purchase

	purchase.Name = temp.Name
	purchase.Price = temp.Price
	purchase.Amount = temp.Amount
	purchase.Category = temp.Category
	purchase.Evaluation = temp.Evaluation
	purchase.Note = temp.Note
	purchase.ImageUrl = temp.ImageUrl
	purchase.Link = temp.Link

	// save the changes
	db.Save(&purchase)

	// return the updated purchase
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "purchase updated",
		"data":    purchase})
}

func DeletePurchaseById(c *fiber.Ctx) error {
	db := database.DB
	purchase := new(model.Purchase)

	// read the param id
	id := c.Params("id")

	// find the purchase with the given id
	db.Find(&purchase, "id = ?", id)

	// if no such purchase exist return an error
	if purchase.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "no such purchase exist",
			"data":    nil})
	}

	// delete the purchase and return error if encountered
	err := db.Delete(&purchase, "id = ?", id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to delete purchase",
			"data":    nil})
	}

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "purchase deleted"})
}
