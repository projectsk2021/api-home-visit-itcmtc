package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/utils"
	"github.com/spf13/viper"
)

var UserClaims utils.JwtUserClaims

func UserProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(viper.GetString("token.secretKeyJWT")),
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			uuid, _ := uuid.Parse(claims["UserId"].(string))
			UserClaims = utils.JwtUserClaims{
				UserId: &uuid,
			}
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ErrorResponse(utils.Message{
				ErrorCode: "1204",
			}))
		},
	})

}
