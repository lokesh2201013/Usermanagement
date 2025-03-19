package controllers

import(
	"github.com/gofiber/fiber/v2"
	"github.com/lokesh2201013/Usermanagement/models"
	"github.com/lokesh2201013/Usermanagement/database" 
	"github.com/lokesh2201013/Usermanagement/utils"   
    "log"
)

func Register(c *fiber.Ctx)error{
	var user models.Admin
	if err:= c.BodyParser(&user); err!=nil{
	  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid input"})
	}

	var existingUser models.Admin

	if err:=database.DB.Where("email=?",user.Email).First(&existingUser).Error;err==nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use"})
	}

	hashpassword, err:=utils.HashPassword(user.Password)

	if err != nil {
		log.Println("Password hashing error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	user.Password = hashpassword

	if err:=database.DB.Create(&user).Error;err!=nil{
		log.Println("Database error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error{

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invald input"})
	}

	var user models.Admin

	if err:= database.DB.Where("email=?",loginData.Email).First(&user).Error;err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"User not found"})
	}
	
	if err := utils.CheckPassword(user.Password, loginData.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		log.Println("Token generation error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func GetAllAdmins(c *fiber.Ctx) error {
	var admins []models.Admin

	// selctet * from admins
	if err := database.DB.Find(&admins).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve admins"})
	}
	
	return c.Status(fiber.StatusOK).JSON(admins)
}