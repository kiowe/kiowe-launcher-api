package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateNewAccessToken(id int, devPubAcc bool) (string, error) {
	claims := jwt.MapClaims{
		"id":          id,
		"dev_pub_acc": devPubAcc,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func JwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": err.Error(),
	})
}
