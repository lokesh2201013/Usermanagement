package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lokesh2201013/Usermanagement/models"
	"github.com/lokesh2201013/Usermanagement/database"
	"log"
	"time"
	"github.com/google/uuid"
)

func Createusers(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
    
	adminID, ok := c.Locals("admin_id").(string) // Convert to string first
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	} 
	parsedAdminID, err := uuid.Parse(adminID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	
	var existingUser string
	if err:=database.DB.Where("email=?",user.Email).First(&existingUser).Error;err==nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use"})
	}

    user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.UpdatedBy = parsedAdminID
	user.Createdby = parsedAdminID


	if err:=database.DB.Create(&user).Error;err!=nil{
		log.Println("Database error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func Updateusers(c *fiber.Ctx) error{
	var user models.User
	if err:= c.BodyParser(&user); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid input"})
	}
	
	adminID, ok := c.Locals("admin_id").(string) // Convert to string first
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	} 
	parsedAdminID, err := uuid.Parse(adminID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	var existingUser models.User
	if err:=database.DB.Where("id=?",user.ID).First(&existingUser).Error;err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"User not found"})
	}
	
	user.UpdatedAt = time.Now()
	user.UpdatedBy = parsedAdminID

	if err:=database.DB.Save(&user).Error;err!=nil{
		log.Println("Database error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}