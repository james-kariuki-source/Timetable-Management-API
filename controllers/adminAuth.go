package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	db "github.com/james-kariuki-source/Timetable-Management-API/connection"
	"github.com/james-kariuki-source/Timetable-Management-API/models"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var admin models.Admin

	db.DB.Where("email=?", data["email"]).First(&admin)

	if admin.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(admin.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Invalid password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(admin.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Invalid or Expired token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Success": true,
		"message": "Successful login",
	})

}

func Admin(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Unauthenticated Admin",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var admin models.Admin

	db.DB.Where("id=?", claims.Issuer).First(&admin)

	return c.JSON(admin)

}

func AdminLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"success":true,
		"message":"Successful logout",
	})
}

// func Register(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil{
// 		return err
// 	}

// 	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

// 	admin := models.Admin{
// 		Name: data["name"],
// 		Email: data["email"],
// 		Password: password,
// 	}

// 	db.DB.Create(&admin)

// 	return c.JSON(admin)
// }
