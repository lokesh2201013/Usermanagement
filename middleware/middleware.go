package middleware

import (
	"log"
	//"reflect"
	"strings"
    "github.com/lokesh2201013/Usermanagement/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	//"github.com/google/uuid"
)

var secretKey = []byte("your-secret-key")

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenHeader := c.Get("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized Bearer token required",
			})
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			log.Println("Token validation error:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized token error"})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("ID", claims["ID"])

		return c.Next()
	}
}

func AdminOnly(handler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID:= c.Locals("ID")
		if ID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized ID required"})
		}

		var exists bool
		query := "SELECT EXISTS(SELECT 1 FROM admins WHERE id = ?);"
		err := database.DB.Raw(query, ID).Scan(&exists).Error
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
		
		if !exists {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Admin access required"})
		}

		return handler(c)
	}
}